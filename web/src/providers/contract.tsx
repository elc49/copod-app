"use client";

import { createContext, PropsWithChildren, useContext } from "react";
import { WalletContext } from "@/providers/wallet";

const ContractContext = createContext({})

const ContractProvider = ({ children }: PropsWithChildren) => {
  const {} = useContext(WalletContext)

  return (
    <ContractContext.Provider
      value={{
      }}
    >
      {children}
    </ContractContext.Provider>
  )
}

export { ContractContext, ContractProvider }
