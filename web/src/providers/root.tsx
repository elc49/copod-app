"use client";

import { PropsWithChildren } from "react";
import { WalletProvider } from "./wallet";
import { ApolloProvider } from "./apollo";
import { ChakraUIProvider } from "./chakra-ui";
import { Toaster } from "@/components/ui/toaster";

const Providers = ({ children }: PropsWithChildren) => {
  return (
    <ChakraUIProvider>
      <WalletProvider>
        <ApolloProvider>
          <Toaster />
          {children}
         </ApolloProvider>
      </WalletProvider>
    </ChakraUIProvider>
  )
}

export { Providers }
