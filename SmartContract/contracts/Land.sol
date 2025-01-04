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
    uint256 registration;
}

contract Land is ERC721, Ownable {
    // Land details
    LandDetails private land;

    constructor(string memory titleNo_, string memory symbol_, address owner_, uint256 size_, uint256 tokenId_, uint256 registration_) ERC721(titleNo_, symbol_) Ownable(owner_) {
        land = LandDetails(titleNo_, size_, symbol_, payable(owner_), tokenId_, registration_);
        // Transfer tokenId_ to owner_
        _safeMint(owner_, tokenId_);
    }

    // Get land
    function getLand() public view returns (LandDetails memory) {
        return land;
    }
}
