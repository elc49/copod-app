import type { Metadata } from "next";
import { space_grotesk } from "../fonts/fonts";
import "./globals.css";
import { Provider } from "@/components/ui/provider";
import { Toaster } from "@/components/ui/toaster";
import { GoogleAnalytics } from "@next/third-parties/google";

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
        <Provider>
          <Toaster />
          {children}
          <GoogleAnalytics gaId={`${process.env.NEXT_PUBLIC_GA_ID}`} />
        </Provider>
      </body>
    </html>
  );
}
