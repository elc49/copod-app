import { buildModule } from "@nomicfoundation/hardhat-ignition/modules"

export default buildModule("Registry", m => {
  const land = m.contract("Registry")
  return { land }
})
