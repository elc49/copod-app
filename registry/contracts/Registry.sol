// contracts/Land.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "./Land.sol";

// Land usage
struct UsageRight {
    uint size;
    uint duration;
    address payable tenant;
    uint256 cost;
    string titleNo;
}

contract Registry {
    // Lands
    mapping(string => address) private lands;
    // Land title usage rights
    mapping(address => mapping(string => UsageRight)) private usage;
    // Count tokenized lands
    uint256 private tokenizedLands;

    // Land events
    event LandCreated(string titleNo, uint size);
    event GrantLandUsageRights(string titleNo, uint256 size, uint256 duration, address tenant);
    event ReclaimBackRights(uint256 size, address tenant, string titleNo);

    // Error
    error LandAlreadyExists(string titleNo);
    
    // Register land
    function addLand(string memory titleNo_, string memory symbol_, address owner_, uint256 size_, uint256 tokenId_) public returns (Land landAddress) {
        // Don't mint same land
        require(lands[titleNo_] == address(0), LandAlreadyExists(titleNo_));

        Land land = new Land(titleNo_, symbol_, owner_, size_, tokenId_);
        lands[titleNo_] = address(land);
        tokenizedLands += 1;
        emit LandCreated(titleNo_, size_);
        return land;
    }

    // Total tokenized properties
    function getCountTokenizedLands() public view returns (uint256) {
        return tokenizedLands;
    }

    // Get land ERC721 contract
    function getLandERC721Contract(string memory titleNo_) public view returns (address) {
        return lands[titleNo_];
    }

    // Grant usage rights
    function setUsageRights(string memory titleNo_, uint256 size_, uint256 duration_, uint256 cost_, address tenant_) public {
        usage[tenant_][titleNo_] = UsageRight(size_, duration_, payable(tenant_), cost_, titleNo_);
        // TODO: properly update land size
        emit GrantLandUsageRights(titleNo_, size_, duration_, tenant_);
    }

    // Reclaim back rights - only callable by land owner
    function reclaimBackRights(address tenant_, string memory titleNo_) public {
        // TODO: reclaim back rights- size
        emit ReclaimBackRights(32, tenant_, titleNo_);
    }
}
