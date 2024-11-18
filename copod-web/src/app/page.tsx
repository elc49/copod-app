"use client";

import { useContext } from "react";
import { AbsoluteCenter } from "@chakra-ui/react";
import { Button } from "@/components/ui/button"
import { AuthContext } from "@/context/Auth";
import withAuth from "@/hoc/withAuth";

export default withAuth(Home)
function Home() {
  const { web3auth, loading, setLoading, setProvider, setIsLoggedIn } = useContext(AuthContext);
  
  const login = async () => {
    setLoading(true)
    try {
      const web3authProvider = await web3auth?.connect()
      setProvider(web3authProvider)
      if (web3auth?.connected) {
        setIsLoggedIn(true)
      }
    } catch (error) {
      console.error(error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <AbsoluteCenter axis="both">
      <Button
        colorPalette="green"
        p={4}
        onClick={login}
        loading={loading}
      >
        Sign in
      </Button>
    </AbsoluteCenter>
  );
}
