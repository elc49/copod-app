import { gql } from "@apollo/client"

export default gql`
  query GetSupportingDocById($id: UUID!) {
    getSupportingDocById(id: $id) {
      id
      govt_id
      verified
    }
  }
`
