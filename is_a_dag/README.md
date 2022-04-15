## WTF is a Directed Acyclic Graph?

### Notes:
- nodes don't have to validate the transactions that take place on the network
- Gossip about gossip
- vitrual voting
- thousands of transactions per second because no PoW
- transactions that prevail simply require majority support within the network
- equal significance to each node on the network
- miner can issue a transaction and validate transaction at same time
- each vertex represents a transaction
- each transaction is built on top of another
- for a new transaction to be added it must build on top of older ones
- Alice creates a new transation
- for it to be acknowledged, the transaction must reference previous ones
- instead of pruning new growth, growth is woven back into the body of the ledger
- any user can create a transaction, that is put in a "block" and spread through network
- in hashgraph no transactions are discarded

### Question:
- can I lock just part of a dag

#### Sources: 
- https://en.wikipedia.org/wiki/Directed_acyclic_graph
- https://pkg.go.dev/github.com/hashicorp/terraform/dag
- https://medium.com/@kotsbtechcdac/dag-will-overcome-blockchain-problems-dag-vs-blockchain-9ca302651122