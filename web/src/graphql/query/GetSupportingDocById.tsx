import { gql } from "@apollo/client"

export const GET_SUPPORTING_DOC_BY_ID = gql`
  query GetSupportingDocById($id: UUID!) {
    getSupportingDocById(id: $id) {
      id
      govt_id
      verified
    }
  }
`
