import json
import asyncio
import requests

from starkware.cairo.lang.vm.crypto import pedersen_hash
from starkware.starknet.definitions.general_config import StarknetGeneralConfig
from starkware.starknet.services.api.feeder_gateway.block_hash import calculate_block_hash, calculate_event_hash

async def calc_block(parent_hash, block_number, global_state_root, block_timestamp, tx_hashes, tx_signatures, event_hashes, block_hash):
    hash = await calculate_block_hash(
        StarknetGeneralConfig,
        parent_hash,
        block_number,
        global_state_root,
        block_timestamp,
        tx_hashes,
        tx_signatures,
        event_hashes,
        pedersen_hash
    )
    print("Fetched Hash: ", block_hash)
    print("Calculated Hash: 0x%x" % (hash))
    print("Match: " , int(block_hash, 16) == hash)
    return

url = requests.get("https://alpha4.starknet.io/feeder_gateway/get_block?blockNumber=157908")
text = url.text
data = json.loads(text)

parentHash = int(data["parent_block_hash"], 16)
blockNum = data["block_number"]
stateRoot = bytes.fromhex(data["state_root"])
ts = data["timestamp"]
blockHash = data["block_hash"]
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
    for e in j["events"]:
        keys = []
        for k in e["keys"]:
            keys.append(int(k, 16))

        data = []
        for d in e["data"]:
            data.append(int(d, 16))

        evHash = calculate_event_hash(
            int(e["from_address"], 16),
            keys,
            data,
            pedersen_hash
        )
        eventHashes.append(evHash)

loop = asyncio.get_event_loop()
loop.run_until_complete(calc_block(
        parentHash,
        blockNum,
        stateRoot,
        ts,
        txHashes,
        txSignatures,
        eventHashes,
        blockHash
    ))
loop.close()
