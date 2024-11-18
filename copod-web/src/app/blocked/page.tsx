"use client";

import { AbsoluteCenter, Heading } from "@chakra-ui/react"
import withAuth from "@/hoc/withAuth";

export default withAuth(Page)
function Page() {
  return (
    <AbsoluteCenter axis="both">
      <Heading>You are not an admin!</Heading>
    </AbsoluteCenter>
  )
}
