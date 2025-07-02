// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract TaxOracle {
    mapping(address => string) public merchantCategory;
    mapping(string => uint256) public taxRates;

    constructor() {
        taxRates["luxury"] = 75;
        taxRates["essential"] = 0;
    }

    function setMerchant(address merchant, string memory category) external {
        merchantCategory[merchant] = category;
    }

    function getTaxRate(address merchant) external view returns (uint256) {
        return taxRates[merchantCategory[merchant]];
    }
}
