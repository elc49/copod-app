// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type DocUploadInput struct {
	URL           string `json:"url"`
	Email         string `json:"email"`
	WalletAddress string `json:"walletAddress"`
}

type Land struct {
	ID            uuid.UUID    `json:"id"`
	Title         string       `json:"title"`
	Size          int          `json:"size"`
	Symbol        string       `json:"symbol"`
	Town          *string      `json:"town,omitempty"`
	TitleDocument string       `json:"titleDocument"`
	Verified      Verification `json:"verified"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
}

type Mutation struct {
}

type PayWithMpesaInput struct {
	Reason        PaymentReason `json:"reason"`
	Phone         string        `json:"phone"`
	Email         string        `json:"email"`
	WalletAddress string        `json:"walletAddress"`
	Currency      string        `json:"currency"`
	PaymentFor    uuid.UUID     `json:"paymentFor"`
}

type Payment struct {
	ID            uuid.UUID `json:"id"`
	ReferenceID   string    `json:"reference_id"`
	Status        string    `json:"status"`
	Title         *Title    `json:"title,omitempty"`
	TitleID       uuid.UUID `json:"title_id"`
	WalletAddress string    `json:"wallet_address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentUpdate struct {
	ReferenceID   string `json:"referenceId"`
	Status        string `json:"status"`
	WalletAddress string `json:"walletAddress"`
}

type Query struct {
}

type Subscription struct {
}

type SupportingDoc struct {
	ID        uuid.UUID    `json:"id"`
	GovtID    string       `json:"govt_id"`
	Verified  Verification `json:"verified"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Title struct {
	ID        uuid.UUID    `json:"id"`
	Title     string       `json:"title"`
	Verified  Verification `json:"verified"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type User struct {
	ID            uuid.UUID `json:"id"`
	Firstname     *string   `json:"firstname,omitempty"`
	Lastname      *string   `json:"lastname,omitempty"`
	Email         string    `json:"email"`
	WalletAddress string    `json:"wallet_address"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaidFor string

const (
	PaidForPaid    PaidFor = "PAID"
	PaidForNotPaid PaidFor = "NOT_PAID"
)

var AllPaidFor = []PaidFor{
	PaidForPaid,
	PaidForNotPaid,
}

func (e PaidFor) IsValid() bool {
	switch e {
	case PaidForPaid, PaidForNotPaid:
		return true
	}
	return false
}

func (e PaidFor) String() string {
	return string(e)
}

func (e *PaidFor) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaidFor(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaidFor", str)
	}
	return nil
}

func (e PaidFor) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaymentReason string

const (
	PaymentReasonLandRegistry PaymentReason = "LAND_REGISTRY"
)

var AllPaymentReason = []PaymentReason{
	PaymentReasonLandRegistry,
}

func (e PaymentReason) IsValid() bool {
	switch e {
	case PaymentReasonLandRegistry:
		return true
	}
	return false
}

func (e PaymentReason) String() string {
	return string(e)
}

func (e *PaymentReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentReason", str)
	}
	return nil
}

func (e PaymentReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaymentStatus string

const (
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusTimeout PaymentStatus = "timeout"
)

var AllPaymentStatus = []PaymentStatus{
	PaymentStatusSuccess,
	PaymentStatusFailed,
	PaymentStatusPending,
	PaymentStatusTimeout,
}

func (e PaymentStatus) IsValid() bool {
	switch e {
	case PaymentStatusSuccess, PaymentStatusFailed, PaymentStatusPending, PaymentStatusTimeout:
		return true
	}
	return false
}

func (e PaymentStatus) String() string {
	return string(e)
}

func (e *PaymentStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentStatus", str)
	}
	return nil
}

func (e PaymentStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Verification string

const (
	VerificationOnboarding Verification = "ONBOARDING"
	VerificationVerified   Verification = "VERIFIED"
	VerificationRejected   Verification = "REJECTED"
)

var AllVerification = []Verification{
	VerificationOnboarding,
	VerificationVerified,
	VerificationRejected,
}

func (e Verification) IsValid() bool {
	switch e {
	case VerificationOnboarding, VerificationVerified, VerificationRejected:
		return true
	}
	return false
}

func (e Verification) String() string {
	return string(e)
}

func (e *Verification) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Verification(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Verification", str)
	}
	return nil
}

func (e Verification) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
