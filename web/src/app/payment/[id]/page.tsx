"use client";

import { useContext, useMemo, useState } from "react";
import { useParams } from "next/navigation";
import { useQuery } from "@apollo/client";
import { Flex, Heading, SimpleGrid } from "@chakra-ui/react";
import Image from "next/image";
import { GET_PAYMENT_DETAILS_BY_ID } from "@/graphql/query";
import { parseUnits } from "viem";
import LandDetails from "../form/LandDetails";
import Loader from "@/components/loader";
import withAuth from "@/providers/withAuth";
import { WalletContext } from "@/providers/wallet";
import { getAccounts, publicClient, privateClient } from "@/blockchain/rpc";

function Page() {
  const [registering, setRegistering] = useState(false)
  const { provider } = useContext(WalletContext)
  const params = useParams()
  const { data, loading } = useQuery(GET_PAYMENT_DETAILS_BY_ID, {
    variables: {
      id: params.id,
    },
  })
  const paymentDetails = useMemo(() => {
    return data?.getPaymentDetailsById
  }, [data])

  // TODO: break this down further to simplify
  const registerLand = async (title: string, size: number, unit: string) => {
    try {
      setRegistering(true)
      const registryContractAddress = await import("../../../../../SmartContract/ignition/deployments/chain-11155420/deployed_addresses.json")
      const abi: any = await import("../../../../../SmartContract/ignition/deployments/chain-11155420/artifacts/Registry#Registry.json")
      const account = await getAccounts(provider!)
      const { request } = await publicClient(provider!).simulateContract({
        account: account?.[0],
        address: registryContractAddress.default["Registry#Registry"],
        abi: abi?.abi,
        functionName: "register",
        args: [title, unit, account?.[0], parseUnits(size.toString(), 10)],
      })
      const hash = await privateClient(provider!).writeContract(request)
      const receipt = await publicClient(provider!).waitForTransactionReceipt({ hash })
      console.log(receipt)
    } catch (e) {
      console.error(e)
    } finally {
      setRegistering(false)
    }
  }

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
        <LandDetails registerLand={registerLand} registering={registering} />
      </Flex>
    </SimpleGrid>
  )
}

export default withAuth(Page)
