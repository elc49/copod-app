import type { Metadata } from "next";
import { roboto, roboto_mono } from "./fonts/fonts"
import "./globals.css";
import { Provider } from "@/components/ui/provider"
import { AuthProvider } from "@/context/Auth";
import ApolloProvider from "@/context/apollo";
import Header from "@/components/ui/header";

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
            <ApolloProvider>
              <Header />
              {children}
            </ApolloProvider>
          </AuthProvider>
        </Provider>
      </body>
    </html>
  );
}
