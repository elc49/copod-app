package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Upload_Controller(t *testing.T) {
	defer func() {
		q.ClearTestUploads(context.Background())
	}()

	t.Run("create_upload", func(t *testing.T) {
		uc := controller.UC
		u, _ := uc.CreateUser(context.Background(), sql.CreateUserParams{
			Email:         fmt.Sprintf("%s@em.com", RandomStringByLength(4)),
			WalletAddress: fmt.Sprintf("0x41eD3Ce6DC13fD4F67Eb715f5c3B105B%s", RandomStringByLength(8)),
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
