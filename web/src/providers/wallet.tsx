"use client";

import { createContext, PropsWithChildren, useCallback, useEffect, useMemo, useState } from "react";
import "viem/window";
import { EthereumPrivateKeyProvider } from "@web3auth/ethereum-provider";
import { Web3AuthNoModal } from "@web3auth/no-modal";
import { AuthAdapter } from "@web3auth/auth-adapter";
import { WEB3AUTH_NETWORK, WALLET_ADAPTERS, IProvider, UserInfo } from "@web3auth/base";
import { createWalletClient, custom, WalletClient } from "viem";
import { optimismSepolia } from "viem/chains";
import { useRouter } from "next/navigation";
import chainConfig from "@/blockchain/chains";

interface IWalletContext {
  provider: IProvider | null
  isLoggedIn: boolean
  initializing: boolean
  login: () => void
  logout: () => void
  user: Partial<UserInfo> | null
  wallet: WalletClient | null
}

const WalletContext = createContext<IWalletContext>({
  provider: null,
  isLoggedIn: false,
  initializing: true,
  login: () => {},
  logout: () => {},
  user: null,
  wallet: null,
})

const WalletProvider = ({ children }: PropsWithChildren) => {
  const [provider, setProvider] = useState<IProvider | null>(null)
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false)
  const [initializing, setInitializing] = useState<boolean>(true)
  const [user, setUser] = useState<Partial<UserInfo> | null>(null)
  const router = useRouter()
  const [wallet, setWallet] = useState<WalletClient | null>(null)

  const privateKeyProvider = useMemo(() => {
    return new EthereumPrivateKeyProvider({ config: { chainConfig } })
  }, [])

  const adapter = useMemo(() => {
    return new AuthAdapter({
      adapterSettings: {
        uxMode: "redirect",
        loginConfig: {
          // Google config
          google: {
            verifier: "google-web3auth-dev",
            typeOfLogin: "google",
            clientId: process.env.NEXT_PUBLIC_WEB3_AUTH_GOOGLE_CLIENT_ID!,
          },
        },
      },
    })
  }, [])

  const web3auth = useMemo(() => {
    const web3Client = new Web3AuthNoModal({
      clientId: process.env.NEXT_PUBLIC_WEB3_AUTH_CLIENT_ID!,
      web3AuthNetwork: WEB3AUTH_NETWORK.SAPPHIRE_DEVNET,
      privateKeyProvider,
    })
    web3Client.configureAdapter(adapter)
    return web3Client
  }, [])

  useEffect(() => {
    async function init() {
      try {
        await web3auth.init()
        setProvider(web3auth.provider)

        if (web3auth.status === "connected") {
          const user = await web3auth.getUserInfo()
          setUser(user)
          setIsLoggedIn(true)
          router.replace("/dashboard")
        }
      } catch (e) {
        console.error(e)
      } finally {
        setInitializing(false)
      }
    }

    init()
  }, [web3auth])

  useEffect(() => {
    if (provider != null) {
      const wallet = createWalletClient({
        chain: optimismSepolia,
        transport: custom(web3auth.provider!),
      })
      setWallet(wallet)
    }
  }, [provider, web3auth])

  const login = useCallback(() => {
    async function web3authLogin() {
      try {
        setInitializing(true)
        const provider = await web3auth.connectTo(WALLET_ADAPTERS.AUTH, {
          loginProvider: "google",
        })
        setProvider(provider)
        if (web3auth.status === "connected") {
          const user = await web3auth.getUserInfo()
          setUser(user)
          setIsLoggedIn(true)
        }
      } catch (e) {
        console.error(e)
      }
    }

    web3authLogin()
  }, [web3auth])

  const logout = useCallback(() => {
    async function web3authLogout() {
      try {
        setInitializing(true)
        await web3auth.logout()
        web3auth.clearCache()
        setProvider(null)
        setUser(null)
        setIsLoggedIn(false)
        setInitializing(false)
      } catch (e) {
        console.error(e)
      }
    }

    web3authLogout()
  }, [web3auth])

  return (
    <WalletContext.Provider
      value={{
        provider,
        isLoggedIn,
        initializing,
        login,
        logout,
        user,
        wallet,
      }}
    >
      {children}
    </WalletContext.Provider>
  )
}

export { WalletContext, WalletProvider }
