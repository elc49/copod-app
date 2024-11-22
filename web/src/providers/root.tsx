"use client";

import { PropsWithChildren } from "react";
import { WalletProvider } from "./wallet";

const Providers = ({ children }: PropsWithChildren) => {
  return (
    <WalletProvider>
      {children}
    </WalletProvider>
  )
}

export { Providers }
