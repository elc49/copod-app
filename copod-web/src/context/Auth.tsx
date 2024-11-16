"use client";

import { createContext, Dispatch, PropsWithChildren, SetStateAction, useEffect, useState } from "react"
import { IAdapter, IProvider, UserInfo } from "@web3auth/base"
import { Web3Auth } from "@web3auth/modal";
import { getWeb3AuthOptions } from "@/web3/Web3";
import { getDefaultExternalAdapters } from "@web3auth/default-evm-adapter";

interface AuthContext {
  isLoggedIn: boolean
  setIsLoggedIn: Dispatch<SetStateAction<boolean>>
  provider: IProvider | null | undefined
  web3auth: Web3Auth | undefined
  user: Partial<UserInfo> | undefined
  setUser: Dispatch<SetStateAction<Partial<UserInfo> | undefined>>
  setProvider: Dispatch<SetStateAction<IProvider | null | undefined>>
}

const AuthContext = createContext<AuthContext>({
  isLoggedIn: false,
  setIsLoggedIn: () => {},
  provider: null,
  web3auth: undefined,
  user: undefined,
  setUser: () => {},
  setProvider: () => {},
})

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [provider, setProvider] = useState<IProvider | null | undefined>(null)
  const [loading, setLoading] = useState(true)
  const [web3auth, setWeb3auth] = useState<Web3Auth | undefined>()
  const [user, setUser] = useState<Partial<UserInfo> | undefined>()

  useEffect(() => {
    const init = async() => {
      const options = getWeb3AuthOptions()
      const web3auth = new Web3Auth(options)

      try {
        const adapters = getDefaultExternalAdapters({ options: options })
        adapters.forEach((adapter: IAdapter<unknown>) => {
          web3auth.configureAdapter(adapter)
        })
        await web3auth.initModal()
        setWeb3auth(web3auth)
        setProvider(web3auth.provider)

        if (web3auth.status === "ready") {
          setLoading(false)
        }
      } catch (error) {
        console.error(error)
        setLoading(false)
      }
    }

    init()
  }, [])

  if (loading) return (
    <div className="grid bg-green-600 w-full place-content-center">
      Loading...
    </div>
  )

  return (
    <AuthContext.Provider
      value={{
        isLoggedIn,
        setIsLoggedIn,
        provider,
        web3auth,
        user,
        setUser,
        setProvider,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
