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
	t.Run("create_upload", func(t *testing.T) {
		uc := controller.GetUploadController()
		uploads := []sql.CreateUploadParams{
			{
				Type:          model.DocLandTitle.String(),
				Uri:           docUri,
				WalletAddress: superUserWallet,
			},
			{
				Type:          model.DocLandTitle.String(),
				Uri:           docUri,
				WalletAddress: superUserWallet,
			},
		}
		for _, i := range uploads {
			p, err := uc.CreateUpload(context.Background(), sql.CreateUploadParams{
				Type:          i.Type,
				Uri:           i.Uri,
				WalletAddress: i.WalletAddress,
			})

			assert.Equal(t, p.URI, i.Uri)
			assert.Nil(t, err)
		}
	})
}
