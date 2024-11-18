import { HttpLink } from "@apollo/client"

const api = process.env.NEXT_PUBLIC_COPOD_API

export const httpLink = new HttpLink({
  uri: api,
})
