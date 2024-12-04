package test

import (
	"context"
	"testing"

	"github.com/elc49/copod/controller"
	"github.com/elc49/copod/graph/model"
	sql "github.com/elc49/copod/sql/sqlc"
	"github.com/stretchr/testify/assert"
)

func Test_SupportDoc_Controller(t *testing.T) {
	sc := controller.GetSupportingDocController()
	ctx := context.Background()
	var s *model.SupportingDoc
	var err error

	t.Run("create_support_doc", func(t *testing.T) {
		args := sql.CreateSupportDocParams{
			GovtID: docUri,
			Email:  email,
		}
		s, err = sc.CreateSupportingDoc(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, s.GovtID, docUri)
	})

	t.Run("should_only_be_one_supporting_doc", func(t *testing.T) {
		e, err := sc.CreateSupportingDoc(ctx, sql.CreateSupportDocParams{
			GovtID: docUri,
			Email:  email,
		})

		assert.Nil(t, err)
		assert.Equal(t, s.ID, e.ID)
	})

	t.Run("get_support_docs_by_verification", func(t *testing.T) {
		docs, err := sc.GetSupportingDocsByVerification(ctx, model.VerificationOnboarding)

		assert.Nil(t, err)
		assert.True(t, len(docs) > 0)
	})

	t.Run("get_supporting_doc_by_id", func(t *testing.T) {
		doc, err := sc.GetSupportingDocByID(ctx, s.ID)

		assert.Nil(t, err)
		assert.Equal(t, doc.GovtID, docUri)
	})

	t.Run("should_update_supporting_doc_by_email", func(t *testing.T) {
		args := sql.UpdateUserSupportDocByEmailParams{
			GovtID:       "3849#",
			Verification: model.VerificationVerified.String(),
			Email:        email,
		}
		doc, err := sc.UpdateSupportingDocByEmail(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, doc.Verified, model.VerificationVerified)
	})
}
