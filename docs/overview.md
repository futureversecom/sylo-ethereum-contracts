# Sylo Network Phase 2 Smart Contracts

**Protocol and Economic Incentives For Decentralized Communications Infrastructure**

Paul Freeman            <paul@sylo.io> </br>
John Carlo San Pedro    <john@sylo.io> </br>
Joshua Dawes            <josh@sylo.io> </br>


## Table of Contents ###########################################

* [Introduction and Background](#introduction-and-background)
* [Probabilistic Micropayments](#probabilistic-micropayments)
	* [Micropayment Tickets](#micropayment-tickets)
* [The Event Relay Protocol](#the-event-relay-protocol)
	* [Nodes](#nodes)
* [Staking](#staking)
	* [Delegated Staking](#delegated-staking)
	* [Scanning for a node](#scanning-for-a-node)
* [Epochs](#epochs)
* [Rewards](#rewards)
* [Ticket decay with time](#ticket-decay-with-time)



## Introduction and Background ###########################################

Sylo Nodes are an application that anyone can run, to help provide network services to Sylo users in a truly private, fully decentralised way.

Sylo Nodes provide Incentivised Event Relay to applications and users of the Sylo Network.

Network traffic is allocated to Nodes based on the amount of Sylo Tokens staked. This information is saved on-chain in a Stake Directory.

Sylo Tickets are the payment mechanism on the network; they are used to pay for very small units of work, off-chain.

An Epoch is the main unit of time in the Sylo Network, measured as a number of on-chain blocks.


## Probabilistic Micropayments ###########################################

Sylo Tickets are probabilistic micropayments used for rewarding Sylo Nodes for their work. Because Sylo Nodes can be run by anybody, they need a financial incentive to perform relay - the Sylo Network cannot rely on altruism and still provide a reliable, high quality service.

A single relay is so inexpensive that the blockchain transaction costs associated with paying per message would be unreasonably high. Because of this, we need a way to exchange value for every relay performed, without relying on an on-chain transaction each time.



### Micropayment Tickets

Instead of transferring currency every transaction, Sylo Tickets are a probabilistic payment. A ticket is sent along with every transaction, but not every ticket can be redeemed for currency.

Tickets only have a small percentage chance of winning, which means that most Tickets are not winning. As a result, the vast majority of relay transactions can occur without an on-chain transfer of funds. Neither the sender nor the receiver of a ticket knows whether a ticket is a winning ticket until the relay it is paying for is completed.

If the Ticket wins, it is worth it's full face value, and can be submitted on-chain to claim that full payout for the sender's funds in escrow. These funds are locked in escrow for a substantial period of time, to ensure that Tickets redeemed in the future will still have funds available to redeem winnings from.

In addition, the owner of the escrow must post penalty escrow, which is burned if the payment escrow is ever empty when a ticket is redeemed. This penalty escrow prevents the owner of the escrow from emptying their own escrow balance into another wallet, as a means of avoiding paying nodes for their work - an attack known as front-running.

The value of a Ticket is it’s expected value: the Ticket's “face value” that is paid out if it wins, multiplied by its probability of winning. By choosing values of these two parameters, these probabilistic micropayment Tickets can be used to pay for arbitrarily small units of work in a gas-efficient way.

## The Event Relay Protocol ###########################################


### Asynchronous Event Relay

Peer Alice wishes to send a packet to peer Bob via Bob’s node.

Both Alice and Bob use the Sylo network intermittently - in the worst case, they may never be online at the same time as each other. This means that Alice is unable to check in with Bob later, to ensure that her relays were delivered.  Alice needs a protocol that is "fire and forget" - once she has left a relay request with Bob's node, Alice needs to be sure that the node will do it's best to deliver her message to Bob, with no further input from her.

This means that Alice needs a trustless method of setting payment aside for Bob's node, to be claimed once the relay is delivered. This is accomplished by signing a ticket, which will pay out from money held in escrow in a smart contract when the ticket is redeemed.

Alice also needs a way to release payment to Bob's node only once Bob has received the relay packet. This is the difficult part, because Alice may never have the opportunity to learn anything more about the outcome of her relay request.


In reality, only two actors know when the relay packet has been delivered - Bob, and Bob's node.

Bob's node cannot be trusted to be truthful about delivery, because it has a financial incentive to lie, and claim payment for delivery without doing the work. This is true in both the one-off case, and in the iterative game, where the node’s optimal strategy is a mixed strategy that includes some cheating.

In general, Bob also cannot be trusted to unlock Alice’s payment for Bob's node. Because Alice is sending to Bob, Bob and Alice may have some relationship with one another, and so Bob is assumed to have some incentive to refuse to unlock Alice's payment.

However, Bob only has one node - all of Bob’s relay traffic, from a variety of peers, comes through that node. This gives Bob's node a mechanism to punish Bob with, if Bob withholds payment - blacklisting.

By threatening to withhold Bob's future traffic as punishment for bad behavior, Bob can be trusted to unlock Alice’s payment - the small incentive Bob has to help any individual relay sender is outweighed by Bob's desire to remain in good standing with his node.

This allows Bob to act as the service completion oracle. When Alice leaves a relay request with Bob's node, it contains a secret from Alice that is encrypted for Bob. Once Bob receives the relay, he decrypts the secret and passes it back to the node, allowing the node to claim payment.


Blacklisting also prevents abuse in the case where Alice and Bob are the same financial entity, with shared cost incentives. The goal in this case is to limit the amount of free relay that Alice and Bob are able to extract from the network.
Bob's node is able to see which relays do not result in valid payment, and adjust the reputation of both Alice and Bob accordingly. After a small number of failures to deliver payment, Bob's node can blacklist Bob, and decrease Alice’s reputation as a relay sender, eventually resulting in the node ignoring relay requests from Alice altogether.

In the worst case attack, the attacker spins up many alternative receiving peers (for free) that scan to a variety of different nodes, sending relay to each of these peers in turn. This continues until Alice’s reputation has been “spent” with all nodes on the network. With 100 nodes on the network, and 50 relays worth of non-payment evidence required for the node to blacklist A (an allowance for legitimate connection issues), this results in 5000 free relays before all nodes will turn away Alice’s requests - about $0.00005 worth of relay, based on monthly node operating costs and throughput.

After this, Alice’s escrow must be moved to another wallet to begin the process again, which takes time for the escrow to unlock, and has an associated gas cost. This gas cost puts an effective price on the “free” relay obtainable this way, and makes the strategy of repeatedly spinning up new sending peers worse than just paying for relay in the first place.


### Nodes

Sylo Nodes are an application that anyone can run on their own server, to help provide network services to Sylo users in a truly private, fully decentralised way. Sylo Nodes currently provide incentivised event relay and will provide additional services in future.

Sylo Nodes are financially incentivized to provide and maintain good service. For good service they receive payouts via Sylo Tickets at a service price, which is set each epoch.

Sylo Nodes require Stake before they are eligible to receive work to do.

## Staking ###########################################
In order to be allocated work on the Sylo Network, Sylo Nodes must have SYLO Tokens staked against them on-chain. These tokens are still owned by the person who staked them, but are locked in a smart contract for a period of time.

Sylo Node runners provide Node Staking. Sylo users can stake toward Nodes via Delegated Staking.

Using the total amount staked to Nodes across the Network, the Stake Directory is updated each epoch. The Stake Directory is used by the stake-weighted scan function for users to know which Node is assigned to which user.

Staking is important because it ensures that the owners of nodes have a financial stake in the overall network. Staking ensures that node owners’ incentives align with the network as a whole - taking actions that add value to the network adds value to their stake, and taking actions that harm the network reduces the value of their stake.

Stakers can initiate the withdrawal of their stake at any time, but the stake is locked in for a period of time before it is withdrawn, to ensure that speculators cannot take short-term control of the network’s services in order to manipulate them for profit.

Staking against a node does two things:
- It increases the amount of traffic that the node receives from the network, increasing that node’s income.
- It entitles the owner of the stake to a portion of the payout from each winning ticket that the node redeems.


### Delegated staking

A node can be staked by the Node’s owner, or by other holders of the SYLO token - the latter is known as delegated staking.

- A percentage of every winning ticket is paid to the node owner first - the remaining payment is then shared among all stakers, delegated or not,  in proportion to the amount they have staked.

### Scanning for a node

When a Sylo Node has SYLO Token staked against it, this is recorded in the stake directory.

The stake directory is an on-chain data structure that holds information about which Sylo Nodes are staked to provide services in the Sylo Network, and how much SYLO Token is staked against each node.

Any peer with access to the blockchain can see the full list of staked Sylo Nodes, and use the stake directory to determine which node they are assigned to for the current epoch, by a process known as scanning.


When a peer wants to identify their Sylo Node, they query the blockchain using a “scan” function. This function takes the peer ID as input, and pseudo-randomly assigns them a Sylo Node.

The scan function does this using the following steps:

- Compute the hash of the peer ID, concatenated with the current epoch number and the peer’s on-chain channel number. This hash can be mapped to a pseudo-random number between 0 and 1. It is unique to each peer, and changes each epoch.
- Multiply the total amount of stake in the stake tree by this random number, to produce a number between 0 and total_staked.
- Binary search the stake directory to find which node is associated with the number produced above, returning that node.


This scan function allows anyone who knows your peer ID to efficiently identify your node for a given Epoch, and therefore identify which Sylo Node will accept a relay message for you.

It also ensures that Sylo Nodes are able to tell when a relay message is unlikely to be collected.  If the recipient scans to a different Sylo Node, then it is unlikely that the recipient will come to this Node to collect it. This incentivizes nodes to only provide services to peers that “scan” to them, and incentivizes peers to only use their assigned node.

SYLO holders can alter their staking arrangements at any time - however, to avoid peer’s Sylo Nodes changing mid-epoch, these changes only take effect on the stake directory when the epoch changes. This ensures that a peer’s Sylo Node is consistent for the entire epoch.

Because the stake directory is updated each epoch, and because the scan function takes the epoch number as input to it’s hash, each peer’s Sylo Node is randomized each epoch. This provides additional security against traffic analysis, and ensures that nodes receive work in proportion to their stake over the long run, regardless of epoch-by-epoch traffic variation.

## Epochs ###########################################

An epoch is the main unit of time in the Sylo Network, measured as a number of on-chain blocks.
At the start of each epoch, the stake directory is updated based on all staking changes made during in the previous epoch, and each peer is randomly reassigned to a Sylo Node.

Epochs have several benefits:

They reduce the gas overhead of running the on-chain components of the network. They also provide a predictable time for changes to come into effect, simplifying the process of peers monitoring the network for changes - e.g.  has my Node changed?


## Rewards ###########################################

Rewards gained from redeeming tickets are held in escrow, and will continue to accumulate until either the Node or a delegated staker withdraws their rewards.
On redeeming a ticket, a portion of the reward is allocated to the Node itself as direct payment for running the Event Relay service. The remaining portion is then split amongst the Node's stakers on a pro-rata basis. Unclaimed staking rewards are also automatically reconsidered as part of the Node's total stake, bringing the Node additional network traffic until that stake is claimed.

Further detail of the staking rewards are calculated over multiple epochs can be found in the [technical specification](spec.md#reward-calculation-and-cumulative-reward-factor).

## Ticket decay with time ###########################################

To incentivize relays to be delivered as soon as possible, the mechanism that pays out winning Sylo Tickets is modified, so that the probability of a Ticket winning decreases as the time since the relay request was made increases.

The current blockchain block number is included in each Sylo Ticket when the ticket is issued, and signed by the ticket's sender. This block number defines the start of the ticket's valid duration, and is used to measure time. This trustless measure of elapsed time is then used to modify the ticket's probability of winning, decreasing the ticket's expected value as time goes on.

This creates a direct incentive for Sylo Nodes to perform relay as soon as possible, to maximise their income from performing the service. It similarly incentivises them to invest in throughput and uptime improvements, so that they can claim tickets as soon as possible, and earn more from the work that they do.
