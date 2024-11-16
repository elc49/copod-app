import { Roboto_Mono, Roboto } from "next/font/google"

const roboto = Roboto({
  weight: "300",
  variable: "--font-roboto",
  subsets: ["latin"],
  display: "swap",
})

const roboto_mono = Roboto_Mono({
  weight: "400",
  variable: "--font-roboto-mono",
  style: "normal",
  subsets: ["latin"],
  display: "swap",
})

export { roboto_mono, roboto }
