// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.28;

contract DebtIncinerator {
    mapping(address => uint256) public loans;
    uint256 public totalBurned;

    event DebtRedeemed(address indexed citizen, uint256 amount);

    function repayLoan() external {
        uint256 amount = loans[msg.sender];
        require(amount > 0, "No loan to repay");
        loans[msg.sender] = 0;
        totalBurned += amount;
        emit DebtRedeemed(msg.sender, amount);
    }

    function borrow(uint256 amount) external {
        loans[msg.sender] += amount;
    }
}
