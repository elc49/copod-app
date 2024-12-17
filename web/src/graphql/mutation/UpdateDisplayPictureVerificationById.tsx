import { gql } from "@apollo/client"

export default gql`
  mutation UpdateDisplayPictureVerificationById($input: UpdateDisplayPictureVerificationByIdInput!) {
    updateDisplayPictureVerificationById(input: $input) {
      id
    }
  }
`
