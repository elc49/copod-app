import { useContext, useEffect } from "react";
import { useRouter } from "next/navigation";
import { AuthContext } from "@/context/Auth";

const DASHBOARD_ROUTE = "/dashboard"
const LOGIN_ROUTE = "/"

type RouteRole = "auth" | "optional" | "all"

export default function withAuth(Component: React.ComponentType, routeRole: RouteRole) {
  const ComponentWithAuth: React.FC = (props) => {
    const { isLoggedIn, isAdmin } = useContext(AuthContext)
    const router = useRouter()

    useEffect(() => {
      if (!isAdmin) router.replace("/blocked")
      if (isLoggedIn) router.replace(DASHBOARD_ROUTE)
      if (routeRole !== "auth" && routeRole !== "optional") router.replace(LOGIN_ROUTE)

    }, [isAdmin, isLoggedIn, router])

    // Optionally return null or a loading state while redirecting
    if (isLoggedIn && routeRole === "auth") {
      return null // Prevent rendering the wrapped component while redirecting
    }

    return <Component {...(props)} />
  }

  return ComponentWithAuth
}
