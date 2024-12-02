"use client";

import { Box, HStack, Heading, IconButton } from "@chakra-ui/react";
import { useRouter } from "next/navigation";
import { ArrowLeft } from "@/components/icons";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter()

  return (
    <Box p="4">
      <HStack gap="4" py="4" align="center">
        <IconButton aria-label="Go back" onClick={() => router.back()}>
          <ArrowLeft />
        </IconButton>
        <Heading>Payment details</Heading>
      </HStack>
      {children}
    </Box>
  )
}
