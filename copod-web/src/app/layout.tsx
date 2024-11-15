import type { Metadata } from "next";
import "./globals.css";
import { roboto_mono, roboto } from "./fonts/fonts"
import { Providers } from "@/providers/providers";

export const metadata: Metadata = {
  title: "Copod",
  description: "Register land. Search land. Buy land space.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${roboto_mono.className} ${roboto.className} antialiased`}>
      <body>
        <Providers>
          {children}
        </Providers>
      </body>
    </html>
  );
}
