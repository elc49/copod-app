import { gql } from "@apollo/client";

export const GET_SUPPORTING_DOCS_BY_VERIFICATION = gql`
  query GetSupportingDocsByVerification($verification: Verification!) {
    getSupportingDocsByVerification(verification: $verification) {
      id
      govt_id
      verified
    }
  }
`
