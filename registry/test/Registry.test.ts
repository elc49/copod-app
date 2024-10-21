import { expect } from "chai"
import hre, { ethers } from "hardhat"
import { Signer } from "ethers"
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs"
import { time } from "@nomicfoundation/hardhat-network-helpers"

interface Land {
  titleNo: string
  size: number
  unit: string
  tokenId: number
}

let signers: Signer[], owner: Signer, tenant: Signer
let registryContract: any, landContract: any
const land: Land = {
  titleNo: "FE/E32/HZ",
  size: 32,
  unit: "HA",
  tokenId: 4842,
}
const timeInFuture = Math.floor((Date.now() + (60*60*24))/1000)

describe("Registry", () => {
  before(async () => {
    signers = await ethers.getSigners()
    owner = signers[0]
    tenant = signers[1]
    registryContract = await hre.ethers.deployContract("Registry")
  })

  describe("register", async () => {
    it("register success", async () => {
      await registryContract.register(land.titleNo, land.unit, await owner.getAddress(), land.size, 4842)

      landContract = (await ethers.getContractFactory("Land")).attach(await registryContract.getLandERC721Contract(land.titleNo))
      const result = await landContract.getLand()
      const count = await landContract.balanceOf(await owner.getAddress())
      const countLands = await registryContract.countTokenizedLands()

      expect(result.titleNo).to.be.equal(land.titleNo)
      expect(result.symbol).to.be.equal(land.unit)
      expect(count).to.be.equal(1)
      expect(countLands).to.be.equal(1)
    })

    it("Dont't register same land parcel", async () => {
      await expect(
        registryContract.register(land.titleNo, land.unit, await owner.getAddress(), 34, 482)
      ).to.be.reverted
    })
  })

  describe("grantLandUsageRights", async () => {
    it("not land owner grant usage", async () => {
      await expect(
        registryContract.grantLandUsageRights(land.titleNo, 21, timeInFuture, 25000, await owner.getAddress(), await tenant.getAddress())
      ).to.be.revertedWithCustomError(registryContract, "NotAuthorized")
    })

    it("land should cover requested usage size", async () => {
      await expect(
        registryContract.grantLandUsageRights(land.titleNo, 34, timeInFuture, 25000, await tenant.getAddress(), await owner.getAddress())
      ).to.be.revertedWithCustomError(registryContract, "UnavailableLandSpace")
    })

    it("request usage to verified land", async () => {
      await expect(
        registryContract.grantLandUsageRights("ER/34", 21, timeInFuture, 25000, await tenant.getAddress(), await owner.getAddress())
      ).to.be.revertedWithCustomError(registryContract, "NoTokenizedLand")
    })

    it("grant rights success", async () => {
      await expect(
        registryContract.grantLandUsageRights(land.titleNo, 21, timeInFuture, 25000, await tenant.getAddress(), await owner.getAddress())
      ).to.emit(registryContract, "GrantLandUsageRights").withArgs(land.titleNo, anyValue, anyValue, anyValue)

      const l = await landContract.getLand()
      expect(l.size).to.be.equal(11)
    })
  })

  describe("claimUsageRights", async () => {
    it("return true", async () => {
      expect(
        await registryContract.claimUsageRights(await tenant.getAddress(), land.titleNo)
      ).to.be.true
    })

    it("return false", async () => {
      const futureTime = Math.floor((Date.now() + (60*60*24))/1000)
      await time.increase(futureTime)

      expect(
        await registryContract.claimUsageRights(await tenant.getAddress(), land.titleNo)
      ).to.be.false
    })
  })
})
