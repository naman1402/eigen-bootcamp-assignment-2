dark pool - secret trading room, hides your trading activity while being decentralized.

use of zk - proves wallet state, validates your order, confirm ownership of wallet
use of MPC - private order matching where multiple nodes collaborate to find matches without seeing individual orders, distributed computations, collective decision making.

components:
1. frontend: UI to place order, client-side encryption of order details, creates zk proof of wallet state and order validity
2. AVS: performer nodes (validation of orders and proof), attester network (matches order using MPC), Consensus Mechanism (agreement on order matching and execution)
3. Privacy layer - zk, mcp for private order matching, encrypted order book
4. execution layer - uniswap v4, final trade confirmation and fund transfer

┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Trader A      │    │   Trader B      │    │   Trader C      │
│                 │    │                 │    │                 │
│ Order: Buy ETH  │    │ Order: Sell ETH │    │ Order: Buy BTC  │
│ + ZK Proof      │    │ + ZK Proof      │    │ + ZK Proof      │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          └──────────────────────┼──────────────────────┘
                                 │
                    ┌─────────────▼──────────────┐
                    │      Performer Nodes       │
                    │                            │
                    │  • Validate ZK Proofs     │
                    │  • Check Order Format     │
                    │  • Verify Wallet State    │
                    └─────────────┬──────────────┘
                                 │
                    ┌─────────────▼──────────────┐
                    │     Attester Network       │
                    │                            │
                    │  Node 1  │  Node 2  │ ... │
                    │    ▲     │    ▲     │     │
                    │    │     │    │     │     │
                    │    └─────┼────┘     │     │
                    │        MPC Protocol │     │
                    │     (Private Match) │     │
                    └─────────────┬──────────────┘
                                 │
                    ┌─────────────▼──────────────┐
                    │      Uniswap V4 Hook      │
                    │                            │
                    │  • Execute Matched Trades  │
                    │  • Update Pool State       │
                    │  • Transfer Funds          │
                    └────────────────────────────┘