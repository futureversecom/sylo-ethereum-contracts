// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import "./ISeekerStakingManager.sol";
import "./SeekerStatsOracle.sol";

contract SeekerStakingManager is
    ISeekerStakingManager,
    Initializable,
    Ownable2StepUpgradeable,
    ERC165
{
    /**
     * @notice ERC721 contract for bridged Seekers. Used for verifying ownership
     * of a seeker.
     */
    IERC721 public rootSeekers;

    /**
     * @notice SeekerStatsOracle contract
     */
    SeekerStatsOracle public oracle;

    /**
     * @notice mapping to track staked seekers by seeker ID
     */
    mapping(uint256 => StakedSeeker) public stakedSeekersById;

    /**
     * @notice mapping to track staked seekers by node address
     */
    mapping(address => uint256[]) public stakedSeekersByNode;

    /**
     * @notice mapping to track staked seekers by user address
     */
    mapping(address => uint256[]) public stakedSeekersByUser;

    /**
     * @notice default staked seeker
     */
    StakedSeeker public defaultStakedSeeker;

    error NodeAddressCannotBeNil();
    error FromNodeAddressCannotBeNil();
    error ToNodeAddressCannotBeNil();
    error SeekerProofIsEmpty();
    error InvalidSignatureForSeekerProof();
    error RootSeekersCannotBeZeroAddress();
    error SeekerStatsOracleCannotBeZeroAddress();
    error SenderAccountMustOwnSeekerId();
    error SeekerNotYetStaked();
    error SeekerNotStakedToNode();
    error SeekerNotStakedBySender();
    error CannotTransferSeekerToSameNode();

    function initialize(IERC721 _rootSeekers, SeekerStatsOracle _oracle) external initializer {
        if (address(_rootSeekers) == address(0)) {
            revert RootSeekersCannotBeZeroAddress();
        }
        if (address(_oracle) == address(0)) {
            revert SeekerStatsOracleCannotBeZeroAddress();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

        rootSeekers = _rootSeekers;
        oracle = _oracle;
    }

    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return
            interfaceId == type(ISeekerStakingManager).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @notice Stakes a seeker
     * @param node Address of node to stake seeker against
     * @param seeker The object containing the seekers statistics
     * @param seekerStatsProof The signature of the seekers proof message, signed by the oracle account.
     */
    function stakeSeeker(
        address node,
        SeekerStatsOracle.Seeker calldata seeker,
        bytes calldata seekerStatsProof
    ) external {
        if (node == address(0)) {
            revert NodeAddressCannotBeNil();
        }
        if (rootSeekers.ownerOf(seeker.seekerId) != msg.sender) {
            revert SenderAccountMustOwnSeekerId();
        }
        if (!oracle.isSeekerRegistered(seeker)) {
            if (seekerStatsProof.length == 0) {
                revert SeekerProofIsEmpty();
            }
            oracle.registerSeeker(seeker, seekerStatsProof);
        }

        stakedSeekersById[seeker.seekerId] = StakedSeeker({
            seekerId: seeker.seekerId,
            node: node,
            user: msg.sender
        });
        stakedSeekersByNode[node].push(seeker.seekerId);
        stakedSeekersByUser[msg.sender].push(seeker.seekerId);
    }

    /**
     * @notice Transfers a staked seeker between nodes
     * @param fromNode Address of node to transfer the staked seeker from
     * @param toNode Address of node to transfer the staked seeker to
     * @param seekerId Seeker ID of the staked seeker to transfer
     */
    function transferStakedSeeker(address fromNode, address toNode, uint256 seekerId) external {
        if (fromNode == address(0)) {
            revert FromNodeAddressCannotBeNil();
        }
        if (toNode == address(0)) {
            revert ToNodeAddressCannotBeNil();
        }
        if (fromNode == toNode) {
            revert CannotTransferSeekerToSameNode();
        }
        if (rootSeekers.ownerOf(seekerId) != msg.sender) {
            revert SenderAccountMustOwnSeekerId();
        }
        if (
            keccak256(abi.encode(stakedSeekersById[seekerId])) ==
            keccak256(abi.encode(defaultStakedSeeker))
        ) {
            revert SeekerNotYetStaked();
        }
        if (stakedSeekersById[seekerId].user != msg.sender) {
            revert SeekerNotStakedBySender();
        }
        if (stakedSeekersById[seekerId].node != fromNode) {
            revert SeekerNotStakedToNode();
        }

        for (uint256 i = 0; i < stakedSeekersByNode[fromNode].length; i++) {
            if (stakedSeekersByNode[fromNode][i] == seekerId) {
                stakedSeekersByNode[fromNode][i] = stakedSeekersByNode[fromNode][
                    stakedSeekersByNode[fromNode].length - 1
                ];

                stakedSeekersByNode[fromNode].pop();

                stakedSeekersByNode[toNode].push(seekerId);
            }
        }

        stakedSeekersById[seekerId].node = toNode;
    }

    /**
     * @notice Unstake a seeker
     * @param node Address of node to unstake seeker from
     * @param seekerId Seeker ID of staked seeker
     */
    function unstakeSeeker(address node, uint256 seekerId) external {
        if (node == address(0)) {
            revert FromNodeAddressCannotBeNil();
        }
        if (rootSeekers.ownerOf(seekerId) != msg.sender) {
            revert SenderAccountMustOwnSeekerId();
        }
        if (
            keccak256(abi.encode(stakedSeekersById[seekerId])) ==
            keccak256(abi.encode(defaultStakedSeeker))
        ) {
            revert SeekerNotYetStaked();
        }
        if (stakedSeekersById[seekerId].user != msg.sender) {
            revert SeekerNotStakedBySender();
        }
        if (stakedSeekersById[seekerId].node != node) {
            revert SeekerNotStakedToNode();
        }

        for (uint256 i = 0; i < stakedSeekersByNode[node].length; i++) {
            if (stakedSeekersByNode[node][i] == seekerId) {
                stakedSeekersByNode[node][i] = stakedSeekersByNode[node][
                    stakedSeekersByNode[node].length - 1
                ];

                stakedSeekersByNode[node].pop();
                break;
            }
        }

        for (uint256 i = 0; i < stakedSeekersByUser[msg.sender].length; i++) {
            if (stakedSeekersByUser[msg.sender][i] == seekerId) {
                stakedSeekersByUser[msg.sender][i] = stakedSeekersByUser[msg.sender][
                    stakedSeekersByUser[msg.sender].length - 1
                ];

                stakedSeekersByUser[msg.sender].pop();
                break;
            }
        }

        delete stakedSeekersById[seekerId];
    }

    /**
     * @notice Get a staked seeker by node address
     * @param node Address of node seeker is staked against
     */
    function getStakedSeekersByNode(address node) external view returns (uint256[] memory) {
        return stakedSeekersByNode[node];
    }

    /**
     * @notice Get a staked seeker by user address
     * @param user Address of user that staked seeker
     */
    function getStakedSeekersByUser(address user) external view returns (uint256[] memory) {
        return stakedSeekersByUser[user];
    }
}
