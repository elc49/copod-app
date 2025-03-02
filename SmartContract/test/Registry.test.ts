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
const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000"

describe("Registry", () => {
  before(async () => {
    signers = await ethers.getSigners()
    owner = signers[0]
    registryContract = await hre.ethers.deployContract("Registry")
  })

  describe("Register", async () => {
    it("Register", async () => {
      const registration = Date.parse("2002-01-15")
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

    it("Don't register same land parcel", async () => {
      await expect(
        registryContract.register(land.titleNo, land.unit, await owner.getAddress(), 34, Date.parse("2002-01-15"))
      ).to.be.reverted
    })

    it("Query land title", async () => {
      const address = await registryContract.getLandERC721Contract(land.titleNo)

      expect(address).to.not.be.equal(ZERO_ADDRESS)
    })

    it("Query non-existent land title", async () => {
      const address = await registryContract.getLandERC721Contract("FE/E32")

      expect(address).to.be.equal(ZERO_ADDRESS)
    })

    it("Query land size divisibility decimal", async () => {
      const address = await registryContract.getLandERC721Contract(land.titleNo)
      landContract = (await ethers.getContractFactory("Land")).attach(address)
      const decimals = await landContract.decimals()
      
      expect(decimals).to.be.equal(18)
    })
  })
})
