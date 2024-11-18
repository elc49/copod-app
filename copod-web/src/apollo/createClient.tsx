import {
  ApolloClient,
  InMemoryCache,
  from,
  type NormalizedCacheObject,
} from "@apollo/client";
import { RetryLink } from "@apollo/client/link/retry";
import { httpLink } from "@/apollo/link";

export const createClient = (): ApolloClient<NormalizedCacheObject> => {
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
      // error retry link
      retryLink,
      // api link
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
