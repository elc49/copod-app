"use client";

import { useContext } from "react";
import { useQuery } from "@apollo/client";
import { WalletContext } from "@/providers/wallet";
import { GET_PAYMENTS_BY_STATUS } from "@/graphql/GetPaymentsByStatus";

function Page() {
  const { account } = useContext(WalletContext)
  const { data } = useQuery(GET_PAYMENTS_BY_STATUS, {
    variables: {
      status: "success",
    },
    skip: !account,
  })
  console.log(data)

  return (
    <h1 className="text-4xl font-bold">Dashboard</h1>
  )
}

export default Page
