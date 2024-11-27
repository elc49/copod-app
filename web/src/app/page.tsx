"use client";

import { useContext } from "react";
import { AbsoluteCenter, Heading, Stack } from "@chakra-ui/react";
import { WalletContext } from "@/providers/wallet";
import Loader from "@/components/loader";

const lines = ["Register land.", "Search land.", "Buy land usage rights."]

function Home() {
  const { initializing } = useContext(WalletContext)

  return (
    <AbsoluteCenter axis="both">
      {initializing ? <Loader /> : (
        <Stack gap="4">
          {lines.map((line, index) => <Heading key={index} size={{ base: "4xl", md: "5xl", lg: "6xl" }}>{line}</Heading>)}
          <p>&copy; Copod {new Date().getFullYear()}</p>
        </Stack>
      )}
    </AbsoluteCenter>
  );
}

export default Home
