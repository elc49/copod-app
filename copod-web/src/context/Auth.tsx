"use client";

import { createContext, Dispatch, PropsWithChildren, SetStateAction, useEffect, useState } from "react"
import { IProvider, UserInfo } from "@web3auth/base"
import { Web3Auth } from "@web3auth/modal";
import { getAccounts, getWeb3AuthOptions } from "@/web3/web3";
import { AbsoluteCenter, Spinner } from "@chakra-ui/react";
import { AuthAdapter } from "@web3auth/auth-adapter";

const ADMINS = process.env.NEXT_PUBLIC_WEB3_ADMINS!

interface IAuthContext {
  isLoggedIn: boolean
  setIsLoggedIn: Dispatch<SetStateAction<boolean>>
  provider: IProvider | null | undefined
  web3auth: Web3Auth | undefined
  user: Partial<UserInfo> | undefined
  setUser: Dispatch<SetStateAction<Partial<UserInfo> | undefined>>
  setProvider: Dispatch<SetStateAction<IProvider | null | undefined>>
  isAdmin: boolean
}

export const AuthContext = createContext<IAuthContext>({
  isLoggedIn: false,
  setIsLoggedIn: () => {},
  provider: null,
  web3auth: undefined,
  user: undefined,
  setUser: () => {},
  setProvider: () => {},
  isAdmin: false,
})

export const AuthProvider = ({ children }: PropsWithChildren) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [provider, setProvider] = useState<IProvider | null | undefined>(null)
  const [loading, setLoading] = useState(true)
  const [web3auth, setWeb3auth] = useState<Web3Auth | undefined>()
  const [user, setUser] = useState<Partial<UserInfo> | undefined>()
  const [isAdmin, setIsAdmin] = useState<boolean>(false)

  useEffect(() => {
    const init = async() => {
      const options = getWeb3AuthOptions()
      const web3auth = new Web3Auth(options)

      try {
        const passwordlessAdapter = new AuthAdapter({
          adapterSettings: {
            loginConfig: {
              // Email passwordless config
              email_passwordless: {
                verifier: "copod-email-passwordless",
                typeOfLogin: "email_passwordless",
                clientId: options.clientId,
              },
            },
          },
        })
        web3auth.configureAdapter(passwordlessAdapter)
        await web3auth.initModal()
        setWeb3auth(web3auth)
        setProvider(web3auth.provider)

        if (web3auth.status === "ready") {
          setLoading(false)
        } else if (web3auth.connected) {
          if (provider != null) {
            const account = await getAccounts(provider)
            if (ADMINS.indexOf(account) != -1) {
              setIsAdmin(true)
            }
          }
          setIsLoggedIn(true)
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
    <AbsoluteCenter axis="both">
        <Spinner color="green.600" animationDuration="0.8s" borderWidth="4px" size="md" />
    </AbsoluteCenter>
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
        isAdmin,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}
