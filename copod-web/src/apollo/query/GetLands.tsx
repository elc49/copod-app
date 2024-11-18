import { gql } from "@apollo/client";

export const GetLands = gql`
  query GetLands {
    getLands {
      id
    }
  }
`
