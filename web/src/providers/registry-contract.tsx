"use client";

import { createContext, PropsWithChildren, useContext, useEffect, useState } from "react";
import { getContract } from "viem";
import { WalletContext } from "@/providers/wallet";

interface IRegistryContractContext {
  registryContract: any
}

interface Deployments {
  readonly "Registry#Registry": string
}

const RegistryContractContext = createContext<IRegistryContractContext>({
  registryContract: undefined
})

const RegistryContractProvider = ({ children }: PropsWithChildren) => {
  const { wallet } = useContext(WalletContext)
  const [deployments, setDeployments] = useState<Deployments>()
  const [abi, setAbi] = useState<any>()
  const [registryContract, setRegistryContract] = useState<any>()

  useEffect(() => {
    async function loadDeployments() {
      try {
        const deployments = await import("../../../SmartContract/ignition/deployments/chain-11155420/deployed_addresses.json")
        const abi = await import("../../../SmartContract/ignition/deployments/chain-11155420/artifacts/Registry#Registry.json")
        setAbi(abi.abi)
        setDeployments(deployments.default)
      } catch (e) {
        console.error(e)
      }
    }

    loadDeployments()
  }, [])

  useEffect(() => {
    const contract = getContract({
      address: deployments?.["Registry#Registry"],
      abi: abi ?? [],
      client: wallet,
    })
    setRegistryContract(contract)
  }, [abi, deployments])

  return (
    <RegistryContractContext.Provider
      value={{
        registryContract,
      }}
    >
      {children}
    </RegistryContractContext.Provider>
  )
}

export { RegistryContractContext, RegistryContractProvider }
