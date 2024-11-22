"use client";

import { useContext } from "react";
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuItem,
  DropdownMenuSeparator,
} from "@/components/ui/dropdown-menu";
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar";
import { User } from "lucide-react";
import { WalletContext} from "@/providers/wallet";
import { Button } from "@/components/ui/button";

function Header() {
  const { account, connect, disconnect } = useContext(WalletContext)

  return (
    <nav className="flex flex-row p-4">
      <div>
        <h1 className="font-bold text-2xl">Copod</h1>
      </div>
      <div className="ml-auto">
        {account ? (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Avatar>
                <AvatarImage src="https://effigy.im/a/0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045.png" />
                <AvatarFallback>
                  <User />
                </AvatarFallback>
              </Avatar>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuLabel>Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={disconnect}
              >
                Log out
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button
            className="font-bold"
            onClick={connect}
          >
            Connect Wallet
          </Button>
        )}
      </div>
    </nav>
  )
}

export default Header
