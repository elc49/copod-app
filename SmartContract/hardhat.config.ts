import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@nomicfoundation/hardhat-ignition-ethers"
import "@nomicfoundation/hardhat-ethers"
import "@nomicfoundation/hardhat-chai-matchers"
import "@nomicfoundation/hardhat-network-helpers"
import { vars } from "hardhat/config";

const INFURA_API_KEY = vars.get("INFURA_API_KEY", "6c43b5c347ff65879365")
const SEPOLIA_PRIVATE_KEY = vars.get("SEPOLIA_PRIVATE_KEY", "1699516346e0e652d6c43b5c347ff65879365ede5aaca0b7b3004efb992c4f0b")

const config: HardhatUserConfig = {
  solidity: "0.8.27",
  networks: {
    sepolia: {
      url: `https://optimism-sepolia.infura.io/v3/${INFURA_API_KEY}`,
      accounts: [SEPOLIA_PRIVATE_KEY],
    },
  },
};

export default config;
