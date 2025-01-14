import type { Metadata } from "next";
import { space_grotesk } from "../fonts/fonts";
import "./globals.css";
import { Provider } from "@/components/ui/provider";

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
    <html lang="en" suppressHydrationWarning className={`${space_grotesk.variable} antialiased`}>
      <body>
        <Provider>{children}</Provider>
      </body>
    </html>
  );
}
