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
  const { isLoggedIn, login, user, logout } = useContext(WalletContext)

  return (
    <nav className="flex flex-row p-4">
      <div>
        <h1 className="font-bold text-2xl">Copod</h1>
      </div>
      <div className="ml-auto">
        {isLoggedIn ? (
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Avatar>
                <AvatarImage src={`${user?.profileImage}`} />
                <AvatarFallback>
                  <User />
                </AvatarFallback>
              </Avatar>
            </DropdownMenuTrigger>
            <DropdownMenuContent>
              <DropdownMenuLabel>Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem
                onClick={logout}
              >
                Log out
              </DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
        ) : (
          <Button
            className="font-bold"
            onClick={login}
          >
            Connect Wallet
          </Button>
        )}
      </div>
    </nav>
  )
}

export default Header
