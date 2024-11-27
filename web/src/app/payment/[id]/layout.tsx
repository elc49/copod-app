"use client";

import { Box, IconButton } from "@chakra-ui/react";
import { ArrowLeft } from "lucide-react";
import { useRouter } from "next/navigation";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter()

  return (
    <Box p="4">
      <Box py="4">
        <IconButton aria-label="Go back" onClick={() => router.back()}>
          <ArrowLeft />
        </IconButton>
      </Box>
      {children}
    </Box>
  )
}
