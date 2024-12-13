import { gql } from "@apollo/client"

export const GET_PAYMENTS_BY_STATUS = gql`
  query GetPaymentsByStatus($status: PaymentStatus!) {
    getPaymentsByStatus(status: $status) {
      id
      email
      status
      onboarding {
        verification
        titleId
        title {
          id
          url
        }
        displayPictureId
        displayPicture {
          id
          url
        }
        supportDocId
        supportingDoc {
          id
          url
        }
      }
      onboarding_id
      verified
      reference_id
    }
  }
`
