import { buildModule } from "@nomicfoundation/hardhat-ignition/modules"

export default buildModule("Land", m => {
  const land = m.contract("Land", ["titleNo", "ha", "0xcA17C601f9b6a1Ad58C7051B66e1342E7BC487A7", 34])
  return { land }
})
