import { PropsWithChildren } from "react";
import { Provider } from "@/components/ui/provider";

export function ChakraUIProvider({ children }: PropsWithChildren) {
  return (
    <Provider>{children}</Provider>
  )
}
