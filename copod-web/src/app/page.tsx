"use client";

import { useContext } from "react";
import { AbsoluteCenter } from "@chakra-ui/react";
import { Button } from "@/components/ui/button"
import { AuthContext } from "@/context/Auth";
import withAuth from "@/hoc/withAuth";

export default withAuth(Home, "all")
function Home() {
  const { web3auth, setProvider, setIsLoggedIn } = useContext(AuthContext);
  
  const login = async () => {
    try {const web3authProvider = await web3auth?.connect()
      setProvider(web3authProvider)
      if (web3auth?.connected) {
        setIsLoggedIn(true)
      }
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <AbsoluteCenter axis="both">
      <Button
        colorPalette="green"
        p={4}
        onClick={login}
      >
        Sign in
      </Button>
    </AbsoluteCenter>
  );
}
