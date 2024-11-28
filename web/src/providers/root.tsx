"use client";

import { PropsWithChildren } from "react";
import { WalletProvider } from "./wallet";
import { ApolloProvider } from "./apollo";
import { ChakraUIProvider } from "./chakra-ui";
import { RegistryContractProvider } from "./registry-contract";

const Providers = ({ children }: PropsWithChildren) => {
  return (
    <ChakraUIProvider>
      <WalletProvider>
        <ApolloProvider>
          <RegistryContractProvider>
            {children}
          </RegistryContractProvider>
        </ApolloProvider>
      </WalletProvider>
    </ChakraUIProvider>
  )
}

export { Providers }
