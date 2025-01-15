import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Copod - Privacy Policy",
  description: "Register land. Search land. Buy land space.",
}

function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div>{children}</div>
  )
}

export default RootLayout
