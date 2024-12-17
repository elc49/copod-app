import { Roboto, Roboto_Mono } from "next/font/google"

const roboto_mono = Roboto_Mono({
  weight: "300",
  variable: "--font-roboto-mono",
  style: "normal",
  subsets: ["latin"],
  display: "swap",
})

const roboto = Roboto({
  weight: "400",
  variable: "--font-roboto",
  subsets: ["latin"],
  display: "swap",
})

export { roboto_mono, roboto }
