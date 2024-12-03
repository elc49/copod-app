"use client";

import { useContext, useMemo } from "react";
import { useQuery } from "@apollo/client";
import { GET_SUPPORTING_DOCS_BY_VERIFICATION } from "@/graphql/query";
import { Verification } from "@/graphql/graphql";
import { Box } from "@chakra-ui/react";
import { WalletContext } from "@/providers/wallet";
import SupportingDocsVerificationTable from "./components/SupportingDocsVerificationTable";
import Loader from "@/components/loader";

function Page() {
  const { isLoggedIn } = useContext(WalletContext)
  const { data, loading } = useQuery(GET_SUPPORTING_DOCS_BY_VERIFICATION, {
    variables: {
      verification: Verification.Onboarding,
    },
    skip: !isLoggedIn,
  })
  const supportingDocs = useMemo(() => {
    return data?.getSupportingDocsByVerification || []
  }, [data])

  return loading ? <Loader /> : (
    <Box>
      <SupportingDocsVerificationTable supportingDocs={supportingDocs} />
    </Box>
  )
}

export default Page
