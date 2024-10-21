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
    event ReclaimUsageRights(uint256 size, address tenant, string titleNo);

    // Error
    error LandAlreadyExists(string titleNo);
    error UnavailableLandSpace(string titleNo, uint256 size);
    error NotAuthorized(address caller);
    error GrantSize(uint256 size);
    error NoTokenizedLand(string titleNo);
    
    // Register land
    function register(string memory titleNo_, string memory symbol_, address owner_, uint256 size_, uint256 tokenId_) public returns (Land landAddress) {
        // Don't mint same land
        require(lands[titleNo_] == address(0), LandAlreadyExists(titleNo_));

        Land land = new Land(titleNo_, symbol_, owner_, size_, tokenId_);
        lands[titleNo_] = address(land);
        tokenizedLands += 1;
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

    // Grant usage rights
    // TODO: should be authorized/authenticated by owner/tenant
    function grantLandUsageRights(string memory titleNo_, uint256 size_, uint256 duration_, uint256 cost_, address tenant_, address owner_) public {
        // Get land details
        require(lands[titleNo_] != address(0), NoTokenizedLand(titleNo_));
        Land l = Land(lands[titleNo_]);
        LandDetails memory land = l.getLand();

        // Validate caller
        require(owner_ == land.owner, NotAuthorized(owner_));

        // Validate land space
        require(size_ <= land.size, UnavailableLandSpace(titleNo_, size_));

        // Grant usage size
        (bool success) = l.grantSize(size_, owner_);
        require(success, GrantSize(size_));

        usage[tenant_][titleNo_] = UsageRight(size_, duration_, payable(tenant_), cost_, titleNo_);
        emit GrantLandUsageRights(titleNo_, size_, duration_, tenant_);
    }

    // Reclaim back rights - only callable by land owner
    // TODO: only callable past duration
    function reclaimUsageRights(address tenant_, string memory titleNo_) public {
        // TODO: reclaim back rights- size
        emit ReclaimUsageRights(32, tenant_, titleNo_);
    }
}
