// contracts/Land.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./Land.sol";

contract Registry {
    // Track lands contract
    mapping(string => address) private lands;
    // Count tokenized lands
    uint256 private tokenizedLands;

    // Land events
    event LandCreated(string titleNo, uint size);

    // Error
    error LandAlreadyExists(string titleNo);
    
    // Register land
    function register(string memory titleNo_, string memory symbol_, address owner_, uint256 size_, uint256 registration_) public returns (Land landAddress) {
        // Don't mint same land
        require(lands[titleNo_] == address(0), LandAlreadyExists(titleNo_));

        tokenizedLands += 1;
        Land land = new Land(titleNo_, symbol_, owner_, size_, tokenizedLands, registration_);
        lands[titleNo_] = address(land);
        emit LandCreated(titleNo_, size_);
        return land;
    }

    // Total tokenized properties
    function countTokenizedLands() public view returns (uint256) {
        return tokenizedLands;
    }

    // Get land ERC721 contract
    function getLandERC721Contract(string memory titleNo_) public view returns (address) {
        return lands[titleNo_];
    }
}
