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
	var err error

	t.Run("create_support_doc", func(t *testing.T) {

		assert.Nil(t, err)
		assert.Equal(t, supportdoc.URL, docUri)
	})

	t.Run("should_only_be_one_supporting_doc", func(t *testing.T) {
		e, err := sc.CreateSupportingDoc(ctx, sql.CreateSupportDocParams{
			Url:   docUri,
			Email: email,
		})

		assert.Nil(t, err)
		assert.Equal(t, supportdoc.ID, e.ID)
	})

	t.Run("get_support_docs_by_verification", func(t *testing.T) {
		docs, err := sc.GetSupportingDocsByVerification(ctx, model.VerificationOnboarding)

		assert.Nil(t, err)
		assert.True(t, len(docs) > 0)
	})

	t.Run("get_supporting_doc_by_id", func(t *testing.T) {
		doc, err := sc.GetSupportingDocByID(ctx, supportdoc.ID)

		assert.Nil(t, err)
		assert.Equal(t, doc.Email, email)
	})

	t.Run("should_update_supporting_doc_by_email", func(t *testing.T) {
		args := sql.UpdateSupportDocByIDParams{
			Verification: model.VerificationVerified.String(),
			ID:           supportdoc.ID,
		}
		doc, err := sc.UpdateSupportDocByID(ctx, args)

		assert.Nil(t, err)
		assert.Equal(t, doc.Verified, model.VerificationVerified)
	})

	t.Run("udpate_support_doc_verification_by_id", func(t *testing.T) {
		u, err := sc.UpdateSupportDocVerificationByID(ctx, email, sql.UpdateSupportDocVerificationByIDParams{
			ID:           supportdoc.ID,
			Verification: model.VerificationVerified.String(),
		})

		assert.Nil(t, err)
		assert.Equal(t, u.Verified, model.VerificationVerified)
	})
}
