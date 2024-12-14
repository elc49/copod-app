import { gql } from "@apollo/client"

export default gql`
  query GetTitleById($id: UUID!) {
    getTitleById(id: $id) {
      id
      url
      verified
    }
  }
`
