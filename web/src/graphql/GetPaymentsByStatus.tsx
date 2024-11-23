import { gql } from "@apollo/client"

export const GET_PAYMENTS_BY_STATUS = gql`
  query GetPaymentsByStatus($status: PaymentStatus!) {
    getPaymentsByStatus(status: $status) {
      id
    }
  }
`
