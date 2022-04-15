## WTF is Gossip(Distributed Consensus)?

### ELI5
- you give your transaction to random people who give them to random people 
- spread exponentially fast and everyone gets every transaction fast
- no leader
- no taking turns
- no consensus on order, add a shim for who talked to who and when
- you're explaining the details about gossip as you gossip

#### Sources:
- https://ilyasergey.net/CS6213/week-03-bft.html#:~:text=3.2.-,FLP%20Theorem,one%20node%20may%20experience%20failure
- https://hedera.com/learning/what-is-gossip-about-gossip
- https://github.com/bgokden/gossip-to-gossip
- https://www.hashicorp.com/resources/everybody-talks-gossip-serf-memberlist-raft-swim-hashicorp-consul
- https://www.swirlds.com/downloads/SWIRLDS-TR-2016-01.pdf
- https://www.serf.io/docs/internals/gossip.html
- https://docs.avax.network/learn/platform-overview/avalanche-consensus