package main

import (
	"os"
	"fmt"
	"math"
	"math/big"
	"encoding/json"

	"github.com/dontpanicdao/caigo"
)

var curve caigo.StarkCurve

// const SEQUENCER_ADDRESS = "5b98b836969a60fec50fa925905dd1d382a7db43"
const SEQUENCER_ADDRESS = "37b2cd6baaa515f520383bee7b7094f892f4c770695fc329a8973e841a971ae"

func main() {
	curve, _ = caigo.SC(caigo.WithConstants())

	rawBlockFile, err := os.ReadFile("../rawStarkNetBlock.json")
	if err != nil {
		panic(err.Error())
	}
	var block Block
	err = json.Unmarshal(rawBlockFile, &block)
	if err != nil {
		panic(err.Error())
	}

	hash, err := block.Hash()

	fmt.Println("HASH ERRR: ", hash, err)
}

/*
StarkNet Block Hash
h(B) = h(
	block_number,
	global_state_root,
	sequencer_address,
	block_timestamp,
	transaction_count,
	transaction_commitment,
	event_count,
	event_commitment,
	0,
	0,
	parent_block_hash
)
*/
func (b Block) Hash() (hash *big.Int, err error) {
	// txCommitement := b.TxCommitment()
	_ = b.TxCommitment()
	// eventCount, eventCommitment := b.EventDetails()



	return hash, err
}

func (b Block) TxCommitment() *big.Int {
	var merkleHold []*big.Int
	var merklePass []*big.Int

	for _, tx := range b.Transactions {
		if len(tx.Signature) == 2 {
			sigHash, _ := curve.PedersenHash([]*big.Int{caigo.HexToBN(tx.Signature[0]), caigo.HexToBN(tx.Signature[1])})
			txSigHash, _ := curve.PedersenHash([]*big.Int{caigo.HexToBN(tx.TransactionHash), sigHash})
			merkleHold = append(merkleHold, txSigHash)
		} else {
			merkleHold = append(merkleHold, caigo.HexToBN(tx.TransactionHash))
		}
	}

	height := int(math.Ceil(math.Log2(float64(len(merkleHold)))))

	

	fmt.Println("GET TX COMM: ", merkleHold)
	fmt.Println("GET TX COMM LEN: ", len(merkleHold), len(b.Transactions))
	return new(big.Int)
}

// func (b Block) EventDetails() (eventCount int64, eventCommitment *big.Int) {

// }
	
type Block struct {
	BlockHash       string `json:"block_hash"`
	ParentBlockHash string `json:"parent_block_hash"`
	BlockNumber     int64    `json:"block_number"`
	StateRoot       string `json:"state_root"`
	Status          string `json:"status"`
	Transactions    []struct {
		ContractAddress     string        `json:"contract_address"`
		ContractAddressSalt string        `json:"contract_address_salt,omitempty"`
		ConstructorCalldata []string      `json:"constructor_calldata,omitempty"`
		TransactionHash     string        `json:"transaction_hash"`
		Type                string        `json:"type"`
		EntryPointSelector  string        `json:"entry_point_selector,omitempty"`
		EntryPointType      string        `json:"entry_point_type,omitempty"`
		Calldata            []string      `json:"calldata,omitempty"`
		Signature           []string `json:"signature,omitempty"`
		MaxFee              string        `json:"max_fee,omitempty"`
	} `json:"transactions"`
	Timestamp           int64 `json:"timestamp"`
	TransactionReceipts []struct {
		TransactionIndex   int64           `json:"transaction_index"`
		TransactionHash    string        `json:"transaction_hash"`
		L2ToL1Messages     []interface{} `json:"l2_to_l1_messages"`
		Events             []interface{} `json:"events"`
		ExecutionResources struct {
			NSteps                 int64 `json:"n_steps"`
			BuiltinInstanceCounter struct {
				PedersenBuiltin   int64 `json:"pedersen_builtin"`
				RangeCheckBuiltin int64 `json:"range_check_builtin"`
				OutputBuiltin     int64 `json:"output_builtin"`
				EcdsaBuiltin      int64 `json:"ecdsa_builtin"`
				BitwiseBuiltin    int64 `json:"bitwise_builtin"`
				EcOpBuiltin       int64 `json:"ec_op_builtin"`
			} `json:"builtin_instance_counter"`
			NMemoryHoles int64 `json:"n_memory_holes"`
		} `json:"execution_resources"`
	} `json:"transaction_receipts"`
}
