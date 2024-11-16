import type { Metadata } from "next";
import { roboto, roboto_mono } from "./fonts/fonts"
import "./globals.css";
import { Provider } from "@/components/ui/provider"
import { AuthProvider } from "@/context/Auth";

export const metadata: Metadata = {
  title: "Copod",
  description: "Register land. Search land. Buy land space",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body className={`${roboto_mono.className} ${roboto.variable}`}>
        <Provider>
          <AuthProvider>
            {children}
          </AuthProvider>
        </Provider>
      </body>
    </html>
  );
}
