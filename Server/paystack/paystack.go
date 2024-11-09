package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"sync"

	"github.com/elc49/copod/config"
	"github.com/elc49/copod/logger"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/elc49/copod/util"
	"github.com/sirupsen/logrus"
)

var p *paystackClient

type Paystack interface {
	ChargeMpesa(context.Context, MpesaCharge) (*MpesaChargeResponse, error)
}

var _ Paystack = (*paystackClient)(nil)

type paystackClient struct {
	log      *logrus.Logger
	sqlStore *sql.Queries
	mu       sync.Mutex
	http     *http.Client
}

func New(sqlStore *sql.Queries) {
	log := logger.GetLogger()
	p = &paystackClient{log, sqlStore, sync.Mutex{}, &http.Client{}}
}

func GetPaystackService() Paystack {
	return p
}

func (p *paystackClient) ChargeMpesa(ctx context.Context, input MpesaCharge) (*MpesaChargeResponse, error) {
	var chargeResponse *MpesaChargeResponse
	chargeApi := config.C.Paystack.BaseApi + "/charge"
	input.Provider.Provider = "mpesa"
	if !config.IsProd() {
		input.Provider.Phone = config.C.Paystack.MobileTestAccount
	}

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
	req.Header.Add("Authorization", config.C.Paystack.SecretKey)

	res, err := p.http.Do(req)
	if err != nil {
		p.log.WithError(err).WithFields(logrus.Fields{"req": payload}).Errorf("paystack: p.http.Do")
		return nil, err
	}

	if err := util.DecodeHttp(res.Body, chargeResponse); err != nil {
		p.log.WithError(err).Errorf("paystack: decode")
		return nil, err
	}

	return chargeResponse, nil
}
