package test

import (
	"context"
	db "database/sql"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_Upload_Controller(t *testing.T) {
	t.Run("create_upload", func(t *testing.T) {
		pc := controller.GetUploadController()
		p, err := pc.CreateUpload(context.Background(), sql.CreateUploadParams{
			Type:     model.DocLandTitle.String(),
			TitleDoc: db.NullString{String: docUri, Valid: true},
			GovtID:   db.NullString{String: docUri, Valid: true},
			Email:    superUserEmail,
		})

		assert.Equal(t, *p.TitleDoc, docUri)
		assert.Nil(t, err)
	})
}
