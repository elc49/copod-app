"use client";

import { useContext, useMemo, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useMutation, useQuery } from "@apollo/client";
import { Flex, SimpleGrid } from "@chakra-ui/react";
import Image from "next/image";
import { UPDATE_TITLE_VERIFICATION } from "@/graphql/mutation";
import { parseUnits } from "viem";
import LandDetails from "../../../form/LandDetails";
import Loader from "@/components/loader";
import { toaster } from "@/components/ui/toaster";
import { DoneIcon } from "@/components/icons";
import withAuth from "@/providers/withAuth";
import { WalletContext } from "@/providers/wallet";
import { getAccounts, publicClient, privateClient } from "@/blockchain/rpc";

export default withAuth(Page)
function Page() {
  const [registering, setRegistering] = useState(false)
  const { provider } = useContext(WalletContext)
  const params = useParams()
  console.log(params)
  const [updateTitleVerification, { loading: updatingTitleVerification }] = useMutation(UPDATE_TITLE_VERIFICATION)
  const router = useRouter()

  const saveLandLocally = (status: string) => {
    updateTitleVerification({
      variables: {
        input: {
          id: "",
          verification: status,
        },
      },
      onCompleted: () => {
        toaster.create({
          title: "Success",
          description: "Land saved locally",
          type: "success",
        })
        router.back()
      },
    })
  }

  const onSuccess = (status: string) => {
    toaster.create({
      title: "Success",
      description: "Land registered",
      type: "success",
    })
    saveLandLocally(status)
  }

  const onFailure = () => {
    toaster.create({
      title: "Error",
      description: "Something went wrong",
      type: "error",
    })
  }

  // TODO: break this down further to simplify
  const registerLand = async (title: string, size: number, unit: string, status: string) => {
    try {
      setRegistering(true)
      const registryContractAddress = await import("../../../../../../../SmartContract/ignition/deployments/chain-11155420/deployed_addresses.json")
      const abi: any = await import("../../../../../../../SmartContract/ignition/deployments/chain-11155420/artifacts/Registry#Registry.json")
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
      onSuccess(status)
      console.log(receipt)
    } catch (e) {
      onFailure()
      console.error(e)
    } finally {
      setRegistering(false)
    }
  }

  return false ? <Loader /> : (
    <SimpleGrid columns={{ base: 1, sm: 2}} p="2" gap={{ base: "40px", sm: "24px" }}>
      <Flex direction="column" align="center" gap="4">
        {/*<Image
          src={""}
          alt={""}
          priority={true}
          width={500}
          height={500}
          />*/}
      </Flex>
      {/*<Flex direction="column" gap="4">
        {paymentDetails.title.verified === "VERIFIED" ? (
          <DoneIcon />
        ) : (
          <LandDetails registerLand={registerLand} registering={registering || updatingTitleVerification} />
        )}
        </Flex>*/}
    </SimpleGrid>
  )
}
