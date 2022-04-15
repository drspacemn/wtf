## WTF is Bitcoin?

### Notes:
- Peer-to-peer distributed timestamp server to generate computational proof of the chronological order of transactions
- Transactions must be publicly announced and need a system for participants to agree on a single history
- Timestamp server takes hash of block of items and publishes the hash(each timestamp includes the previous timestamp)
- POW: scanning for a value that when hashed, the hash begins with a number of zero bits
  - increment a nonce in the block until

### Run Examples:
PoW example
change the 'targetBits' constant to see how the mining difficulty affects block creation
```
cd mini/go
go run main.go
```

Pull an example bitcoin block to verify
```
cd verify/
curl https://blockchain.info/rawblock/0000000000000000000836929e872bb5a678546b0a19900b974c206c338f0947 > rawBTCBlock.json
cd go/
go run *.go
```

Run verify benchmark
```
cd verify/go/
go test -bench=. -count 5
```

#### Sources:
- https://bitcoin.org/bitcoin.pdf
- https://github.com/Jeiwan/blockchain_go/tree/part_2