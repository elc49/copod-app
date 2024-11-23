import { CHAIN_NAMESPACES } from "@web3auth/base";
import { optimismSepolia } from "viem/chains";
import { numberToHex } from "viem";

const chainConfig = {
  chainNamespace: CHAIN_NAMESPACES.EIP155,
  chainId: numberToHex(optimismSepolia.id),
  rpcTarget: process.env.NEXT_PUBLIC_BLOCKCHAIN_RPC_API!,
  displayName: "Ethereum",
  blockExplorerUrl: "https://sepolia-optimism.etherscan.io/",
  ticker: "ETH",
  tickerName: "Ethereum",
  logo: "https://cryptologos.cc/logos/ethereum-eth-logo.png",
};

export default chainConfig;
