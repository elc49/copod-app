import { expect } from "chai"
import hre, { ethers } from "hardhat"
import { Signer } from "ethers"

interface Land {
  titleNo: string
  size: number
  unit: string
}

let signers: Signer[], owner: Signer
let registryContract: any, landContract: any
const land: Land = {
  titleNo: "FE/E32/HZ",
  size: 32,
  unit: "HA",
}

describe("Registry", () => {
  before(async () => {
    signers = await ethers.getSigners()
    owner = signers[0]
    registryContract = await hre.ethers.deployContract("Registry")
  })

  describe("register", async () => {
    it("register success", async () => {
      const registration = Date.now()
      await registryContract.register(land.titleNo, land.unit, await owner.getAddress(), land.size, registration)

      landContract = (await ethers.getContractFactory("Land")).attach(await registryContract.getLandERC721Contract(land.titleNo))
      const result = await landContract.getLand()
      const count = await landContract.balanceOf(await owner.getAddress())
      const countLands = await registryContract.countTokenizedLands()

      expect(result.titleNo).to.be.equal(land.titleNo)
      expect(result.symbol).to.be.equal(land.unit)
      expect(count).to.be.equal(1)
      expect(countLands).to.be.equal(1)
      expect(result.registration).to.be.equal(registration)
    })

    it("Dont't register same land parcel", async () => {
      await expect(
        registryContract.register(land.titleNo, land.unit, await owner.getAddress(), 34, Date.now())
      ).to.be.reverted
    })

    it("count land usages", async () => {
      const count = await registryContract.getLandUsagesCount(land.titleNo)

      expect(count).to.be.equal(0)
    })
  })
})
