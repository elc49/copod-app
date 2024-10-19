// contracts/Land.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

// Land parcel
struct LandDetails {
    string titleNo;
    uint256 size;
    string symbol;
    address payable owner;
    uint256 tokenId;
}

contract Land is ERC721, Ownable {
    // Land titles to details- title no. are unique
    LandDetails land;

    constructor(string memory titleNo_, string memory symbol_, address owner_, uint256 size_, uint256 tokenId_) ERC721(titleNo_, symbol_) Ownable(owner_) {
        land = LandDetails(titleNo_, size_, symbol_, payable(owner_), tokenId_);
        // Transfer titleNo_ to owner_
        _safeMint(owner_, tokenId_);
    }

    function getLand() public view returns (string memory titleNo, string memory symbol, uint256 size) {
        return (land.titleNo, land.symbol, land.size);
    }
}

