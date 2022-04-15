## WTF is StarkNet?

### State:
- mapping between addresses(felts) and contract's state
- transition function: transaction tx transitions the system from state s to S'
  - tx is an invoke transaction and storage of S' is the result of executing the gtarget contract code with respect to previous statte S
  - tx is deploy transaction S' contains new contracts state and contract's address, and storage of S is updated according to the execution of the contract's constructor
- commitment is the root of teh binary merkle-patricia tree of height 251, 2 level structure where the contract address determines the path ffrom the root to the leafencoding the contract state
- contract_hash, storage_root 

### Contract State:
- contract code
- contract storage


### Run Examples:

Pull an example StarkNet block to verify
```
cd verify/
curl https://alpha4.starknet.io/feeder_gateway/get_block?blockNumber=145996 > rawStarkNetBlock.json
cd go/
go run *.go
```

#### Sources:
- https://starknet.io/glossary
- https://docs.starknet.io/docs/State/starknet-state
- https://aszepieniec.github.io/stark-anatomy