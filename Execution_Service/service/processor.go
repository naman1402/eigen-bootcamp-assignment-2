package services

import (
	"math/big"
	"encoding/hex"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/naman1402/eigen-bootcamp-assignment-2/Execution_Service/config"
)

// Init loads config
func Init() {
	config.Init()
}

type TaskParams struct {
	ProofOfTask      string
	AchievementData  string
	TaskDefinitionId int
	PerformerAddress string
	Signature        string
}

// SendTask submits the proof and achievement data to the AVS network
func SendTask(proofOfTask string, achievementData string, taskDefinitionId int) {
	privateKey, err := crypto.HexToECDSA(config.PrivateKey)
	if err != nil {
		log.Println("[SendTask] Invalid private key:", err)
		return
	}
	publicKey, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Println("[SendTask] Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	performerAddress := crypto.PubkeyToAddress(*publicKey).Hex()

	// Prepare ABI arguments for encoding
	arguments := abi.Arguments{
		{Type: abi.Type{T: abi.StringTy}},
		{Type: abi.Type{T: abi.BytesTy}},
		{Type: abi.Type{T: abi.AddressTy}},
		{Type: abi.Type{T: abi.UintTy}},
	}

	dataPacked, err := arguments.Pack(
		proofOfTask,
		[]byte(achievementData),
		common.HexToAddress(performerAddress),
		big.NewInt(int64(taskDefinitionId)),
	)
	if err != nil {
		log.Println("[SendTask] Error encoding arguments:", err)
		return
	}
	messageHash := crypto.Keccak256Hash(dataPacked)

	sig, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		log.Println("[SendTask] Error signing message:", err)
		return
	}
	sig[64] += 27
	serializedSignature := hexutil.Encode(sig)
	log.Println("[SendTask] Signature:", serializedSignature)

	client, err := rpc.Dial(config.OTHENTIC_CLIENT_RPC_ADDRESS)
	if err != nil {
		log.Println("[SendTask] Error connecting to RPC:", err)
		return
	}

	params := TaskParams{
		ProofOfTask:      proofOfTask,
		AchievementData:  "0x" + hex.EncodeToString([]byte(achievementData)),
		TaskDefinitionId: taskDefinitionId,
		PerformerAddress: performerAddress,
		Signature:        serializedSignature,
	}

	response := makeRPCRequest(client, params)
	log.Println("[SendTask] RPC response:", response)
}

// makeRPCRequest sends the sendTask RPC call
func makeRPCRequest(client *rpc.Client, params TaskParams) interface{} {
	var result interface{}

	err := client.Call(&result, "sendTask", params.ProofOfTask, params.AchievementData, params.TaskDefinitionId, params.PerformerAddress, params.Signature)
	if err != nil {
		log.Println("[makeRPCRequest] RPC call error:", err)
	}
	return result
}