"use client";

import { PropsWithChildren } from "react";
import { WalletProvider } from "./wallet";
import { ApolloProvider } from "./apollo";

const Providers = ({ children }: PropsWithChildren) => {
  return (
    <WalletProvider>
      <ApolloProvider>
        {children}
      </ApolloProvider>
    </WalletProvider>
  )
}

export { Providers }
