// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts-upgradeable/access/Ownable2StepUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "solidity-trigonometry/src/Trigonometry.sol";
import "hardhat/console.sol";

import "./ISeekerStatsOracle.sol";

contract SeekerStatsOracle is ISeekerStatsOracle, Initializable, Ownable2StepUpgradeable, ERC165 {
    /**
     * @notice The oracle account. This contract accepts any attestations of
     * Seeker power that have been signed by this account.
     */
    address public SeekerStatsOracleAccount;

    /**
     * @notice Tracks the set of Seeker Stats and Rank with Seeker ID
     */
    mapping(uint256 => Seeker) public seekers;

    /**
     * @notice Holds the angle used for coverage calculation in radians
     */
    int256 private angle = Trigonometry.sin(((Trigonometry.TWO_PI / 6) + Trigonometry.TWO_PI));

    event SeekerStatsUpdated(
        uint256 indexed seekerId,
        uint256 attr_chip,
        uint256 attr_durability,
        uint256 attr_sensors,
        uint256 attr_cores,
        uint256 attr_storage,
        uint256 attr_reactor
    );

    error OracleCannotBeZeroAddress();
    error SeekerProofIsEmpty();
    error UnauthorizedRegisterSeekerStatsCall();
    error InvalidSignatureForSeekerProof();
    error SeekerNotRegistered(uint256 seekerId);

    function initialize(address _seekerStatsOracleAccount) external initializer {
        if (_seekerStatsOracleAccount == address(0)) {
            revert OracleCannotBeZeroAddress();
        }

        Ownable2StepUpgradeable.__Ownable2Step_init();

        SeekerStatsOracleAccount = _seekerStatsOracleAccount;
    }

    /**
     * @notice Returns true if the contract implements the interface defined by
     * `interfaceId` from ERC165.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return
            interfaceId == type(ISeekerStatsOracle).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @notice Sets the oracle account.
     * @param _seekerStatsOracleAccount The oracle account.
     */
    function setOracle(address _seekerStatsOracleAccount) external onlyOwner {
        if (_seekerStatsOracleAccount == address(0)) {
            revert OracleCannotBeZeroAddress();
        }
        SeekerStatsOracleAccount = _seekerStatsOracleAccount;
    }

    /**
     * @notice Returns true if the oracle account signed the proof message for the given seeker.
     * @param seeker The object containing the seekers statistics.
     * @param signature The signature of the seekers proof message, signed by the oracle account.
     */
    function validateSeekerStatsProof(
        Seeker calldata seeker,
        bytes calldata signature
    ) internal view returns (bool) {
        if (SeekerStatsOracleAccount == address(0)) {
            revert OracleCannotBeZeroAddress();
        }

        bytes memory proof = createStatsMessage(seeker);
        bytes32 ecdsaHash = ECDSA.toEthSignedMessageHash(proof);
        address signerAddress = ECDSA.recover(ecdsaHash, signature);
        if (signerAddress == SeekerStatsOracleAccount) {
            return true;
        } else {
            return false;
        }
    }

    /**
     * @notice Creates a proofing message unique to the provided seeker.
     * @param seeker The object containing the seekers statistics.
     */
    function createStatsMessage(Seeker calldata seeker) public pure returns (bytes memory) {
        return
            abi.encodePacked(
                seeker.seekerId,
                seeker.rank,
                seeker.attr_chip,
                seeker.attr_durability,
                seeker.attr_sensors,
                seeker.attr_cores,
                seeker.attr_storage,
                seeker.attr_reactor
            );
    }

    function registerSeekerRestricted(Seeker calldata seeker) external {
        if (msg.sender != SeekerStatsOracleAccount) {
            revert UnauthorizedRegisterSeekerStatsCall();
        }

        seekers[seeker.seekerId] = seeker;
        emit SeekerStatsUpdated(
            seeker.seekerId,
            seeker.attr_chip,
            seeker.attr_durability,
            seeker.attr_sensors,
            seeker.attr_cores,
            seeker.attr_storage,
            seeker.attr_reactor
        );
    }

    /**
     * @notice Registers a seeker
     * @param seeker The object containing the seekers statistics.
     * @param signature The signature of the seekers proof message, signed by the oracle account.
     */
    function registerSeeker(Seeker calldata seeker, bytes calldata signature) external {
        bool valid = validateSeekerStatsProof(seeker, signature);
        if (!valid) {
            revert InvalidSignatureForSeekerProof();
        }

        seekers[seeker.seekerId] = seeker;
        emit SeekerStatsUpdated(
            seeker.seekerId,
            seeker.attr_chip,
            seeker.attr_durability,
            seeker.attr_sensors,
            seeker.attr_cores,
            seeker.attr_storage,
            seeker.attr_reactor
        );
    }

    /**
     * @notice Calculates the coverage score for the given seekers
     * @param seekersList A list containing seekers, will revert if any seeker is not registered.
     */
    function calculateAttributeCoverage(
        Seeker[] calldata seekersList
    ) external view returns (int256) {
        int256 sumCoverage = 0;
        Seeker memory defaultSeeker;

        for (uint256 i = 0; i < seekersList.length; i++) {
            Seeker memory seeker = seekersList[i];
            Seeker memory registeredSeeker = seekers[seeker.seekerId];
            if (keccak256(abi.encode(registeredSeeker)) == keccak256(abi.encode(defaultSeeker))) {
                revert SeekerNotRegistered(seeker.seekerId);
            }

            sumCoverage += (int256(seeker.attr_reactor) * angle * int256(seeker.attr_cores)) / 2;
            sumCoverage +=
                (int256(seeker.attr_cores) * angle * int256(seeker.attr_durability)) /
                2;
            sumCoverage +=
                (int256(seeker.attr_durability) * angle * int256(seeker.attr_sensors)) /
                2;
            sumCoverage += (int256(seeker.attr_sensors) * angle * int256(seeker.attr_storage)) / 2;
            sumCoverage += (int256(seeker.attr_storage) * angle * int256(seeker.attr_chip)) / 2;
            sumCoverage += (int256(seeker.attr_chip) * angle * int256(seeker.attr_reactor)) / 2;
        }

        return sumCoverage;
    }
}
