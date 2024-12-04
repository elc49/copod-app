"use client";

import { useContext, useMemo, useState } from "react";
import { useParams, useRouter } from "next/navigation";
import { useQuery, useMutation } from "@apollo/client";
import { Flex, SimpleGrid } from "@chakra-ui/react";
import Image from "next/image";
import { CREATE_USER } from "@/graphql/mutation";
import { GET_SUPPORTING_DOC_BY_ID } from "@/graphql/query";
import { WalletContext } from "@/providers/wallet";
import Loader from "@/components/loader";
import UserDetailsForm from "../components/form";
import withAuth from "@/providers/withAuth";
import { toaster } from "@/components/ui/toaster";

export default withAuth(Page)
function Page() {
  const router = useRouter()
  const params = useParams()
  const [saving, setSaving] = useState(false)
  const { isLoggedIn, user } = useContext(WalletContext)
  const { data: documentDetails, loading: documentDetailsLoading } = useQuery(GET_SUPPORTING_DOC_BY_ID, {
    skip: !isLoggedIn,
    variables: {
      id: params.id,
    },
  })
  const docDetails = useMemo(() => {
    return documentDetails?.getSupportingDocById
  }, [documentDetails])
  const [creaeUser, { loading: updatingUserDetails }] = useMutation(CREATE_USER)

  const saveDetails = (values: any) => {
    try {
      setSaving(true)
      creaeUser({
        variables: {
          input: {
            email: user?.email,
            firstname: values.firstname,
            lastname: values.lastname,
            govtid: values.govtid,
            verification: values.verification[0],
            supportDocId: params.id,
          },
        },
        onCompleted: () => {
          toaster.create({
            title: "Success",
            description: "User details updated",
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
    } finally {
      setSaving(false)
    }
  }

  return documentDetailsLoading ? <Loader /> : (
    <SimpleGrid columns={{ base: 1, sm: 2}} p="2" gap={{ base: "40px", sm: "24px" }}>
      <Flex direction="column" align="center" gap="4">
        <Image
          src={docDetails.govt_id}
          alt={docDetails.__typename}
          priority={true}
          width={500}
          height={500}
        />
      </Flex>
      <Flex direction="column" gap="4">
        <UserDetailsForm updating={saving || updatingUserDetails} saveDetails={saveDetails} />
      </Flex>
    </SimpleGrid>
  )
}
