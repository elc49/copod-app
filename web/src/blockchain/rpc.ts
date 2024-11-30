import { IProvider } from "@web3auth/base";
import { createWalletClient, custom, createPublicClient, GetAddressesReturnType } from "viem";
import { optimismSepolia } from "viem/chains";

export const getAccounts = async (provider: IProvider): Promise<GetAddressesReturnType | undefined> => {
  try {
    const walletClient = createWalletClient({
      chain: optimismSepolia,
      transport: custom(provider),
    })
    const address = await walletClient.getAddresses()
    return address
  } catch (e) {
    console.error(e)
  }
}

export const publicClient = (provider: IProvider): any => {
  try {
    return createPublicClient({
      chain: optimismSepolia,
      transport: custom(provider),
    })
  } catch (e) {
    console.error(e)
  }
}

export const privateClient = (provider: IProvider): any => {
  try {
    return createWalletClient({
      chain: optimismSepolia,
      transport: custom(provider),
    })
  } catch (e) {
    console.error(e)
  }
}
