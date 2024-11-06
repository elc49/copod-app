package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Upload_Controller(t *testing.T) {
	defer func() {
		ctx := context.Background()
		q.ClearTestUploads(ctx)
		q.ClearTestUsers(ctx)
	}()

	t.Run("create_upload", func(t *testing.T) {
		uc := controller.UC
		u, _ := uc.CreateUser(context.Background(), sql.CreateUserParams{
			Email:         RandomEmailAddress(),
			WalletAddress: RandomWalletAddress(),
			GovtID:        RandomGovtID(),
		})

		pc := controller.PC
		p, err := pc.CreateUpload(context.Background(), sql.CreateUploadParams{
			Type:          model.DocLandTitle.String(),
			Uri:           docUri,
			WalletAddress: u.WalletAddress,
		})

		assert.Equal(t, p.URI, docUri)
		assert.Nil(t, err)
	})

	t.Run("get_upload", func(t *testing.T) {
	})
}
