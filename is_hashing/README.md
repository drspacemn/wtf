## WTF is a "hash"?


### WTF is a Merkle Tree?
- ask for file by hash not by name
- hash list, list of hashes which are hashed again to get a single hash representing the lsit
- binary tree where each node is a  hash over its child nodes
- efficiently recalculate our hash by rehashing just the part of the tree that changed

#### Sources
- https://hackmd.io/@benjaminion/bls12-381