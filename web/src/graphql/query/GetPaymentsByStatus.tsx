import { gql } from "@apollo/client"

export default gql`
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
          verified
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
      reference_id
    }
  }
`
