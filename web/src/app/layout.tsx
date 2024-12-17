import type { Metadata } from "next";
import "./globals.css";
import { roboto, roboto_mono } from "./fonts/fonts";
import { Providers } from "@/providers/root";
import Header from "@/components/header";

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
    <html lang="en" suppressHydrationWarning className={`${roboto.variable} ${roboto_mono.variable} antialiased`}>
      <body>
        <Providers>
          <Header />
          {children}
        </Providers>
      </body>
    </html>
  );
}
