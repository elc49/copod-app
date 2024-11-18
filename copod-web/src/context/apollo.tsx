"use client";

import { PropsWithChildren } from "react";
import { ApolloProvider as ApolloClientProvider } from "@apollo/client";
import { createClient } from "@/apollo/createClient";

const ApolloProvider = ({ children }: PropsWithChildren) => {
  const client = createClient()

  return <ApolloClientProvider client={client}>{children}</ApolloClientProvider>
}

export default ApolloProvider
