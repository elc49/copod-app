import { gql } from "@apollo/client";

export default gql`
  mutation UpdateTitleVerificationById($input: UpdateTitleVerificationByIdInput!) {
    updateTitleVerificationById(input: $input) {
      id
      verified
    }
  }
`
