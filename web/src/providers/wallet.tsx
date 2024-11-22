"use client";

import { createContext, PropsWithChildren, useEffect, useState } from "react";
import { MetaMaskSDK, SDKProvider } from "@metamask/sdk";
import { type Address, type Chain, type WalletClient, custom, createWalletClient } from "viem";
import "viem/window";

const infuraAPIKey = process.env.NEXT_PUBLIC_INFURA_API_KEY

interface IWalletContext {
  account: Address | undefined
  chain: Chain | undefined
  connect: () => void
  disconnect: () => void
  wallet: WalletClient | undefined
  connecting: boolean
}

const WalletContext = createContext<IWalletContext>({
  account: undefined,
  chain: undefined,
  connect: () => {},
  disconnect: () => {},
  wallet: undefined,
  connecting: false,
})

const WalletProvider = ({ children }: PropsWithChildren) => {
  const [sdk, setSdk] = useState<MetaMaskSDK>()
  const [provider, setProvider] = useState<SDKProvider>()
  const [account, setAccount] = useState<Address>()
  const [chain, setChain] = useState<Chain>()
  const [wallet, setWallet] = useState<WalletClient>()
  const [connecting, setConnecting] = useState<boolean>(false)

  useEffect(() => {
    const doAsync = async () => {
      try {
        setConnecting(true)
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
        setChain(await clientSDK.getProvider().request({ method: "eth_chainId" }))
      } catch (e) {
        console.error(e)
      } finally {
        setConnecting(false)
      }
    }
    doAsync()
  }, [])

  useEffect(() => {
    if (!sdk || !provider) {
      return;
    }

    console.log("Setting account from active provider")
    try {
      setConnecting(true)
      setWallet(
        createWalletClient({
          chain: chain,
          transport: custom(provider),
        })
      )
      wallet?.requestAddresses()
      .then((accounts) => {
        setAccount(accounts?.[0])
      })
      .catch((e) => console.error(e))
    } catch (e) {
      console.error(e)
    } finally {
      setConnecting(false)
    }

    const onChainChanged = (chain: unknown) => {
      console.log("Blockchain id changed")
      setChain(chain as Chain)
    }

    const onAccountChanged = (accounts: unknown) => {
      console.log("Account changed")
      setAccount((accounts as Address[])?.[0])
    }

    const onDisconnect = (error: unknown) => {
      console.log("Disconnecting\n", error)
      setAccount(undefined)
      setProvider(null)
      setChain(undefined)
    }

    provider.on("accountsChanged", onAccountChanged)
    provider.on("chainChanged", onChainChanged)
    provider.on("disconnect", onDisconnect)
    
    return () => {
      console.log("Clean up window.ethereum events");
      provider.removeListener('chainChanged', onChainChanged);
      provider.removeListener('accountsChanged', onAccountChanged);
      provider.removeListener('disconnect', onDisconnect);
    }
  }, [provider, chain])

  const connect = async () => {
    if (!provider) {
      throw new Error(`No window.ethereum in the current browser session`)
    }

    try {
      setConnecting(true)
      setWallet(
        createWalletClient({
          chain: chain,
          transport: custom(provider),
        })
      )
      const address = await wallet?.requestAddresses()
      setAccount(address?.[0])
    } catch (e) {
      console.error(e)
    } finally {
      setConnecting(false)
    }
  }

  const disconnect = () => {
    try {
      sdk?.terminate()
    } catch (e) {
      console.error(e)
    }
  }

  return (
    <WalletContext.Provider
      value={{
        chain: chain,
        account: account,
        connect: connect,
        wallet: wallet,
        connecting: connecting,
        disconnect: disconnect,
      }}
    >
      {children}
    </WalletContext.Provider>
  )
}

export { WalletContext, WalletProvider }
