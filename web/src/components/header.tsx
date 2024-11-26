"use client";

import { useContext } from "react";
import { Flex, Heading } from "@chakra-ui/react";
import { WalletContext} from "@/providers/wallet";
import {
  MenuContent,
  MenuItem,
  MenuTrigger,
  MenuRoot,
} from "@/components/ui/menu";
import {
  Avatar,
} from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";

function Header() {
  const { isLoggedIn, login, user, logout } = useContext(WalletContext)

  return (
    <Flex gap="4" p="4" justify="space-between">
      <Flex marginEnd="auto">
        <Heading>Copod</Heading>
      </Flex>
      {isLoggedIn ? (
        <Flex>
          <MenuRoot>
            <MenuTrigger>
              <Avatar
                src={user?.profileImage}
              />
            </MenuTrigger>
            <MenuContent>
              <MenuItem fontWeight="bold" onClick={logout} value="sign-out">
                Sign out
              </MenuItem>
            </MenuContent>
          </MenuRoot>
        </Flex>
      ) : (
        <Button onClick={login}>
          Connect Wallet
        </Button>
      )}
    </Flex>
  )
}

export default Header
