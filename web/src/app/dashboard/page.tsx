"use client";

import { useContext, useMemo } from "react";
import { useQuery } from "@apollo/client";
import { WalletContext } from "@/providers/wallet";
import { GET_PAYMENTS_BY_STATUS } from "@/graphql/query/GetPaymentsByStatus";
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

  return loading ? <Loader /> : <PaymentsByStatusTable payments={payments} />
}

export default withAuth(Page)
