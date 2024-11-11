package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Payment_Controller(t *testing.T) {
	ctx := context.Background()
	pc := controller.GetPaymentController()

	t.Run("create_payment", func(t *testing.T) {
		uc := controller.GetUploadController()
		p, err := uc.CreateUpload(context.Background(), sql.CreateUploadParams{
			Type:          model.DocLandTitle.String(),
			Uri:           docUri,
			WalletAddress: superUserWallet,
		})
		assert.Nil(t, err)

		b, err := pc.CreatePayment(ctx, sql.CreatePaymentParams{
			Email:       superUserEmail,
			ReferenceID: RandomStringByLength(10),
			Status:      "success",
			Reason:      model.PaymentReasonLandRegistration.String(),
			Amount:      1500,
			Currency:    "KES",
			UploadID:    uuid.NullUUID{UUID: p.ID, Valid: true},
		})

		assert.Nil(t, err)
		assert.True(t, *b)
	})
}
