"use client";

import { useContext, useMemo } from "react";
import { useQuery } from "@apollo/client";
import { Box } from "@chakra-ui/react";
import { WalletContext } from "@/providers/wallet";
import { GET_PAYMENTS_BY_STATUS } from "@/graphql/query";
import withAuth from "@/providers/withAuth";
import { PaymentStatus } from "@/graphql/graphql";
import PaymentsByStatusTable from "./components/PaymentsByStatusTable";

import Loader from "@/components/loader";

function Page() {
  const { isLoggedIn } = useContext(WalletContext)
  const { data, loading } = useQuery(GET_PAYMENTS_BY_STATUS, {
    variables: {
      status: PaymentStatus.Success,
    },
    skip: !isLoggedIn,
  })
  const payments = useMemo(() => {
    return data?.getPaymentsByStatus || []
  }, [data])

  return loading ? <Loader /> : (
    <Box p="4">
      <PaymentsByStatusTable payments={payments} />
    </Box>
  )
}

export default withAuth(Page)
