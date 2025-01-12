import { Roboto_Mono, Space_Grotesk } from "next/font/google"

const roboto_mono = Roboto_Mono({
  weight: "400",
  variable: "--font-roboto-mono",
  style: "normal",
  subsets: ["latin"],
  display: "swap",
})

const space_grotesk = Space_Grotesk({
  weight: "500",
  variable: "--font-space-grotesk",
  subsets: ["latin"],
  display: "swap",
})

export { roboto_mono, space_grotesk }
