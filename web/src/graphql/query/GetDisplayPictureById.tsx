import { gql } from "@apollo/client"

export default gql`
  query GetDisplayPictureById($id: UUID!) {
    getDisplayPictureById(id: $id) {
      id
      url
      verified
    }
  }
`
