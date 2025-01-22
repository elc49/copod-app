package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/elc49/copod/cache"
	"github.com/elc49/copod/config"
	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/util"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var p *paystackClient

type Paystack interface {
	ChargeMpesa(context.Context, uuid.UUID, MpesaCharge) (*MpesaChargeResponse, error)
	HappyPaystack(context.Context, *PaystackWebhook) error
}

var _ Paystack = (*paystackClient)(nil)

type paystackClient struct {
	log               *logrus.Logger
	sqlStore          *sql.Queries
	mu                sync.Mutex
	http              *http.Client
	paymentController controller.PaymentController
	redis             *redis.Client
}

func New(sqlStore *sql.Queries) {
	log := logger.GetLogger()
	p = &paystackClient{
		log,
		sqlStore,
		sync.Mutex{},
		&http.Client{},
		controller.GetPaymentController(),
		cache.GetCache().Redis(),
	}
	log.Infoln("paystackservice: Connected")
}

func GetPaystackService() Paystack {
	return p
}

func (p *paystackClient) ChargeMpesa(ctx context.Context, paymentFor uuid.UUID, input MpesaCharge) (*MpesaChargeResponse, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var chargeResponse *MpesaChargeResponse
	chargeApi := config.C.Paystack.BaseApi + "/charge"
	input.Provider.Provider = "mpesa"
	if !config.IsProd() {
		input.Provider.Phone = config.C.Paystack.MobileTestAccount
	}

	fees := 0
	switch input.Reason {
	case model.PaymentReasonLandRegistry.String():
		i, err := strconv.Atoi(config.C.Paystack.LandFees)
		if err != nil {
			p.log.WithError(err).WithFields(logrus.Fields{"int": config.C.Paystack.LandFees}).Errorf("paystack: strconv.Atoi fees")
			return nil, err
		}
		fees = i
	}
	input.Amount = fees * 100

	payload, err := json.Marshal(input)
	if err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"input": input}).Errorf("paystack: json.Marshal")
		return nil, err
	}

	req, err := http.NewRequest("POST", chargeApi, bytes.NewBuffer(payload))
	if err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"req": payload}).Errorf("paystack: http.NewRequest")
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.C.Paystack.SecretKey)

	res, err := p.http.Do(req)
	if err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"req": payload}).Errorf("paystack: p.http.Do")
		return nil, err
	}

	if err := util.DecodeHttp(res.Body, &chargeResponse); err != nil {
		p.log.WithError(err).Errorf("paystack: decode")
		return nil, err
	}

	go func() {
		ctx := context.Background()
		args := sql.CreatePaymentParams{
			Email:        input.Email,
			ReferenceID:  chargeResponse.Data.Reference,
			Status:       chargeResponse.Data.Status,
			Reason:       input.Reason,
			Amount:       int32(input.Amount),
			Currency:     input.Currency,
			OnboardingID: uuid.NullUUID{UUID: paymentFor, Valid: true},
		}
		_, err := controller.GetPaymentController().CreatePayment(ctx, args)
		if err != nil {
			p.log.WithError(err).Errorf("paystack: CreatePayment async")
			return
		}
	}()

	return chargeResponse, nil
}

func (p *paystackClient) HappyPaystack(ctx context.Context, input *PaystackWebhook) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	payment, err := p.sqlStore.GetPaymentByReferenceID(ctx, input.Data.Reference)
	if err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"reference_id": input.Data.Reference}).Errorf("paystack: GetPaymentByReferenceID")
		return err
	}

	payload := model.PaymentUpdate{
		ReferenceID: input.Data.Reference,
		Status:      input.Data.Status,
		Email:       payment.Email,
	}
	b, err := util.EncodeJson(payload)
	if err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"payload": payload}).Errorf("paystack: EncodeJson webhook payload")
		return err
	}

	// update payment status in db
	go func() {
		ctx := context.Background()
		args := sql.UpdatePaymentStatusParams{
			ReferenceID: input.Data.Reference,
			Status:      input.Data.Status,
		}
		_, err := p.sqlStore.UpdatePaymentStatus(ctx, args)
		if err != nil {
			p.log.WithError(err).Errorf("paystack: UpdatePaymentStatus async")
			return
		}
	}()

	if err := p.redis.Publish(ctx, cache.PAYMENT_UPDATED_CHANNEL, b).Err(); err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"payload": payload}).Errorf("paystack: redis.Publish payment update")
		return err
	}

	return nil
}
