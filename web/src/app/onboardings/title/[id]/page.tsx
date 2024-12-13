"use client";

import { SimpleGrid } from "@chakra-ui/react";
import withAuth from "@/providers/withAuth";

export default withAuth(Page)
function Page() {
  return (
    <SimpleGrid columns={{ base: 1, sm: 2 }} p="2" gap={{ base: "40px", sm: "24px" }}>
      Title details
    </SimpleGrid>
  )
}
