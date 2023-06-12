# Sylo Network Protocol Technical Specification

Paul Freeman <paul@sylo.io> </br>
John Carlo San Pedro <john@sylo.io> </br>
Joshua Dawes <josh@sylo.io> </br>

## Table of Contents

- [Introduction](#introduction)
- [Users](#users)
- [Network Parameters](#network-parameters)
- [Smart Contract Specification](#smart-contract-specification)
  - [Data Types](#data-types)
  - [Functions](#functions)
- [Deployment Timeline](#deployment-timeline)
- [Appendix](#appendix)

## Introduction

The Sylo Network Protocol is a suite of smart contracts deployed on Ethereum,
which govern a system of compensating service providers via micro-probabilistic
payments. This protocol will be used to power the [Event Relay
Protocol](overview.md#the-event-relay-protocol) that allows for trustless,
decentralized communication. The purpose of this document is to help understand
the current implementation of the contracts. A more general
[overview](overview.md) of the system is also available. This suite of contracts
is also currently scoped for [phase two](#phase-two) of the Sylo Network incentivization plan.

## Users

The Sylo Network consists of multiple types of users interacting with the smart
contracts.

- **Nodes**: Users that wish to operate and maintain a Sylo Node that will
  support the Event Relay Protocol. Running a Node will allow for compensation
  via redeeming micro-payment tickets. Node operators will be required to have
  `SYLO` tokens staked against their node in order to participate in the
  network, and the amount of work (thus the amount of compensation) that a
  Node receives will be based on their proportion of stake relative to other
  Nodes within the network. The overview covers the [scan weighted
  function](overview.md#scanning-for-a-node) in more detail.
- **Delegated Stakers**: Delegated Stakers (or delegators) are users that wish
  to participate in the Sylo Network and earn `SYLO` tokens without needing to
  run a Sylo Node themselves. These users can supply additional staked `SYLO`
  to an existing Node within the Network in order to increase the Node's
  potential for generating revenue. Delegated Stakers will be rewarded on a
  pro-rata basis.
- **Senders**: Senders are users who hold `SYLO` tokens and wish to utilize
  Sylo Nodes for their decentralized communication service. Senders must
  deposit `SYLO` tokens into both an `escrow` and `penalty` balance held
  within a smart contract. Nodes will be paid via these balances, and senders
  are required to maintain a healthy level of both `escrow` and `penalty` to
  be able to participate in the network. The contracts and event relay service
  allows for the `escrow` to be deposited by a third-party.
- **Receivers**: Receivers do not explicitly interact with the contracts but
  play a critical role in the Event Relay protocol. On receiving an event,
  receivers will reveal the necessary information in order for a Node to
  redeem a winning ticket and be compensated. The overview document goes into
  more detail with regards to the [event relay
  mechanism](overview.md#asynchronous-event-relay)
- **Sylo**: The Sylo team will be deploying the contracts to Ethereum and will
  have "ownership" of the contracts. Ownership allows certain privileged
  functions to be called on the contracts. These functions range from manually
  adjusting network parameters to making the call to initialize the next
  Epoch. These responsibilities will be passed over to a DAO.

A [sequence diagram showcasing the interactions](contracts_sequence_diagram.png)
between the users and the contracts is also available.

## Network Parameters

The following network parameters will be manually set and adjusted by the Sylo
Team over the course of phase two. Note: `SYLO` token parameters are all
specified in [`SOLO`](#sylo/solo) units instead.

#### **epochDuration**

The minimum duration in blocks the next epoch will last for. Attempting to
initialize the next epoch if the current epoch's duration has not yet reached
this value will result in failure.

#### **defaultPayoutPercentage**

The payout percentage refers to the percentage of a ticket's face value that
will be divvied out to a Node's stakers. The remaining value is then
given to the Node as a fee for providing Event Relay service.

Example:

If this value was set to `40%`, and a ticket's value was `1000 SOLO`. Then on
redeeming a ticket, `400 SOLO` would be set aside for those who have stake in the Node, and shared out in proportion to the amount staked. The
remaining (`600 SOLO`) is given directly to the Node.

The `defaultPayoutPercentage` parameter is the default value for this. **Note**:
For phase two, the default value is used for all Nodes and supersedes the
`payoutPercentage` value set in a Node's registry.

Changes to this value will only take effect in the next epoch.

#### **faceValue**

The value in `SOLO` of a winning ticket.

Changes to this value will only take effect in the next epoch.

#### **baseLiveWinProb**

The probability of a ticket winning immediately after the ticket is issued.

Changes to this value will only take effect in the next epoch.

#### **expiredWinProb**

The probability of a ticket winning after the ticket's entire duration has
elapsed. (Payouts for expired tickets are not currently implemented).

Changes to this value will only take effect in the next epoch.

#### **ticketDuration**

The duration in blocks a ticket is alive for.

Changes to this value will only take effect in the next epoch.

#### **decayRate**

The rate at which a ticket's winning probability will decay over its `ticketDuration`,
expressed as a percentage.

Example:

A decay rate of `80%` and a base live winning probability of `10%` indicates
that at the block immediately before a ticket has expired, the ticket's winning
probability will have decayed to `2%`.

#### **unlockDuration**

The duration in blocks that must elapse before either deposits or stake can be
withdrawn (once the unlocking phase has begun).

#### **minimumStakeProportion**

The minimum amount of stake a Node must own for itself, expressed as a
percentage of the Node's overall stake. This requirement must always
be must met whenever the Node unlocks stake, or if other delegators attempt to
add more stake to the Node. Failing to meet this requirement will prevent the
Node from participating in the next epoch.

Example:

A minimum stake proportion of `20%` indicates that the Node must own 20% of its
total stake. Thus if the stake total was `1000 SOLO`, then the Node must
own at least `200 SOLO` of this stake to participate in the network.

## Smart Contract Specification

The Sylo Network Protocol contracts are written in Solidity and will initially
be deployed to the Ethereum mainnet. The current system includes:

- `SyloToken`: ERC20 contract for the Sylo Token which has already been
  deployed.
- `SyloTicketing`: Contract that manages user deposits for payments, and
  implements the `redeem` function for redeeming winning tickets.
- `StakingManager`: Tracks the amount of stake and the delegated stakers for
  each Node
- `Directory`: Creates and manages a `Directory` structure every epoch based
  on the stake held by each stakee. The `Directory` is used as the backend for
  the stake-weighted scan function.
- `RewardsManager`: Tracks rewards for each epoch when winning tickets are
  redeemed.
- `EpochsManager`: Manages initializing of each epoch and stores the Network
  parameters for every epoch.
- `Registries`: Stores a `Registry` struct for every Node

### Data Types

---

#### **Epoch**

Network parameters for the current epoch are saved into this structure by the
`EpochsManager` contract when every new epoch is initialized.

| Field                   | Description                                                                                                 |
| ----------------------- | ----------------------------------------------------------------------------------------------------------- |
| startBlock              | The block the epoch started                                                                                 |
| duration                | The duration in blocks the epoch will last for                                                              |
| endBlock                | The block the epoch ended. Initially set to 0 but will be updated when the next epoch is initialized        |
| defaultPayoutPercentage | [See defaultPayoutPercentage](#defaultPayoutPercentage)                                                     |
| faceValue               | [See faceValue](#faceValue)                                                                                 |
| baseLiveWinProb         | [See baseLiveWinProb](#baseLiveWinProb)                                                                     |
| expiredWinProb          | [See expiredWinProb](#expiredWinProb)                                                                       |
| ticketDuration          | [See ticketDuration](#ticketDuration)                                                                       |
| decayRate               | [See defaultPayoutPercentage](#decayRate)                                                                   |

#### **Ticket**

Tickets are created by senders at the client level and are given to Nodes as
compensation for providing event relay.

| Field           | Description                                              |
| --------------- | -------------------------------------------------------- |
| epochId         | The id of the epoch the ticket was generated in          |
| sender          | Address of the sender and its delegated addresss         |
| receiver        | Address of the receiver and its delegated addresss       |
| redeemer        | Address of the redeemer (Usually the node)               |
| generationBlock | The approximate block number the ticket was generated at |
| redeemerCommit  | Hash of the secret random number of the redeemer         |

#### **Stake**

The _Stake_ datatype tracks a Node's current total managed stake, and each
individual delegated stake entry.

| Field             | Description                                                                        |
| ----------------- | ---------------------------------------------------------------------------------- |
| stakeEntries      | A mapping between each delegated staker and it's `StakeEntry` for the given stakee |
| totalManagedStake | The sum of all delegated stake amounts for the stakee                              |

#### **StakeEntry**

A datatype that tracks a delegated staker's stake for a particular stakee.

| Field     | Description                                               |
| --------- | --------------------------------------------------------- |
| amount    | The amount of delegated stake in `SOLO`                   |
| updatedAt | The block number this stake entry was updated at          |
| epochId   | The epoch id of the epoch this stake entry was updated in |

#### **Directory**

A snapshot of all staking entries at the time it was created.

| Field      | Description                                                             |
| ---------- | ----------------------------------------------------------------------- |
| entries    | An array of `DirectoryEntry` that is iterated over during [scan](#scan) |
| stakes     | A mapping of each stakee to their total stake                           |
| totalStake | The sum of all stakes                                                   |

#### **DirectoryEntry**

This datatype helps to make the `scan` implementation more efficient. A
`DirectoryEntry` value is created as Nodes join the next epoch's directory. The
entry includes a `boundary` value which is a sum of the current directory's
total stake, and the current Node's total stake. This entry is then pushed to
the end of the `entries` array for the given directory.

| Field    | Description                       |
| -------- | --------------------------------- |
| stakee   | The address of the stakee         |
| boundary | The boundary value for this entry |

#### **RewardPool**

Each Node must initialize a reward pool for every epoch they wish to participate
in. The reward pool will help track the portion of rewards from redeeming
tickets that will be distributed to a Node's delegated stakers. The reward pool
also tracks the [cumulative reward
factor](#reward-calculation-and-cumulative-reward-factor) (CRF) to make
calculating the distribution more efficient.

| Field                         | Description                                                                 |
| ----------------------------- | --------------------------------------------------------------------------- |
| stakersRewardTotal            | The balance of the reward pool                                              |
| initializedAt                 | The block number the reward pool was initialized                            |
| totalActiveStakes             | Tracks the total active stake for this reward pool                          |
| initialCumulativeRewardFactor | The CRF at the time the reward pool was initialized                         |
| cumulativeRewardFactor        | The ongoing CRF of the reward pool which is updated as tickets are redeemed |

#### **Registry**

Every Node must also have a `Registry` entry. The entry holds various network
parameters which are configured by Nodes themselves.

| Field            | Description                                                                                                                                                                                                          |
| ---------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| multiAddr        | The libp2p multi address of the Node. This is needed for clients to connect to their Node. Nodes should take care to ensure this value is correct and up to date                                                     |
| payoutPercentage | Percentage of a redeemed tickets value that will be paid out to the Node's delegated stakers. **This value is currently unused and is superseded by the _defaultPayoutPercentage_ network parameter for phase two**. |

### Functions

This section will detail the various contract calls that will be made over the
lifetime of the Sylo Network.

---

#### Registries

#### _register_

Nodes are required to set their `Registry` entry to be able to stake and redeem
tickets.

| Param     | Description                      |
| --------- | -------------------------------- |
| multiAddr | Sets the multi addr for the Node |

---

#### StakingManager

#### _addStake_

Called by both Nodes and delegators. This will transfer `SOLO` from the
`msg.sender` to the `StakingManager` contract, and create or update a stake
entry for the specified stakee. Additionally it will also automatically claim
any outstanding stake rewards. This function will fail if the additional added
stake will cause the Node to own less than require
[minimumStakeProportion](#minimumStakeProportion).

| Param  | Description                          |
| ------ | ------------------------------------ |
| amount | The amount of stake to add in `SOLO` |
| stakee | The address of the stakee            |

#### _unlockStake_

Allows Node and delegators to set their stake for unlocking, which eventually
will allow the stake to be withdrawn once the unlocking phase has ended. This
removes the stake for consideration in the next epoch. If any stake was already
in the unlocking phase, the amount of unlocking stake will instead be increased
and the unlock duration will be reset.

| Param  | Description                             |
| ------ | --------------------------------------- |
| amount | The amount of stake to unlock in `SOLO` |
| stakee | The address of the stakee               |

#### _cancelUnlocking_

Cancels stake that is in the unlocking phase and adds it back to the total
managed stake for that stakee. The re-added stake can be utilized in the next
epoch. Cancelling more stake than is available results in the maximum amount
of unlocking stake being cancelled.

| Param  | Description                                       |
| ------ | ------------------------------------------------- |
| amount | The amount of unlocking stake to cancel in `SOLO` |
| stakee | The address of the stakee                         |

##### _withdrawStake_

Returns stake that has finished unlocking back to the `msg.sender` account. This
function will fail if the stake has not finished unlocking.

| Param  | Description               |
| ------ | ------------------------- |
| stakee | The address of the stakee |

---

#### Directory

#### _joinNextDirectory_

Called by Nodes as a prerequisite to participating in the Sylo Network for the
next epoch. This function allows the stake delegated to a Node be used in the
`scan` function. It will create and append a `DirectoryEntry` based on the sum
of the total managed stake the Node has, plus any unclaimed staking rewards.

There are no explicit parameters for this function, though it only allows a Node to
call this function once per epoch. It is in the Node's best interest to call
this function near the end of the current epoch, in order to maximize the amount
of unclaimed reward that can be included as stake in the next Epoch's directory entry.

#### _scan_

Called by users of the Event Relay service to find the node associated with a
given `point`. The `point` value is any 16 byte value, which is likely to be the
hash of some user identifier. Hashing a user ID will create a psuedo-random
value, which the `scan` function then maps to a value between 0 and the total
stake in the current epoch's directory. This is then used in a binary search
with the directory's entries, eventually returning the address of a Node. Node's
with larger proportions of stake are more likely to be returned by the `scan`
function.

| Param | Description           |
| ----- | --------------------- |
| point | A psuedo-random value |

**Returns**: Address of a Node

---

#### RewardsManager

#### _initializeRewardPool_

Called by Nodes as a prerequisite to participating in the Sylo Network for the
next epoch. This function initializes and stores a new `RewardPool` entry for
the next epoch for this Node. It will calculate the `totalActiveStake` for this
reward pool based on the sum of the total managed stake the Node has. The new
reward pool will also read the `cumulativeRewardFactor` from the previous pool
and begin tracking a new factor for the next epoch.

There are no explicit parameters for this function, though it only allows a Node to
call this function once per epoch. It is in the Node's best interest to call
this function near the end of the current epoch, in order to maximize the amount
of unclaimed reward that can be used for the directory entry.

#### _claimStakeReward_

This function is called by Nodes and delegators when they wish to claim rewards
that their stake has gained for them. This will utilize the current reward
pool's CRF and the CRF at the time their stake became active to calculate the
value of their reward. See [Cumulative Reward Factor] for details of the
calculation though generally having a higher proportion of stake compared to
other delegators will lead to a larger reward claim. Calling this function will
prevent the user from being eligible to claim any further rewards until the next
epoch begins

| Param  | Description                                            |
| ------ | ------------------------------------------------------ |
| stakee | Address of the stakee the user wishes to claim against |

A public function `calculateStakerClaim` is exposed by the `RewardsManager`
contract which allows users to understand the amount in `SOLO` gained if they
were to call `claimStakingReward`. As claiming rewards will also remove the users unclaimed
reward from being used in the total active stake for the next epoch, users may
wish to wait until the reward value is high enough to offset gas costs.

#### _claimNodeReward_

This is called by Node operators when they wish to withdraw rewards gained from
operating a Node. The current value of this reward is a public field of the
`RewardsManager` contract.

Node operators may wish to wait until the reward value is high enough to
minimize the loss in earnings from gas costs.

---

#### EpochsManager

#### _initializeEpoch_

Sylo will take responsibility of calling this function every epoch to initialize
the next epoch. Invoking this function will read the current set of network
parameters and store it a new `Epoch` value. This function will fail if the
current epoch has yet to end.

---

#### Ticketing

#### _depositEscrow_

This function is called by users that wish to utilize the Event Relay service
provided by Node. This function transfers a specified amount of `SOLO` to be
held in escrow by the Ticketing contract. When winning tickets are redeemed, the
face value of the ticket will be paid out from the escrow.

| Param   | Description                                                                                                       |
| ------- | ----------------------------------------------------------------------------------------------------------------- |
| amount  | The amount in `SOLO` to deposit                                                                                   |
| account | The account the deposit will belong to. **Note**: The tokens are still transferred from the `msg.sender` account. |

##### _depositPenalty_

This function should be called in conjunction with `depositEscrow` to also hold
a `penalty` amount in escrow. When winning tickets are redeemed, if the face
value of a ticket is greater than the sender's escrow, then the penalty will be
burned instead. This is to prevent an economic attack on the probabilistic micropayment mechanism called front-running. [Further detail
with regards to the economics of the Sylo Network can be found in the
overview](overview.md).

##### _unlockDeposits_

Moves both existing escrow and penalty values to an unlocking phase, which
eventually allows withdrawal once the unlocking phase has completed. This
function will fail if the user has already begun unlocking their deposits.

##### _lockDeposits_

This function essentially cancels the unlocking phase and allows the token to be
used again as deposits.

##### _withdraw_

Once the unlocking phase has completed, this function can be called to transfer
the tokens held in escrow back to the `msg.sender`.

##### _redeem_

`Redeem` should be called by the Node after completing an event relay and
learning of the ticket sender's secret random value. The Node should only call
this if it understands the ticket will win. The `Ticketing` contract exposes
both [calculateWinningProbability](#calculateWinningProbability) and [isWinningTicket](#isWinningTicket) functions that can be
used to determine if a ticket is winning, though the Node can also perform the
calculation locally.

| Param        | Description                                                           |
| ------------ | --------------------------------------------------------------------- |
| ticket       | The [Ticket](#Ticket) issued by the sender                            |
| redeemerRand | The random value generated by the Node itself                         |
| senderSig    | The signature of the sender (signs the hash of the ticket)            |
| receiverSig  | The signature of the receiver (signs the hash of the ticket)          |

Redeeming a ticket will revert if the Node fails to have a valid `Registry` or if
the Node failed to call both `joinDirectory` and `initializeRewardPool` for the
epoch the ticket was issued in.

If a ticket is successfully redeemed, the ticket's face value is removed from
the sender's deposit and transferred to the `RewardsManager` contract. An
internal function call is made to the `RewardsManager` contract to increment the
reward balance for the node and for it's delegated stakers for the current
epoch.

In the case that a ticket is redeemed though the sender does not have hold
sufficient value in the deposit escrow, the sender's penalty deposit is also
"burned". Burning in this case refers to transferring those tokens to the "deAd"
address (`0x000000000000000000000000000000000000dEaD`).

#### _calculateWinningProbability_

This function calculates the probability of a ticket winning based on the
ticket's [baseLiveWinProb](#baseLiveWinProb) and the number of blocks that has
elapsed since the ticket was generated. The ticket's parameters are retrieved
from the Epoch the ticket is associated with. The calculation is as follows:

<img style="background-color:white;padding:3px" src="https://render.githubusercontent.com/render/math?math=p=baseLiveWinProb - baseLiveProb * decayRate * blocksElapsed / ticketDuration">

| Param  | Description                                                           |
| ------ | --------------------------------------------------------------------- |
| Ticket | The ticket to calculate probability for                               |
| Epoch  | The epoch associated with the epoch that holds the tickets parameters |

**Returns**: A value between 0 and 2^128-1 representing the probability

#### _isWinningTicket_

Given the probability of the ticket winning, and the signature and the redeemer
random number of the ticket, this function checks if the ticket is actually a
winner. This is done by checking if the hash of the ticket (as a numerical
value) is less than the specified probability.

| Param        | Description                                               |
| ------------ | --------------------------------------------------------- |
| senderSig    | The signature of the ticket signed by the sender          |
| receiverSig  | The signature of the ticket signed by the receiver        |
| redeemerRand | The random number generated by the redeemer of the ticket |
| winProb      | The winning probability of the ticket                     |

This function is used in conjunction with
[calculateWinningProbability](#calculateWinningProbability) to determine the
`winProb` parameter.

## Deployment Timeline

The Sylo Network will be updated in [three major phases](https://sylo.io/newsroom/article/sylo-network-incentivisation-release-plan). The first phase involves
deploying the Sylo Token ERC20 contract, and increasing liquidity for
the token through various means. This phase has completed. The current set of
contracts is scoped for **phase two**.

### Phase Two

The second phase includes a deployment of the Sylo Network contracts onto the
Ethereum mainnet, as well as a mechanism to incentivize Sylo Node operators in
order to bootstrap the network. This mechanism involves the Sylo Team running
their own Nodes that periodically generate artificial work for the network.
This artificial work will follow the exact same process specified in the
Event Relay Protocol, and from a Node Operator's perspective will look exactly
the same as "real" work.

This phase allows the Sylo Team to have more time to integrate the Event
Relay Protocol into real applications, such as the Sylo Wallet. It also
allows us to discover any flaws or learnings from the current system. Phase
two does not require all network/economic mechanisms to be present in the system,
and as such has not been fully realized in the contracts yet. This includes:

- **Payouts for expired tickets**. Artificial work will be generated from
  Nodes that the Sylo Team operates and should always be online. There should
  not be a case where a winning ticket is redeemed much later than the time it
  was generated.
- **Slashing/Stake Distribution**. This is a complex process that will likely
  benefit from the learnings gained in phase two. Additionally as the work is
  artificial for phase two, there is no gain from having this system in place
  yet.

### Phase Three

Along with implementing the missing functionality stated above, phase three
will also likely involve a transition of ownership/management of the contracts
to a DAO. The Sylo Network will be completely autonomous and decentralized
at this point.

## Appendix

### SYLO/SOLO

A single `SYLO` token refers `10**18 SOLO`. This mirrors the `ETH/WEI`
representation. All function parameters and data types that work with
token values are represented in `SOLO`.

### Reward Calculation and Cumulative Reward Factor

The `Cumulative Reward Factor` is a variable that significantly improves the gas
cost efficiency of calculating staking reward distributions. A stakers reward
share is determined on a pro-rata basis/

| Epoch         | 0   | 1   | 2   | 3   |
| ------------- | --- | --- | --- | --- |
| Reward Gained | 0   | 10  | 8   | 12  |
| Total Stake   | 20  | 20  | 20  | 25  |
| Alice Stake   | 5   | 5   | 5   | 5   |
| Bob Stake     | 15  | 15  | 15  | 20  |
| Alice Reward  | 0   | 2.5 | 1.6 | 2.4 |
| Bob Reward    | 0   | 7.5 | 6   | 9.6 |

- **Reward Gained** refers to the total amount of rewards gained in that epoch
  from redeeming tickets that will be allocated to a Node's stakers
- **Total Stake** will be the total amount of `SOLO` staked towards the Node in
  the **next** epoch. This value is the sum of each participant's stake value.
- **Alice Stake** and **Bob Stake** are Alice's and Bob's respective stakes.
- **Alice Reward** and **Bob Reward** are Alice's and Bob's respective rewards
  accumulated at the end of the epoch. Their reward is determined by how much
  stake they have.

It is simple enough to calculate each staker's share of the reward for an epoch
manually.

```math
aliceReward_1 = 5 * 10/20 = 2.5
```

Alice’s reward at the end of an epoch can be determined by multiplying alice’s
proportion of stake against the reward gained
at the specified epoch. Then adding that value to their reward from the previous epoch.

For epochs 1 and 2, we can perform:

```math
aliceReward_2 = 5 * 10/20 + 5 * 8/25 = 4.1
```

Calculating the reward share with this approach is problematic as more state
needs to be read for every epoch the calculation is being made for. This
can lead to excessively high gas costs that scale with the number of epochs that
needs to be accounted for. This can easily cause the withdraw process to
become too expensive in terms of gas costs, as essentially the the number of
storage reads that need to be performed will scale linearly with the
number of epochs.

To solve this issue, a cumulative reward factor variable is introduced. If we
examine the above calculations we can notice that Alice's initial stake of `5`
is a constant and the calculation can be easily simplified.

```math
aliceReward_1/5 = 10/20 = 0.5
aliceReward_2/5 = 10/20 + 8/20 = 0.9
```

The CRF at the end of epoch 1 is `0.5` and at the end of epoch 2 is `0.9`.
As long as Alice's stake remains constant for the set of epochs we are calculating for,
we can use the CRF values to calculate Alice's reward without needing to iterate through
each epoch.
Alice's reward is calculated by multiplying the CRF value by Alice's stake.

```math
aliceReward_1 = 0.5 * 5 = 2.5
aliceReward_2 = 0.9 * 5 = 4.5
```

The same values can be used to also calculate Bob's share of rewards at the end of each epoch

```math
bobReward_1 = 0.5 * 15 = 7.5
bobReward_2 = 0.9 * 15 = 13.5
```

The value of calculating `0.9 * 15` is the same as if we were to manually calculate Bob's
reward for both epoch 1 and 2, and then sum them together.

The above approach only works when the participant's stake is constant
throughout the calculation (or has not yet made a claim). When the stake changes or a claim
is made, we need to also track the CRF value when their stake was changed.

For example on Epoch 3, Bob adds 5 SYLOs to stake, bringing his stake to 20, and the total
stake for the node to 25.

The CRF for the end of epoch 3 is:

```math
CRF_3 = 0.9 + 12 / 25 = 1.38
```

The CRF for epoch 2 was `0.9`. To calculate the reward Bob gained at epoch 3, we need to
use the difference in CRF values from when the calculation starts, to when the calculation ends.

```math
bobRewardAtEpoch3 = 20 * (CRF_3 - CRF_2) = 20 * 0.48 = 9.6
```
