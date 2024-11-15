"use client";

import { createContext, Dispatch, PropsWithChildren, SetStateAction, useEffect, useState } from "react"
import { IAdapter, IProvider, UserInfo } from "@web3auth/base"
import { Web3Auth } from "@web3auth/modal";
import { getWeb3AuthOptions } from "@/web3/Web3";
import { getDefaultExternalAdapters } from "@web3auth/default-evm-adapter";

interface AuthContext {
  loading: boolean
  isLoggedIn: boolean
  provider: IProvider | null | undefined
  web3auth: Web3Auth | undefined
  user: Partial<UserInfo> | undefined
  setProvider: Dispatch<SetStateAction<IProvider | null | undefined>>
}

const AuthContext = createContext<AuthContext>({
  loading: false,
  isLoggedIn: false,
  provider: null,
  web3auth: undefined,
  user: undefined,
  setProvider: () => {},
})

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [provider, setProvider] = useState<IProvider | null | undefined>(null)
  const [loading, setLoading] = useState(false)
  const [web3auth, setWeb3auth] = useState<Web3Auth | undefined>()
  const [userInfo, setUserinfo] = useState<Partial<UserInfo> | undefined>()

  useEffect(() => {
    const init = async() => {
      setLoading(true)
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

        if (web3auth.connected) {
          const user = await web3auth.getUserInfo()
          setUserinfo(user)
          setIsLoggedIn(true)
          setLoading(false)
        }
      } catch (error) {
        console.error(error)
      }
    }

    init()
  }, [])

  return (
    <AuthContext.Provider
      value={{
        loading,
        isLoggedIn,
        provider,
        web3auth,
        user: userInfo,
        setProvider,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
