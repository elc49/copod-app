// contracts/Land.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

// Land parcel
struct LandDetails {
    string titleNo;
    uint256 size;
    string symbol;
    address payable owner;
    uint256 tokenId;
}

contract Land is ERC721, Ownable {
    using Math for uint256;

    // Land titles to details- title no. are unique
    LandDetails land;

    // Events
    event GrantSize(address owner, uint256 size);
    event ReclaimSize(address owner, uint256 size);

    // Errors
    error TrySubtracting(uint256 size);
    error TryAdd(uint256 size);

    constructor(string memory titleNo_, string memory symbol_, address owner_, uint256 size_, uint256 tokenId_) ERC721(titleNo_, symbol_) Ownable(owner_) {
        land = LandDetails(titleNo_, size_, symbol_, payable(owner_), tokenId_);
        // Transfer titleNo_ to owner_
        _safeMint(owner_, tokenId_);
    }

    // Get land
    function getLand() public view returns (LandDetails memory) {
        return land;
    }

    // Grant from land space
    function grantSize(uint256 size_, address owner_) public returns (bool) {
        require(owner_ == ownerOf(land.tokenId), OwnableUnauthorizedAccount(owner_));
        (bool success, uint256 result) =  land.size.trySub(size_);
        require(success, TrySubtracting(size_));
        // Update new land size
        land.size = result;
        emit GrantSize(owner_, size_);
        return true;
    }

    // Reclaim granted land space
    function reclaimSize(uint256 size_, address owner_) public returns (bool) {
        require(owner_ == ownerOf(land.tokenId), OwnableUnauthorizedAccount(owner_));
        (bool success, uint256 result) = land.size.tryAdd(size_);
        // Update new land size
        land.size = result;
        emit ReclaimSize(owner_, size_);
        return true;
    }
}

