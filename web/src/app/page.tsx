"use client";

import { useContext } from "react";
import { WalletContext } from "@/providers/wallet";
import Loader from "@/components/loader";
import withAuth from "@/providers/withAuth";

const lines = ["Register land.", "Search land.", "Buy land usage rights."]

function Home() {
  const { initializing } = useContext(WalletContext)

  return (
    <div className="grid grid-rows-[20px_1fr_20px] min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2  items-start">
        {initializing ? <Loader /> : (
          <>
            {lines.map((line, index) => <h2 key={index} className="md:text-9xl text-6xl">{line}</h2>)}
          </>
        )}
      </main>
      <footer className="row-start-3 flex gap-6 flex-wrap items-center justify-center">
        <p>&copy; Copod {new Date().getFullYear()}</p>
      </footer>
    </div>
  );
}

export default withAuth(Home)
