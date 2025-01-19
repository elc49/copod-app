import type { Metadata } from "next";
import "./globals.css";
import { space_grotesk } from "@/font/font";
import { Providers } from "./providers";

export const metadata: Metadata = {
  title: "Copod - Own land. Search land. Use land.",
  description: "Own land, Search land. Use land.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${space_grotesk.variable} antialiased`}>
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
