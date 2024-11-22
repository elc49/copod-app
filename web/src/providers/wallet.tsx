"use client";

import { createContext, PropsWithChildren, useEffect, useState } from "react";
import { MetaMaskSDK, SDKProvider } from "@metamask/sdk";
import { type Address, type Chain, type WalletClient, custom, createWalletClient } from "viem";
import "viem/window";

const infuraAPIKey = process.env.NEXT_PUBLIC_INFURA_API_KEY

interface IWalletContext {
  account: Address | undefined
  connected: boolean
  chain: Chain | undefined
  connectWallet: () => void
  walletClient: WalletClient | undefined
  connecting: boolean
}

const WalletContext = createContext<IWalletContext>({
  account: undefined,
  connected: false,
  chain: undefined,
  connectWallet: () => {},
  walletClient: undefined,
  connecting: false,
})

const WalletProvider = ({ children }: PropsWithChildren) => {
  const [sdk, setSdk] = useState<MetaMaskSDK>()
  const [provider, setProvider] = useState<SDKProvider>()
  const [account, setAccount] = useState<Address>()
  const [connected, setConnected] = useState<boolean>(false)
  const [chain, setChain] = useState<Chain>()
  const [walletClient, setWalletClient] = useState<WalletClient>()
  const [connecting, setConnecting] = useState<boolean>(false)

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
      setChain(chain as Chain)
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
      setAccount((accounts as Address[])?.[0])
      setConnected(true)
    }

    const onConnect = (_connectInfo: any) => {
      console.log("Connecting/\n", _connectInfo)
      setConnected(true)
      setChain(_connectInfo.chainId as Chain)
    }

    const onDisconnect = (error: unknown) => {
      console.log("Disconnecting\n", error)
      setConnected(false)
      setChain(undefined)
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

  const connect = async () => {
    if (!provider) {
      throw new Error(`No window.ethereum in the current browser session`)
    }

    try {
      setConnecting(true)
      setWalletClient(
        createWalletClient({
          chain: chain,
          transport: custom(provider),
        })
      )
      const address = await walletClient?.requestAddresses()
      setAccount(address?.[0])
    } catch (e) {
      console.error(e)
    }
  }

  return (
    <WalletContext.Provider
      value={{
        chain: chain,
        account: account,
        connected: connected,
        connectWallet: connect,
        walletClient: walletClient,
        connecting: connecting,
      }}
    >
      {children}
    </WalletContext.Provider>
  )
}

export { WalletContext, WalletProvider }
