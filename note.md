Dynamic NFTs representing off-chain achievements like Github contributions or PlayStation platinum trophies.

1. nft - ERC-712 with upgradable tokenURI, metadata stored on IPFS or returned via on-chain logic
2. avs operators run script that query external API (Github, PSN, etc), verify the condition, output a signed metadata change, signature is submitted onchain to confirm the change.
