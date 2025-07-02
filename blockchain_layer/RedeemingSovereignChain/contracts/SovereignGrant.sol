// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract SovereignGrant {
    mapping(address => uint256) public balances;

    function issueGrant(address recipient, uint256 amount) public {
        balances[recipient] += amount;
    }

    function balanceOf(address recipient) public view returns (uint256) {
        return balances[recipient];
    }
}
