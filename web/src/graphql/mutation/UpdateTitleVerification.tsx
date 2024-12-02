import { gql } from "@apollo/client"

export const UPDATE_TITLE_VERIFICATION = gql`
  mutation UpdateTitleVerification($input: UpdateTitleVerificationInput!) {
    updateTitleVerification(input: $input) {
      id
      verified
    }
  }
`
