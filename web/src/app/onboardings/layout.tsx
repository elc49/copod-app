"use client";

import { useMemo } from "react";
import { Box, HStack, Heading, IconButton } from "@chakra-ui/react";
import { usePathname, useRouter } from "next/navigation";
import { ArrowLeft } from "@/components/icons";

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const router = useRouter()
  const pathname = usePathname()
  const pageTitle = useMemo(() => {
    if (pathname.includes("onboardings/document")) {
      return "Document details"
    }
    return "Payments"
  }, [pathname])

  return (
    <Box p="4">
      <HStack gap="4" align="center" py="4">
        <IconButton aria-label="Go back" onClick={() => router.back()}>
          <ArrowLeft />
        </IconButton>
        <Heading>{pageTitle}</Heading>
      </HStack>
      {children}
    </Box>
  )
}
