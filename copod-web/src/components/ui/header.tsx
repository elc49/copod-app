"use client";

import { useContext } from "react";
import { Box, Flex, Heading } from "@chakra-ui/react";
import { Button } from "./button";
import { AuthContext } from "@/context/Auth";

export default function Header() {
  const { loading, isLoggedIn, setLoading, web3auth, setProvider, setIsLoggedIn } = useContext(AuthContext)

  const logout = async () => {
    setLoading(true)
    try {
      await web3auth?.logout()
      setIsLoggedIn(false)
      setProvider(null)
    } catch (error) {
      console.error(error)
    } finally {
      setLoading(false)
    }
  }

  return (
    <Flex align="center" gap="4" p="4" justify="space-between">
      <Box marginEnd="auto">
        <Heading>
          Copod
        </Heading>
      </Box>
      {isLoggedIn && (
        <Box>
          <Button
            loading={loading}
            onClick={logout}
            p="4"
          >
            Sign out
          </Button>
        </Box>
      )}
    </Flex>
  )
}
