import type { Metadata } from "next";
import Script from "next/script";
import { space_grotesk } from "../fonts/fonts";
import "./globals.css";
import { Provider } from "@/components/ui/provider";
import { Toaster } from "@/components/ui/toaster";

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
      <Script async src={`https://www.googletagmanager.com/gtag/js?id=${process.env.NEXT_PUBLIC_GA_ID}`} />
      <Script id="google-analytics">
        {
          `
          window.dataLayer = window.dataLayer || [];
          function gtag(){dataLayer.push(arguments);}
          gtag('js', new Date());

          gtag('config', '${process.env.NEXT_PUGLIC_GA_ID}');
          `
        }
      </Script>
      <body>
        <Provider>
          <Toaster />
          {children}
        </Provider>
      </body>
    </html>
  );
}
