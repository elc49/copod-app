"use client";

import { useContext, useMemo } from "react";
import { useQuery } from "@apollo/client";
import { Box } from "@chakra-ui/react";
import { PaymentStatus } from "@/graphql/graphql";
import getPaymentsByStatus from "@/graphql/query/GetPaymentsByStatus";
import PaymentsByStatusTable from "./components/PaymentsByStatusTable";
import { WalletContext } from "@/providers/wallet";
import withAuth from "@/providers/withAuth";

import Loader from "@/components/loader";

export default withAuth(Page)
function Page() {
  const { isLoggedIn } = useContext(WalletContext)
  const { data, loading } = useQuery(getPaymentsByStatus, {
    variables: {
      status: PaymentStatus.Success,
    },
    skip: !isLoggedIn,
  })
  const payments = useMemo(() => {
    return data?.getPaymentsByStatus || []
  }, [data])

  return loading ? <Loader /> : (
    <Box>
      <PaymentsByStatusTable payments={payments} />
    </Box>
  )
}
