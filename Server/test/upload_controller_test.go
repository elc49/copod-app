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
		p, err := uc.CreateUpload(context.Background(), sql.CreateUploadParams{
			Type:          model.DocLandTitle.String(),
			Uri:           docUri,
			WalletAddress: superUserWallet,
		})

		assert.Equal(t, p.URI, docUri)
		assert.Nil(t, err)
	})
}
