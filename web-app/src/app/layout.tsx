import type { Metadata } from "next";
import "./globals.css";
import { space_grotesk } from "@/font/font";

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className={`${space_grotesk.variable} antialiased`}>
      <body>
        {children}
      </body>
    </html>
  );
}
