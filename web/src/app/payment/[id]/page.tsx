"use client";

import { useParams } from "next/navigation";
import { useQuery } from "@apollo/client";
import { Box, Heading, SimpleGrid } from "@chakra-ui/react";
import { GET_PAYMENT_DETAILS_BY_ID } from "@/graphql/query";
import LandDetails from "../form/LandDetails";
import Loader from "@/components/loader";
import withAuth from "@/providers/withAuth";

function Page() {
  const params = useParams()
  const { data, loading } = useQuery(GET_PAYMENT_DETAILS_BY_ID, {
    variables: {
      id: params.id,
    },
  })
  console.log(data)

  return loading ? <Loader /> : (
    <SimpleGrid columns={{ base: 1, sm: 2}} gap={{ base: "40px", sm: "24px" }}>
      <Box>
        <Heading>Land title</Heading>
      </Box>
      <Box>
        <Heading>Registration form</Heading>
        <LandDetails />
      </Box>
    </SimpleGrid>
  )
}

export default withAuth(Page)
