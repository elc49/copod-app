"use client";

import { useMemo } from "react";
import { AbsoluteCenter, Card, Stack } from "@chakra-ui/react";
import NextLink from "next/link";
import withAuth from "@/providers/withAuth";

export default withAuth(Page)
function Page() {
  const cardPoints = useMemo(() => {
    return ["lands", "docs"]
  }, [])

  return (
    <AbsoluteCenter>
      <Stack direction="row" gap="4" wrap="wrap">
        {cardPoints.map((point: string, i: number) => (
          <NextLink key={i} href={`/${point}`}>
            <Card.Root key={i} size="lg" width="320px" onClick={() => {}}>
              <Card.Body textAlign="center" textTransform="capitalize" fontSize="4xl" fontWeight="bold">
                {point}
              </Card.Body>
            </Card.Root>
          </NextLink>
        ))}
      </Stack>
    </AbsoluteCenter>
  )
}
