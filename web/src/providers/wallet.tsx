"use client";

import { createContext, PropsWithChildren, useEffect, useState } from "react";
import { MetaMaskSDK, SDKProvider } from "@metamask/sdk";

const infuraAPIKey = process.env.NEXT_PUBLIC_INFURA_API_KEY

interface IWalletContext {
  account: string
  connected: boolean
  chain: string
  metamaskSdk: MetaMaskSDK
  metamaskProvider: SDKProvider
  connectWallet: () => void
}

const WalletContext = createContext<IWalletContext>({
  account: "",
  connected: false,
  chain: "",
  metamaskSdk: {},
  metamaskProvider: {},
  connectWallet: () => {},
})

const WalletProvider = ({ children }: PropsWithChildren) => {
  const [sdk, setSdk] = useState<MetaMaskSDK>()
  const [provider, setProvider] = useState<SDKProvider>()
  const [account, setAccount] = useState<string>("")
  const [connected, setConnected] = useState<boolean>(false)
  const [chain, setChain] = useState("")

  useEffect(() => {
    const doAsync = async () => {
      const clientSDK = new MetaMaskSDK({
        dappMetadata: {
          name: "Copod",
          url: "http://localhost:3000",
          iconUrl: "http://localhost:3000/favicon.ico",
        },
        infuraAPIKey,
        checkInstallationImmediately: true,
      })
      await clientSDK.init()
      setSdk(clientSDK)
      setProvider(clientSDK.getProvider())
    }
    doAsync()
  }, [])

  useEffect(() => {
    if (!sdk || !provider) {
      return;
    }

    console.log("Setting active provider object")
    if (provider.getSelectedAddress()) {
      console.log("Setting account from provider")
      setAccount(provider.getSelectedAddress() ?? "")
      setConnected(true);
    } else {
      setConnected(false)
    }

    const onChainChanged = (chain: unknown) => {
      console.log("Blockchain id changed")
      setChain(chain as string)
    }

    const onInitialized = () => {
      console.log("Initialized")
      setConnected(true)
      if (provider.getSelectedAddress()) {
        setAccount(provider.getSelectedAddress() ?? "")
      }

      if (provider.chainId()) {
        setChain(provider.chainId())
      }
    }

    const onAccountChanged = (accounts: unknown) => {
      console.log("Account changed")
      setAccount((accounts as string[])?.[0])
      setConnected(true)
    }

    const onConnect = (_connectInfo: any) => {
      console.log("Connecting/\n", _connectInfo)
      setConnected(true)
      setChain(_connectInfo.chainId as string)
    }

    const onDisconnect = (error: unknown) => {
      console.log("Disconnecting\n", error)
      setConnected(false)
      setChain("")
    }

    provider.on("accountsChanged", onAccountChanged)
    provider.on("chainChanged", onChainChanged)
    provider.on("_initialized", onInitialized)
    provider.on("connect", onConnect)
    provider.on("disconnect", onDisconnect)
    
    return () => {
      console.log("Clean up window.ethereum events");
      provider.removeListener('chainChanged', onChainChanged);
      provider.removeListener('_initialized', onInitialized);
      provider.removeListener('accountsChanged', onAccountChanged);
      provider.removeListener('connect', onConnect);
      provider.removeListener('disconnect', onDisconnect);
    }
  }, [provider])

  const connect = () => {
    if (!provider) {
      throw new Error(`No window.ethereum in the current browser session`)
    }

    provider.request({ method: "eth_requestAccounts", params: [] })
      .then((accounts) => {
        console.log("Accounts connected")
        setAccount((accounts as string[])?.[0])
      })
      .catch((error) => console.error("eth_requestAccounts", error))
  }

  return (
    <WalletContext.Provider
      value={{
        chain: chain,
        account: account,
        connected: connected,
        metamaskSdk: sdk,
        metamaskProvider: provider,
        connectWallet: connect,
      }}
    >
      {children}
    </WalletContext.Provider>
  )
}

export { WalletContext, WalletProvider }
