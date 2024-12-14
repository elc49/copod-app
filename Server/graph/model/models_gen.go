// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type CreateOnboardingInput struct {
	Email             string `json:"email"`
	TitleURL          string `json:"titleUrl"`
	SupportdocURL     string `json:"supportdocUrl"`
	DisplayPictureURL string `json:"displayPictureUrl"`
}

type DisplayPicture struct {
	ID        uuid.UUID    `json:"id"`
	URL       string       `json:"url"`
	Email     string       `json:"email"`
	Verified  Verification `json:"verified"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type GetOnboardingByEmailAndVerificationInput struct {
	Email        string       `json:"email"`
	Verification Verification `json:"verification"`
}

type GetUserLandsInput struct {
	Email        string       `json:"email"`
	Verification Verification `json:"verification"`
}

type Mutation struct {
}

type Onboarding struct {
	ID               uuid.UUID       `json:"id"`
	TitleID          uuid.UUID       `json:"titleId"`
	Title            *Title          `json:"title"`
	SupportDocID     uuid.UUID       `json:"supportDocId"`
	SupportingDoc    *SupportingDoc  `json:"supportingDoc"`
	DisplayPictureID uuid.UUID       `json:"displayPictureId"`
	DisplayPicture   *DisplayPicture `json:"displayPicture"`
	Verification     Verification    `json:"verification"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

type PayWithMpesaInput struct {
	Reason     PaymentReason `json:"reason"`
	Phone      string        `json:"phone"`
	Email      string        `json:"email"`
	Currency   string        `json:"currency"`
	PaymentFor uuid.UUID     `json:"paymentFor"`
}

type Payment struct {
	ID           uuid.UUID   `json:"id"`
	ReferenceID  string      `json:"reference_id"`
	Status       string      `json:"status"`
	Email        string      `json:"email"`
	Onboarding   *Onboarding `json:"onboarding,omitempty"`
	OnboardingID uuid.UUID   `json:"onboarding_id"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

type PaymentUpdate struct {
	ReferenceID string `json:"referenceId"`
	Status      string `json:"status"`
	Email       string `json:"email"`
}

type Query struct {
}

type Subscription struct {
}

type SupportingDoc struct {
	ID        uuid.UUID    `json:"id"`
	URL       string       `json:"url"`
	Email     string       `json:"email"`
	Verified  Verification `json:"verified"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type Title struct {
	ID           uuid.UUID    `json:"id"`
	URL          string       `json:"url"`
	Verified     Verification `json:"verified"`
	SupportDocID uuid.UUID    `json:"support_doc_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

type UpdateOnboardingStatusInput struct {
	OnboardingID uuid.UUID    `json:"onboardingId"`
	Verification Verification `json:"verification"`
}

type UpdateTitleVerificationByIDInput struct {
	TitleID      uuid.UUID    `json:"titleId"`
	Verification Verification `json:"verification"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	Firstname *string   `json:"firstname,omitempty"`
	Lastname  *string   `json:"lastname,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
