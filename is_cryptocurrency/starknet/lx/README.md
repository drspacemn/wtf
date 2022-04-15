## WTF is a Fractal Scaling?

### Notes:
- L3s can be realized using validity proofs as long as long as L2 is capable of supporting a verifier smart contract
- When L2 uses validity proofs submitted to L1 we get a recursive structure where L2s applify tx cost reduction
- retain the security of L1
- hyper scalability, better contol by app designer
- L2-L3 interop: on/off ramp flows
  
Building Blocks:
- smart contract tracking the L2 state root on L1(already have)
- verifier smart contract for verifying the validity of state transition proofs
- bridge contracts on L1 managing deposits/withdrawlas tokens
- state tracking and verifier smart contracts on L2

Indexer:
- node operators that extract, transform, and load data into a database by mapping the data into pre-defined schema of tables

#### Sources:
- https://hackmd.io/@kalmanlajko/rkgg9GLG5
- https://medium.com/starkware/fractal-scaling-from-l2-to-l3-7fe238ecfb4f
- https://polynya.medium.com/rollups-data-availability-layers-modular-blockchains-introductory-meta-post-5a1e7a60119d
- https://polynya.medium.com/addressing-common-rollup-misconceptions-eba9d758707e
- https://polynya.medium.com/danksharding-36dc0c8067fe
- https://polynya.medium.com/the-dynamics-around-validity-proof-amortization-519e9ae291c1
- https://polynya.medium.com/why-rollups-data-shards-are-the-only-sustainable-solution-for-high-scalability-c9aabd6fbb48
- https://polynya.medium.com/the-dynamics-around-validity-proof-amortization-519e9ae291c1
- https://polynya.medium.com/anatoly-yakovenko-on-solana-as-an-ethereum-rollup-78329ca69a98
- https://polynya.medium.com/volitions-best-of-all-worlds-cfd313aec9a8