import { gql } from "@apollo/client";

export const GET_PAYMENT_DETAILS_BY_ID = gql`
  query GetPaymentDetailsById($id: UUID!) {
    getPaymentDetailsById(id: $id) {
      id
      title {
        id
        title
      }
    }
  }
`
