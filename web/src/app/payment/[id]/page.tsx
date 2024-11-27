"use client";

import { useMemo } from "react";
import { useParams } from "next/navigation";
import { useQuery } from "@apollo/client";
import { Flex, Heading, SimpleGrid } from "@chakra-ui/react";
import Image from "next/image";
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
  const paymentDetails = useMemo(() => {
    return data?.getPaymentDetailsById
  }, [data])

  return loading ? <Loader /> : (
    <SimpleGrid columns={{ base: 1, sm: 2}} p="2" gap={{ base: "40px", sm: "24px" }}>
      <Flex direction="column" align="center" gap="4">
        <Heading>Land title</Heading>
        <Image
          src={paymentDetails.title.title}
          alt={paymentDetails.__typename}
          priority={true}
          width={500}
          height={500}
        />
      </Flex>
      <Flex direction="column" gap="4">
        <Heading>Registration form</Heading>
        <LandDetails />
      </Flex>
    </SimpleGrid>
  )
}

export default withAuth(Page)
