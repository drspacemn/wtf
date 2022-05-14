## WTF is Ethereum?

### Notes:
- Secure decentralized generalized transaction ledge
- Simple application on a decentralized, but singleton, compute resource, transactional singleton machine with shared-state
- Plurality of resources each with a distinct state and op code, able to interact through message-passing framework
- decentralized value-transfer system that can be shared with the whole world and virtually free
- cryptographically secure, transaction-based state machine
- end-developer a tightly integrated end-to-end system for building sofware, trustful object messaging comput framework
- facilitate transactions between consenting individauls
- State is made up of objects called "accounts" with each account having a 20-byte address
- state transition are direct transferes of value and information between accounts
- Account contains:
  - nonce counter used to make sure each transaction can only be processed once
  - ether balance
  - contract code, if present
  - storage(empty by default)
- 2 Types of accounts: externally owned accounts controlled by private keys and contract accounts controlled by their contract code
- contract account, every time the contract accoiunt receives a message its code activatesallowing it to read and write to internal storage
- and send other messages or create contracts in turn
- "transaction" refers to a signed data package that stores a message to be sent from and externally owned account
- Transaction:
  - recipient of the message
  - signature identifying the sender
  - amount of ether to transfer from the sender to the recipient
  - optional data field 
  - STARTGAS valuerepresenting a maximum number of computational steps the transaction execution is allowed to take
  - GASPRICE value representing the fee the sender pays per computation
- Data field has no function by default, but the virtual machine has an opcode using
- each transaction is required to set a limit to how many computational steps of code execution it can use. fundamental unit of computation is "gas"
- Usually a computational step costs 1 gas, but can cost more gas because they are more computationally expensive. 
- Messages are virtual objects that are never serialized and exist only in the ethereum execution environment
- The code in Ethereum contracts is written in low-level stack based bytecode language aka "ethereum virtual machine code" or "EVM code"
- Series of bytes where each byte represents an operation. Code execution is an infinite loop taht consists of repeatedly carying out the operation the operation at the current program counter(which begins at zero)
- and then incrementing the program counter by one until the end of the code is reached or an error or STOP or RETURN instruction is detected
- The stack a last in first out container to which values can be pushed and popped
- Memory an infinitely expandable byte array
- Contract's long-term storage, key/value store. Unlike stack ang memory reset after computation ends, storage persists for the long term
- Code can also access the value, sender and data of the incoming message and block header data. Code can return a byte array of thedata as an output
- While the EVM is running it's full computational stat can be defined by the tuple (block_state, transaction, message, code, memory, stack, pc, gas)
  - block_state global state containing all accounts and includes balances and storages
  - at the start of every round of execution the current instruction is found by taking the pcth byte of the code
  - and each instruction has its own definition in terms of how it affects the tuple. ADD pops two items off the stack and pushes their sum , reeduces gas by q and increments pc by 1
  - SSTORE pushes the top two items off the stack and inserts the second item in the contracts storage at the index specified by the first item
  - can optimize EVM execution via just-in-time compilation
- Unlike bitcoin ethereum blocks contain a copy of both the transaction list and the most recent state
- Patricia tree(trie) is used to accomplish this, nodes to be inserted and deleted and not just changed.
- transaction based state machine
- 1 ETHER is 10^18 wei
- 1 GWEI is 1 Billion wei

### RLP Encoding
- Recursive Length Prefix encoding
- encode arbitrarily nested arrays of binary data
- serialization in Ethereum

### Merkle Patricia Trie
- persistent data storage, to map between arbitrary-length binary data
  - mutable data structure to map between 256-bit binary fragments
  - typically implemented as a database
  - provide single value that identifies given set of key-value pairs
- Merkle Patricia Trie allows us to verify data integrity
- One can compute Merkle Root Hash of the trie with the Hash function such that any key-value pair was updated
  - the merkle root hash of the trie would be different 
- trie allows us to verify the inclusion of a key value pair without the access to the entire key-value pairs
- trie can provide a proof to prove that a certain key-value pair is included in a key-value mapping that produces a certain merkle root hash
- world state is a key-value mapping, keys are each account address, values are the balances for each account
- light client has merkle root hash for block
- full node can provide merkle proof which contains the merkle root hash, account key, and it's valance value as well as other data
- merkle proof allows light client to verify correctness on its own
- a light cleint can ask for a merkle root hash of the trie state, and use it to verify balance of its account
- light client only needs to trust the merle root hash
- small merkle root hash can be used to verify the state of a giant key-value mapping
- Ethereum has 3 Tries: Transaction Trie, Receipt Trie, and State Trie
  - each block header includes the three hashes, transactionRoot, receiptRoot, stateRoot
  - key is the hex form of the bytes from the RLP Encoding
  - value for key 80 is the result of the RLP Encoding
- EmptyNode, LeafNode, BranchNode
  - 
#### ELI MERKLE PATRICIA
- ordered tree data structure used to store dynamic set of associated array keys usually strings
- Node's position in th etree defines it skey 

WHAT ARE THE KEYS AND VALUES OF THE TRIE?
- the RLP encoding of an unsigned interger

THE PROCESS OF EXECUTING CONTRACT CODE IS PART OF THE DEFINITION OF THE STATE TRANSITION FUNCTION
- so if a transaction is added into block B the code execution spawned by that transaction will be executed by all nodes now and in the future

Full Node Size: 700 GB
Archive Node Size: 10 TB

#### Sources:
- https://ethereum.org/en/whitepaper
- https://ethereum.github.io/yellowpaper/paper.pdf
- https://ethereum.org/en/developers/docs/evm/opcodes
- https://jellopaper.org/evm
- https://medium.com/hackernoon/getting-deep-into-ethereum-how-data-is-stored-in-ethereum-e3f669d96033

