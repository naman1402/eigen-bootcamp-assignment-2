smart contracts
L1 - Governance layer 
AVS Governance: manage operators and governance policies
L1MessageHandler - broadcast messages to L2 via AVS governance
L1AvsTreasury - Manage rewards and handling protocol fee

L2 - Task verification layer to minimize operational cost while maintaining availability, high throughput and low latency
AttestionCenter - Manage verification and historical footprint
OBLS - Implements Multisig operations and handles BLS signature aggregation
L2MessageHandler - Used for communicating with L1 via Attestation center
L2AvsTreasury - Manages rewards and handling protocol fee
InternalTaskHandler - Manages and handles internal task

Operator Roles
Performer - executes off-chain tasks and generate a Proof-of-Task
Attester - Validate the work of Performers and cast a attestations
Aggregator - Collect attestations, aggregate signature and submit final result on-chain
Bootstap Node - Serve as initial points of contact for peers joining the network

Consensus schema:
when a task is triggered, a performer node executes the task and generates a proof-of-task and submit it to the AVS network.
Attester nodes validate the Proof-of-Task submitted by the performer, casting BLS-signed attestations, weight of attestation is proportional to the amount of restaked assets, voting power
Aggregation - colelcts all valid attestation casts by operators and monitors for quorum: Approval or Rejection, once quorum is reched the Aggregator aggregates all the attestations and submits an on-chain txn to the AttestationCenter contract to finalize the task.
Verification and Finality - AttestationCenter smart contract plays a key role in verifying the off-chain execution and maintaining a transparent on-chain record of all activities.

Execution Service - runs performer nodes and is responsible for executing task and generating PoT. this PoT could be result of a calculation, a Zk proof, CID of a JSON on IPFS, etc.

Validation Service - run attester nodes and is responsible for verifying the validity of a Task execution submitted by a performer. Uses validation service API endpoint at /task/validate to verify the proofOfTaskgo mo