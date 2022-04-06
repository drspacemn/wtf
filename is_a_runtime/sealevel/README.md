## WTF is Sealevel?
Parallel smart contract runtime. Single threaded=one contract at a time modifies teh blockchain state.
Solana transactions describe all the states a transaction will read or write while executing. Allows for non-overlapping transactions to execute concurrenlty also 
transactions that are only reading same sate to execute concurrenlty
Mapping of Public Keys to accounts, account maintain balacnces and data, vector of bytes
Programs can only change the data of accounts they own, programs can only debit accounts they own,any program can credit any account, any program can read any account
System program defaults to owner of a ll programs
User defined program is loaded by the loader program, loader program is able to mark data in accountabl eas executable
- create new public key
- transfer coin to the key
- tell system program to allocate memory
- tell system program to assign account to loader
- updload the byte code into the memory in pieces
- tell loader to mark the memory as executable
In KV store there is subset of keys that a program and only that program has write access to
Interfaces like readv or writev tell the kernel ahead of time all the memory the suer wants to read or write. Allows the OS to prefetch, prepare the device, 
and execte teh oepration conccurrently if the device allows it
each instruction tells the VM which accounts it wants to read and wrte ahead of time
- sort millions of pending transactions, schedule all the on-overlapping stransaction sin parallel

### Transactions:
Specify an instruction vector. Each instruction contains the program, program instruction, and list of accounts the transaction wants to read and write
Interface is inspired by low level Operating System interfaces to devices
- size_t \n readv(int d, const struct iovec *iov, int iovcnt); \n struct iovec { char *iov_base; size_t iov_len; \n}

### Single Instruction Multiple Data
- data parallelism: we preform the same computation but on different data..
- sort all the instructions by program ID
- run the same program over all accounts concurrently
- scalable array of multithreaded straming multiprocessors, CPU invokes a kernel grid, the blocks of the grid are enumerated and stributed ot 
- executes the instruction over 80 different inputs in parallel

### IDEAS: 
- delegation of what the account has ability to write feels like a DAG
- (TODO: write a device driver for rpi)

#### Sources:
- https://medium.com/solana-labs/sealevel-parallel-processing-thousands-of-smart-contracts-d814b378192