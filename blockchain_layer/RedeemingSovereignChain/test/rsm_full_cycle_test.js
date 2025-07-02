const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Redeeming Sovereign Model – Full Cycle", function () {
  let SovereignGrant;
  let grant;
  let owner, addr1;

  beforeEach(async function () {
    [owner, addr1] = await ethers.getSigners();

    // ✅ Get the contract factory
    SovereignGrant = await ethers.getContractFactory("SovereignGrant");

    // ✅ Deploy the contract
    grant = await SovereignGrant.deploy();
  });

  it("should issue grant and allocate savings correctly", async function () {
    const totalGrant = ethers.parseEther("280000"); // <-- FIXED

    // ✅ Simulate issuing a grant
    await grant.issueGrant(addr1.address, totalGrant);

    // ✅ Check balance
    const balance = await grant.balanceOf(addr1.address);
    expect(balance).to.equal(totalGrant);
  });
});
