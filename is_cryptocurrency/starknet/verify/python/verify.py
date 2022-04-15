import json
# import asyncio
from typing import Sequence

from starkware.cairo.lang.vm.crypto import pedersen_hash
from starkware.starknet.definitions.general_config import StarknetGeneralConfig
from starkware.starknet.services.api.feeder_gateway.block_hash import calculate_block_hash

async def calc_block(parent_hash, block_number, global_state_root, block_timestamp, tx_hashes, tx_signatures, event_hashes):
    hash = calculate_block_hash(StarknetGeneralConfig, parent_hash, block_number, global_state_root, block_timestamp, tx_hashes, tx_signatures, event_hashes, pedersen_hash)
    print("BLOCK HASH: ", hash)

with open('../rawStarkNetBlock145996.json') as f:
    data = json.load(f)

print(data["timestamp"])

txHashes = []
txSignatures = []
eventHashes = []

for i in data["transactions"]:
    txHashes.append(int(i["transaction_hash"], 16))
    if "signature" in i:
        if len(i["signature"]) == 2:
            innerSig = []
            innerSig.append(int(i["signature"][0], 16))
            innerSig.append(int(i["signature"][1], 16))
            txSignatures.append(innerSig)
        else:
            txSignatures.append([])
    else:
        txSignatures.append([])

for j in data["transaction_receipts"]:
    print("EVENTS ", j["events"])

print("HASHES: ", txHashes)
print("HASHES: ", txSignatures)

# loop = asyncio.get_event_loop()
# loop.run_until_complete(calc_block(
#     data["parent_block_hash"],
#     data["block_number"],
#     data["state_root"],
#     data["timestamp"],
#     data[]))
# loop.close()
