"use client";

import { PropsWithChildren } from "react";
import { WalletProvider } from "./wallet";
import { ApolloProvider } from "./apollo";
import { ChakraUIProvider } from "./chakra-ui";

const Providers = ({ children }: PropsWithChildren) => {
  return (
    <WalletProvider>
      <ApolloProvider>
        <ChakraUIProvider>
          {children}
        </ChakraUIProvider>
      </ApolloProvider>
    </WalletProvider>
  )
}

export { Providers }
