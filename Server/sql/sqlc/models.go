// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sql

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DisplayPicture struct {
	ID           uuid.UUID `json:"id"`
	Url          string    `json:"url"`
	Verification string    `json:"verification"`
	Email        string    `json:"email"`
	SupportDocID uuid.UUID `json:"support_doc_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type EarlySignup struct {
	ID        uuid.UUID    `json:"id"`
	Email     string       `json:"email"`
	Onboarded sql.NullTime `json:"onboarded"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Onboarding struct {
	ID               uuid.UUID      `json:"id"`
	TitleID          uuid.UUID      `json:"title_id"`
	SupportDocID     uuid.UUID      `json:"support_doc_id"`
	DisplayPictureID uuid.UUID      `json:"display_picture_id"`
	Email            string         `json:"email"`
	Verification     string         `json:"verification"`
	PaymentStatus    sql.NullString `json:"payment_status"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

type Payment struct {
	ID           uuid.UUID     `json:"id"`
	Email        string        `json:"email"`
	Amount       int32         `json:"amount"`
	Currency     string        `json:"currency"`
	Reason       string        `json:"reason"`
	Status       string        `json:"status"`
	ReferenceID  string        `json:"reference_id"`
	OnboardingID uuid.NullUUID `json:"onboarding_id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
}

type SupportDoc struct {
	ID           uuid.UUID `json:"id"`
	Url          string    `json:"url"`
	Verification string    `json:"verification"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TitleDeed struct {
	ID           uuid.UUID      `json:"id"`
	Url          string         `json:"url"`
	Title        sql.NullString `json:"title"`
	Verification string         `json:"verification"`
	Email        string         `json:"email"`
	SupportDocID uuid.UUID      `json:"support_doc_id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
