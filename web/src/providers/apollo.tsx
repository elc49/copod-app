"use client";

import { PropsWithChildren } from "react";
import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  from,
  type NormalizedCacheObject,
  ApolloProvider as ApolloClientProvider,
} from "@apollo/client";
import { RetryLink } from "@apollo/client/link/retry";

const api = process.env.NEXT_PUBLIC_COPOD_API

const httpLink = new HttpLink({
  uri: api,
})

const createClient = (): ApolloClient<NormalizedCacheObject> => {
  // caching
  const cache = new InMemoryCache({})
  // error handling
  const retryLink = new RetryLink({
    delay: {
      initial: 300,
      jitter: true,
    },
    attempts: {
      max: 2,
      retryIf: (error) => !!error,
    },
  })

  return new ApolloClient({
    link: from([
      // error link
      retryLink,
      // api
      httpLink,
    ]),
    cache,
    defaultOptions: {
      watchQuery: {
        fetchPolicy: "cache-and-network" as const,
      },
    },
  })
}

export const ApolloProvider = ({ children }: PropsWithChildren) => {
  const client = createClient()

  return (
    <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
  )
}
