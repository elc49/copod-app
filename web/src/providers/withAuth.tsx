"use client";

import * as React from "react";
import { useRouter } from "next/navigation";
import { WalletContext } from "@/providers/wallet";

const LOGIN_ROUTE = "/"
const DASHBOARD_ROUTE = "/dashboard"

function withAuth(Component: React.ComponentType) {
  const ComponentWithAuth: React.FC = (props) => {
    const { isLoggedIn } = React.useContext(WalletContext)
    const router = useRouter()

    React.useEffect(() => {
      if (!isLoggedIn) {
        router.replace(LOGIN_ROUTE)
      } else {
        router.replace(DASHBOARD_ROUTE)
      }
    }, [isLoggedIn, router])

    return <Component {...props} />
  }
  return ComponentWithAuth
}

export default withAuth
