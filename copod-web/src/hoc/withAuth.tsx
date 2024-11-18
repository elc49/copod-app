import { useContext, useEffect } from "react";
import { useRouter } from "next/navigation";
import { AuthContext } from "@/context/Auth";

const DASHBOARD_ROUTE = "/dashboard"
const LOGIN_ROUTE = "/"
const BLOCKED_ROUTE = "/blocked"

export default function withAuth(Component: React.ComponentType) {
  const ComponentWithAuth: React.FC = (props) => {
    const { isLoggedIn, isAdmin } = useContext(AuthContext)
    const router = useRouter()

   useEffect(() => {
     if (isLoggedIn) {
       if (isAdmin) {
         router.replace(DASHBOARD_ROUTE)
       } else {
         router.replace(BLOCKED_ROUTE)
       }
     } else {
         router.replace(LOGIN_ROUTE)
     }
   }, [isLoggedIn, router, isAdmin])

   return <Component {...props} />
  }

  return ComponentWithAuth
}
