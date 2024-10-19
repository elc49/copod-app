import { expect } from "chai"
import hre, { ethers } from "hardhat"
import { Signer } from "ethers"
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs"

let signers: Signer[]
let registryContract: any

describe("Registry", () => {
  before(async() => {
    signers = await ethers.getSigners()
    registryContract = await hre.ethers.deployContract("CopodRegistry")
  })

  it("Register land", async () => {
    let owner = signers[0]
    await registryContract.addLand("FE/E32", "HA", await owner.getAddress(), 34, 4842)

    let landContract: any = (await ethers.getContractFactory("Land")).attach(await registryContract.getLandERC721Contract("FE/E32"))
    let land = await landContract.getLand()
    let count = await landContract.balanceOf(await owner.getAddress())
    let countLands = await registryContract.getCountTokenizedLands()

    expect(land.titleNo).to.be.equal("FE/E32")
    expect(land.symbol).to.be.equal("HA")
    expect(count).to.be.equal(1)
    expect(countLands).to.be.equal(1)
  })

  it("Dont't register same land parcel", async() => {
    let owner = signers[0]
    await expect(
      registryContract.addLand("FE/E32", "HA", await owner.getAddress(), 34, 482)
    ).to.be.reverted
  })

  it("Grant land usage rights", async() => {
    let tenant = signers[1]
    await expect(
      registryContract.setUsageRights("FE/E32", 32, 38203, 16000, await tenant.getAddress())
    ).to.emit(registryContract, "GrantLandUsageRights").withArgs("FE/E32", anyValue, anyValue, anyValue)
  })
})
