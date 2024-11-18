import { WEB3AUTH_NETWORK, IProvider } from "@web3auth/base";
import { Web3AuthOptions } from "@web3auth/modal";
import { chainConfig } from "@/blockchains/chains";
import { EthereumPrivateKeyProvider } from "@web3auth/ethereum-provider";
import { ethers } from "ethers";

const clientId = process.env.NEXT_PUBLIC_WEB3_AUTH_CLIENT_ID!

const privateKeyProvider = new EthereumPrivateKeyProvider({
  config: { chainConfig },
})

export const getWeb3AuthOptions = (): Web3AuthOptions => {
  return {
    clientId,
    web3AuthNetwork: WEB3AUTH_NETWORK.SAPPHIRE_DEVNET,
    privateKeyProvider,
  }
}

export const getAccounts = async (provider: IProvider): Promise<any> => {
  try {
    const ethersProvider = new ethers.BrowserProvider(provider)
    const signer = await ethersProvider.getSigner()

    return signer.getAddress()
  } catch (error) {
    console.error(error)
    return error
  }
}
