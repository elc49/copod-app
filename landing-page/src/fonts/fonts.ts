import { Bricolage_Grotesque, Space_Grotesk } from "next/font/google"

const bricolage_grotesque = Bricolage_Grotesque({
  variable: "--font-bricolage",
  subsets: ["latin"],
  weight: "500",
})

const space_grotesk = Space_Grotesk({
  weight: "500",
  variable: "--font-space",
  subsets: ["latin"],
  display: "swap",
})

export { bricolage_grotesque, space_grotesk }
