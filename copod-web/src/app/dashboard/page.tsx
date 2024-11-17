"use client";

import withAuth from "@/hoc/withAuth";

export default withAuth(Page, "all")
function Page() {
  return (
    <h1>Hello, Dashboard page</h1>
  )
}
