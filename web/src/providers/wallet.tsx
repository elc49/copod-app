"use client";

import { PropsWithChildren } from "react";
import { MetaMaskProvider } from "@metamask/sdk-react";

const infuraAPIKey = process.env.NEXT_PUBLIC_INFURA_API_KEY

const WalletProvider = ({ children }: PropsWithChildren) => {
  return (
    <MetaMaskProvider
      sdkOptions={{
        dappMetadata: {
          name: "Copod",
          url: "http://localhost:3000",
        },
        infuraAPIKey,
      }}
    >
      {children}
    </MetaMaskProvider>
  )
}

export { WalletProvider }
