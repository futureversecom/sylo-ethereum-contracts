# Sylo Network Protocol Technical Specification

## Introduction

The Sylo Network Protocol is a suite of smart contracts deployed on Ethereum, which govern a system of compensating service providers via micro-probabilistic payments. This protocol will be used to power the [Event Relay Protocol] (todo link) that allows for trustless, decentralized communication. This document is a technical specification of the data types and contract calls currently present in the system. For a more general overview of the protocol, refer to these documents:
  - The [whitepaper] -- TODO plugin link to whitepaper

## Mechanics Overview

Quick discussion about micro payment tickets, event relay protocol, and epochs.

## Users

The Sylo Network consists of multiple types of users interacting with the smart contracts.
  - **Nodes**: Users that wish to operate and maintain a [Sylo Node] (todo: Link to Sylo Node here) that will support the Event Relay Protocol. Running a Node will allow for compensation via redeeming micro-payment tickets. Node operators will be required to have `SYLO` tokens staked against their node in order to participate in the network, and the amount of work (thus the amount of compensation) that a Node receives will be based on their proportion of stake relative to  other Nodes within the network. See [stake weighted scan] (todo plugin stake weighted scan link) for details.
  - **Delegated Stakers**: Delegated Stakers (or delegators) are users that wish to participate in the Sylo Network and earn `SYLO` tokens without needing to run a Sylo Node themselves. These users can supply additional staked `SYLO` to an existing Node within the Network in order to increase the Node's potential for generating revenue. Delegated Stakers will be rewarded on a pro-rata basis.
  - **Senders**: Senders are uses which hold `SYLO` tokens and wish to utilize Sylo Nodes for their decentralized communication service. Senders must deposit `SYLO` tokens into both an `escrow` and `penalty` balance held within a smart contract. Nodes will be paid via these balances, and senders are required to maintain a healthy level of both `escrow` and `penalty` to be able to participate in the network.
  - **Receivers**: Receivers do not explicitly interact with the contracts but play a critical role in the Event Relay protocol. On receiving an event, receivers will reveal the necessary information in order for a Node to redeem a winning ticket and be compensated. See [here] (todo link) for more details on the ticketing mechanism.
  - **Sylo**: The Sylo team will be deploying the contracts to Ethereum and will have "ownership" of the contracts. Ownership allows certain privileged functions to be called on the contracts. These functions range from manually adjusting network parameters to making the call to initialize the next Epoch. These responsibilities be passed over to a DAO (todo link to DAO details?).

## Network Parameters

#### **epochDuration**

The minimum duration in blocks the next epoch will last for. Attempting to initialize
the next epoch if the current epoch's duration has not yet reached this value
will result in failure.

#### **defaultPayoutPercentage**

The payout percentage refers to the percentage of a ticket's face value that will be divvied out to a Node's delegated stakers. The remaining value is then given to the Node as a fee for providing Event Relay service.

Example:

If this value was set to `40%`, and a ticket's value was `1000 SOLO`. Then on redeeming a ticket, `400 SOLO` would be set aside for delegated stakers, and the remaining (`600 SOLO`) is given directly to the Node.

The `defaultPayoutPercentage` parameter is the default value for this. **Note**: For phase two, the default value is used for all Nodes and supersedes the `payoutPercentage` value set in a Node's listing.

Changes to this value will only take effect in the next epoch.

#### **faceValue**

The value in `SOLO` of a winning ticket.

Changes to this value will only take effect in the next epoch.

#### **baseLiveWinProb**

The probability of a ticket winning immediately after the ticket is issued.

Changes to this value will only take effect in the next epoch.

#### **expiredWinProb**

The probability of a ticket winning after the ticket's entire duration has elapsed.

Changes to this value will only take effect in the next epoch.

#### **ticketDuration**

The duration in blocks a ticket is alive for.

Changes to this value will only take effect in the next epoch.

#### **decayRate**

The rate at which a ticket's winning probability will decay over its lifetime, expressed as a percentage.

Example:

A decay rate of `80%` and a base live winning probability of `10%` indicates that at the block immediately before a ticket has expired, the ticket's winning probability will have decayed to `2%`.

#### **unlockDuration**

The duration in blocks that must elapse before either deposits or stake can be withdrawn (once the unlocking phase has begun).

#### **minimumStakeProportion**

The minimum amount of stake a Node must own for itself, expressed as a percentage of the Node's overall delegated stake. This requirement must always be must met whenever the Node unlock stake, or if other delegators attempt to add more stake to the Node. Failing to meet this requirement
will prevent the Node from participating in the next epoch.

Example:

A minimum stake proportion of `20%` indicates that the Node must own 20% of its total delegated stake. Thus if the stake total was `1000 SOLO`, then the must own at least `200 SOLO` to participate in the network.

## Smart Contract Specification

The Sylo Network Protocol contracts are written in Solidity and will initially be deployed to the Ethereum mainnet. The current system includes:
  - `SyloToken`: ERC20 contract for the Sylo Token which has already been deployed.
  - `SyloTicketing`: Contract that manages user deposits for payments,
  and implements the `redeem` function for redeeming winning tickets.
  - `StakingManager`: Tracks the amount of stake and the delegated stakers
  for each Node
  - `Directory`: Creates and manages a `Directory` structure every epoch based on the stake held by each stakee. The `Directory` is used as the backend
  for the stake-weighted scan function.
  - `RewardsManager`: Tracks rewards for each epoch when winning tickets are redeemed.
  - `EpochsManager`: Manages initializing of each epoch and stores the Network parameters for every epoch.
  - `Listings`: Stores a `Listing` struct for every Node

### Data Types

---

#### **Epoch**

Network parameters for the current epoch are saved into this structure by the `EpochsManager` contract when every new epoch is initialized.

| Field | Description |
|-------|-------------|
| iteration | A numerical value which is incremented when a new epoch is initialized. Also used as the epoch's identifier |
| startBlock | The block the epoch started |
| duration | The duration in blocks the epoch will last for |
| endBlock | The block the epoch ended. Initially set to 0 but will be updated when the next epoch is initialized |
| defaultPayoutPercentage | [See defaultPayoutPercentage](#-defaultPayoutPercentage) |
| faceValue | [See faceValue](#-faceValue) |
| baseLiveWinProb | [See baseLiveWinProb](#-baseLiveWinProb) |
| expiredWinProb | [See expiredWinProb](#-expiredWinProb) |
| ticketDuration | [See ticketDuration](#-ticketDuration) |
| decayRate | [See defaultPayoutPercentage](#-decayRate) |

#### **Ticket**

Tickets are created by senders at the client level and are given to Nodes as compensation for
providing event relay.

| Field | Description |
|-------|-------------|
| epochId | The id of the epoch the ticket was generated in |
| sender | Address of the sender |
| redeemer | Address of the redeemer (Usually the node) |
| generationBlock | The approximate block number the ticket was generated at |
| senderCommit | Hash of the secret random number of the sender |
| redeemerCommit | Hash of the secret random number of the redeemer |


#### **Stake**

The *Stake* datatype tracks a Node's current total managed stake, and each individual delegated stake entry.

| Field | Description |
|-------|-------------|
| stakeEntries | A mapping between each delegated staker and it's `StakeEntry` for the given stakee
| totalManagedStake | The sum of all delegated stake amounts for the stakee |

#### **StakeEntry**

A datatype that tracks a delegated staker's stake for a particular stakee.

| Field | Description |
|-------|-------------|
| amount | The amount of delegated stake in `SOLO` |
| updatedAt | The block number this stake entry was updated at |
| epochId | The epoch id of the epoch this stake entry was updated in |

#### **Directory**

A snapshot of all staking entries at the time it was created.

| Field | Description |
|-------|-------------|
| entries | An array of `DirectoryEntry` that is iterated over during `scan` (# todo link scan fn)
| stakes | A mapping of each stakee to their total stake |
| totalStake | The sum of all stakes |

#### **DirectoryEntry**

This datatype helps to make the `scan` implementation more efficient. A `DirectoryEntry` value is created as Nodes join the next epoch's directory. The entry includes a `boundary` value which is a sum of the current directory's total stake, and
the current Node's total stake. This entry is then pushed to the end of the `entries` array for the given directory.

| Field | Description |
|-------|-------------|
| stakee | The address of the stakee |
| boundary | The boundary value for this entry |


#### **RewardPool**

Each Node must initialize a reward pool for every epoch they wish to participate in. The reward pool will help track the portion of rewards from redeeming tickets that will be distributed to a Node's delegated stakers. The reward pool also tracks the [cumulative reward factor] (# todo link crf) (CRF) to make calculating the distribution more efficient.

| Field | Description |
|-------|-------------|
| stakersRewardTotal | The balance of the reward pool |
| initializedAt | The block number the reward pool was initialized |
| totalActiveStakes | Tracks the total active stake for this reward pool |
| initialCumulativeRewardFactor | The CRF at the time the reward pool was initialized |
| cumulativeRewardFactor | The ongoing CRF of the reward pool which is updated as tickets are redeemed |


#### **Listing**

Every Node must also have a `Listing` entry. The entry holds various network parameters which are configured by Nodes themselves.

| Field | Description |
|-------|-------------|
| multiAddr | The libp2p multi address of the Node. This is needed for clients to connect to their Node. Nodes should take care to ensure this value is correct and up to date |
| payoutPercentage | Percentage that of a redeemed tickets value that will be paid out to the Node's delegated stakers. **This value is currently unused and is superseded by the *defaultPayoutPercentage* network parameter for phase two**.
| minDelegatedStake | The minimum amount of stake a delegated staker must put forth |


### Functions

This section will detail the various contract calls that will be made over the lifetime of the Sylo Network.

---

#### Listings

#### *setListing*

Nodes are required to set their `Listing` entry to be able to stake and redeem tickets.

| Param | Description |
|-------|-------------|
| multiAddr | Sets the multi addr for the Node |
| minDelegatedStake | Sets the minimum delegated stake for the Node |

---

#### StakingManager

#### *addStake*

Called by both Nodes and delegators. This will transfer `SOLO` from the `msg.sender` to the `StakingManager` contract, and create or update a stake
entry for the specified stakee. Additionally it will also automatically claim any
outstanding stake rewards. This function will fail if the additional added stake
will cause the Node to own less than require [minimumStakeProportion] (# todo link here).

| Param | Description |
|-------|-------------|
| amount | The amount of stake to add in  `SOLO` |
| stakee | The address of the stakee |

#### *unlockStake*

Allows Node and delegators to set their stake for unlocking, which eventually will allow the stake to be withdrawn once the unlocking phase has ended. This removes the stake for consideration in the next epoch. If any stake was already in the unlocking phase, the amount of unlocking stake will instead be increased and the unlock duration will be reset.

| Param | Description |
|-------|-------------|
| amount | The amount of stake to unlock in  `SOLO` |
| stakee | The address of the stakee |

**Algorithm**

#### *cancelUnlocking*

Cancels stake that is in the unlocking phase and adds it back to the total managed stake for that stakee. The re-added stake can be utilized in the next epoch.

| Param | Description |
|-------|-------------|
| amount | The amount of unlocking stake to cancel in  `SOLO` |
| stakee | The address of the stakee |

##### *withdrawStake*

Returns stake that has finished unlocking back to the `msg.sender` account. This function will fail if the stake has not finished unlocking.

| Param | Description |
|-------|-------------|
| stakee | The address of the stakee |

---

#### Directory

#### *joinNextDirectory*

Called by Nodes as a prerequisite to participating in the Sylo Network for the next epoch. This function allows the stake delegated to a Node be used in the `scan` function. It will create and append a `DirectoryEntry` based on the sum of the total managed stake the Node has, plus any unclaimed staking rewards.

There are no explicit parameters for this function though only allows a Node to call this function once per epoch. It is in the Node's best interest to call this function near the end of the current epoch, in order to maximize the amount of unclaimed reward that can be used for the directory entry.

#### *scan*

Called by users of the Event Relay service to find the node associated with a given `point`. The `point` value is any 16 byte value, which is likely to be the hash of some user identifier. Hashing a user ID will create a psuedo-random value, which the `scan` function then maps to a value between 0 and the total stake in the current epoch's directory. This is then used in a binary search with the directory's entries, eventually returning the address of a Node. Node's with larger proportions of stake are more likely to be returned by the `scan` function.

| Param | Description |
|-------|-------------|
| point | A psuedo-random value |

**Returns**: Address of a Node

---

#### RewardsManager

#### *initializeRewardPool*

Called by Nodes as a prerequisite to participating in the Sylo Network for the next epoch. This function initializes and stores a new `RewardPool` entry for the next epoch for this
Node. It will calculate the `totalActiveStake` for this reward pool based on the sum of the total managed stake the Node has. The new reward pool will also read the `cumulativeRewardFactor` from the previous pool and begin tracking a new factor for the next epoch.

There are no explicit parameters for this function though only allows a Node to call this function once per epoch. It is in the Node's best interest to call this function near the end of the current epoch, in order to maximize the amount of unclaimed reward that can be used for the directory entry.

#### *claimStakeReward*

This function is called by Nodes and delegators when they wish to claim rewards that their stake has gained for them. This will utilize the current reward pool's CRF and the CRF at the time their stake became active to calculate the value of their reward. See [Cumulative Reward Factor] for details of the calculation though generally having a higher proportion of stake compared to other delegators will lead to a larger reward claim. Calling this function will prevent the user from being eligible to claim any further rewards until the next epoch begins

| Param | Description |
|-------|-------------|
| stakee | Address of the stakee the user wishes to claim against |

A public function `calculateStakerClaim` is exposed by the `RewardsManager` contract which allows users to understand the amount in `SOLO` gained if they were to call `claimStakingReward`. As this will also remove the users unclaimed reward from being used in the total active stake for the next epoch, users may wish to wait until the reward value is high enough to offset gas costs.

#### *claimNodeReward*

This is called by Node operators when they wish to withdraw rewards gained from operating a Node. The current value of this reward is a public field of the `RewardsManager` contract.

Node operators may wish to wait until the reward value is high enough to minimize the loss in earnings from gas costs.

---

#### EpochsManager

#### *initializeEpoch*

Sylo will take responsibility of calling this function every epoch to initialize the next epoch. Invoking this function will read the current set of network parameters and store it a new `Epoch` value. This function will fail if the current epoch has yet to end.

---

#### Ticketing

#### *depositEscrow*

This function is called by users that wish to utilize the Event Relay service provided by Node. This function transfers a specified amount of `SOLO` to be held in escrow by the Ticketing contract. When winning tickets are redeemed, the face value of the ticket will be paid out from the escrow.

| Param | Description |
|-------|-------------|
| amount | The amount in `SOLO` to deposit |
| account | The account the deposit will belong to. **Note**: The tokens are still transferred from the `msg.sender` account. |


##### *depositPenalty*

This function should be called in conjunction with `depositEscrow` to also hold a `penalty` amount in escrow. When winning tickets are redeemed, if the face value of a ticket is greater than the sender's escrow, then the penalty will be burned instead. This is to prevent certain economic attacks. [Further detail with regards to the economics of the Sylo Network can be found here] (todo link).

##### *unlockDeposits*

Moves both existing escrow and penalty values to an unlocking phase, which eventually allows withdrawal once the unlocking phase has completed. This function will fail if the user has already begun unlocking their deposits.

##### *lockDeposits*

This function essentially cancels the unlocking phase and allows the token to be used again as deposits.

##### *withdraw*

Once the unlocking phase has completed, this function can be called to transfer the tokens held in escrow back to the `msg.sender`.

##### *redeem*

`Redeem` should be called by the Node after completing an event relay and learning of the ticket sender's secret random value. The Node should only call this if it understands the ticket will win. The `Ticketing` contract exposes both `calculateWinningProbability` and `isWinningTicket` functions that can be used to determine if a ticket is winning, though the Node can also perform the calculation locally.

| Param | Description |
|-------|-------------|
| ticket | The `Ticket` (todo link ticket) issued by the sender |
| senderRand | The random value revealed to the Node after completing an event relay |
| redeemerRand | The random value generated by the Node itself |
| sig | The signature of the sender (signs the hash of the ticket) |

Redeeming a ticket will revert if the Node fails to have a valid `Listing` or if the Node failed to call both `joinDirectory` and `initializeRewardPool` for the epoch the ticket was issued in.

If a ticket is successfully redeemed, the ticket's face value is removed from the sender's deposit and transferred to the `RewardsManager` contract. An internal function call is made to the `RewardsManager` contract to increment the reward balance for the node and for it's delegated stakers for the current epoch.

In the case that a ticket is redeemed though the sender does not have hold sufficient value in the deposit escrow, the sender's penalty deposit is also "burned". Burning in this case refers to transferring those tokens to the "deAd" address (`0x000000000000000000000000000000000000dEaD`).

## Appendix

### Cumulative Reward Factor
