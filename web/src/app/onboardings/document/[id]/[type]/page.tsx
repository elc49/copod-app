"use client";

import { useContext, useMemo, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useMutation, useQuery } from "@apollo/client";
import { Button, Image as ChakraImage,Flex, SimpleGrid } from "@chakra-ui/react";
import NextImage from "next/image";
import updateTitleVerificationById from "@/graphql/mutation/UpdateTitleVerificationById";
import createUser from "@/graphql/mutation/CreateUser";
import getSupportDocById from "@/graphql/query/GetSupportingDocById";
import getTitleById from "@/graphql/query/GetTitleById";
import getDisplayPictureById from "@/graphql/query/GetDisplayPictureById";
import { parseUnits } from "viem";
import LandDetails from "../../../form/LandDetails";
import UserDetailsForm from "../../../form/UserDetails";
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
  const { data: title, loading: titleLoading } = useQuery(getTitleById, {
    variables: {
      id: params.id,
    },
    skip: !(params.type === "title"),
  })
  const titleData = useMemo(() => {
    return title?.getTitleById
  }, [title])
  const [updateTitleVerification, { loading: updatingTitle }] = useMutation(updateTitleVerificationById)
  const { data: supportDoc, loading: supportDocLoading } = useQuery(getSupportDocById, {
    variables: {
      id: params.id,
    },
    skip: !(params.type == "supportingdoc"),
  })
  const docDetails = useMemo(() => {
    return supportDoc?.getSupportingDocById
  }, [supportDoc])
  const [createNewUser, { loading: creatingUser }] = useMutation(createUser)
  const { data: displayPicture, loading: gettingDisplayPicture } = useQuery(getDisplayPictureById, {
    variables: {
      id: params.id,
    },
    skip: !(params.type === "displaypicture"),
  })
  const dpDetails = useMemo(() => {
    return displayPicture?.getDisplayPictureById
  }, [displayPicture])
  const router = useRouter()

  const saveUser = (values: any) => {
    try {
      createNewUser({
        variables: {
          input: {
            email: docDetails.email,
            firstname: values.firstname,
            lastname: values.lastname,
            supportDocId: params.id,
            supportDocVerification: values.verification[0],
          },
        },
        onCompleted: () => {
          toaster.create({
            title: "Success",
            description: "User created",
            type: "success",
          })
          router.back()
        },
        onError: (e) => {
          toaster.create({
            title: "Error",
            description: `${e.message}`,
            type: "error",
          })
        },
      })
    } catch (e) {
      console.error(e)
    }
  }

  const saveLandLocally = (status: string) => {
    updateTitleVerification({
      variables: {
        input: {
          titleId: params.id,
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
    saveLandLocally(status)
    toaster.create({
      title: "Success",
      description: "Land registered",
      type: "success",
    })
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

  return (titleLoading || supportDocLoading || gettingDisplayPicture) ? <Loader /> : (
    <SimpleGrid columns={{ base: 1, sm: 2}} p="2" gap={{ base: 4, sm: 8 }}>
      {params.type === "title" && (
        <>
          <Flex direction="column" align="center" gap="4">
            <ChakraImage asChild>
              <NextImage
                src={`${titleData.url}`}
                alt={titleData.__typename}
                priority={true}
                width={500}
                height={500}
              />
            </ChakraImage>
          </Flex>
          <Flex direction="column" gap="4">
            {titleData.verified === "VERIFIED" ? (
              <DoneIcon />
            ) : (
              <LandDetails registering={registering || updatingTitle} registerLand={registerLand} />
            )}
          </Flex>
        </>
      )}
      {params.type === "supportingdoc" && (
        <>
          <Flex direction="column" align="center" gap="4">
            <ChakraImage asChild>
              <NextImage
                src={`${docDetails.url}`}
                alt={docDetails.__typename}
                priority={true}
                width={500}
                height={500}
              />
            </ChakraImage>
          </Flex>
          <Flex direction="column" gap="4">
            <UserDetailsForm updating={creatingUser} saveUser={saveUser} />
          </Flex>
        </>
      )}
      {params.type === "displaypicture" && (
        <>
          <Flex gap="4" align="center">
            <ChakraImage asChild>
              <NextImage
                src={`${dpDetails.url}`}
                alt={dpDetails.__typename}
                priority={true}
                width={500}
                height={500}
              />
            </ChakraImage>
          </Flex>
          <Flex gap="4">
            <Button
              onClick={() => {}}
            >
              Save
            </Button>
          </Flex>
        </>
      )}
    </SimpleGrid>
  )
}
