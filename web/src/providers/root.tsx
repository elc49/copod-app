"use client";

import { PropsWithChildren } from "react";
import { WalletProvider } from "./wallet";
import Header from "@/components/header";

const Providers = ({ children }: PropsWithChildren) => {
  return (
    <WalletProvider>
      <Header />
      {children}
    </WalletProvider>
  )
}

export { Providers }
