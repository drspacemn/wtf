## WTF is Bitcoin?

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