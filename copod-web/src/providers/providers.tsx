"use client";

import { PropsWithChildren } from "react";
import { AuthProvider } from "@/context/Auth";

export const Providers = ({ children }: PropsWithChildren) => (
  <>
    <AuthProvider>
      {children}
    </AuthProvider>
  </>
)
