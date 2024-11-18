"use client";

import { useContext } from "react";
import { useQuery } from "@apollo/client";
import withAuth from "@/hoc/withAuth";
import { GetLands } from "@/apollo/query/GetLands";
import { AuthContext } from "@/context/Auth";

export default withAuth(Page)
function Page() {
  const { isLoggedIn } = useContext(AuthContext)
  const { data } = useQuery(GetLands, {
    skip: !isLoggedIn,
  })
  console.log(data)

  return (
    <h1>Hello, Dashboard page</h1>
  )
}
