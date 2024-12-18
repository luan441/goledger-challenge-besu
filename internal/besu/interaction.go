package besu

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SimpleStorage struct {
	Format                 string      `json:"_format"`
	ContractName           string      `json:"contractName"`
	SourceName             string      `json:"sourceName"`
	Abi                    []Abi       `json:"abi"`
	Bytecode               string      `json:"bytecode"`
	DeployedBytecode       string      `json:"deployedBytecode"`
	LinkReferences         interface{} `json:"linkReferences"`
	DeployedLinkReferences interface{} `json:"deployedLinkReferences"`
}

type Abi struct {
	Inputs          []InputOutput `json:"inputs"`
	Name            string        `json:"name"`
	Outputs         []InputOutput `json:"outputs"`
	StateMutability string        `json:"stateMutability"`
	Type            string        `json:"type"`
}

type InputOutput struct {
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

func ExecContract(value int64) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("error retrieving directory: %v", err)
		return err
	}

	jsonFile, err := os.Open(fmt.Sprintf("%s/besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json", dir))
	if err != nil {
		log.Printf("error opening json file: %v", err)
		return err
	}
	defer jsonFile.Close()

	content, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("error reading json file: %v", err)
		return err
	}

	var simpleStorage SimpleStorage

	err = json.Unmarshal(content, &simpleStorage)
	if err != nil {
		log.Printf("error Unmarshal: %v", err)
		return err
	}

	abiJSON, err := json.Marshal(simpleStorage.Abi)
	if err != nil {
		log.Printf("error Marshal: %v", err)
		return err
	}

	abi, err := abi.JSON(strings.NewReader(string(abiJSON)))
	if err != nil {
		log.Printf("error parsing abi: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Printf("error dialing node: %v", err)
		return err
	}

	slog.Info("querying chain id")

	chainID, err := client.ChainID(ctx)
	if err != nil {
		log.Printf("error querying chain id: %v", err)
		return err
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b")

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	priv, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		log.Printf("error loading private key: %v", err)
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainID)
	if err != nil {
		log.Printf("error creating transactor: %v", err)
		return err
	}

	tx, err := boundContract.Transact(auth, "set", big.NewInt(value))
	if err != nil {
		log.Printf("error transacting: %v", err)
		return err
	}

	fmt.Println("waiting until transaction is mined",
		"tx", tx.Hash().Hex(),
	)

	receipt, err := bind.WaitMined(
		context.Background(),
		client,
		tx,
	)
	if err != nil {
		log.Printf("error waiting for transaction to be mined: %v", err)
		return err
	}

	fmt.Printf("transaction mined: %v\n", receipt)
	return nil
}

func CallContract() (int64, error) {
	var result *big.Int
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("error retrieving directory: %v", err)
		return 0, err
	}

	jsonFile, err := os.Open(fmt.Sprintf("%s/besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json", dir))
	if err != nil {
		log.Printf("error opening json file: %v", err)
		return 0, err
	}
	defer jsonFile.Close()

	content, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("error reading json file: %v", err)
		return 0, err
	}

	var simpleStorage SimpleStorage

	err = json.Unmarshal(content, &simpleStorage)
	if err != nil {
		log.Printf("error Unmarshal: %v", err)
		return 0, err
	}

	abiJSON, err := json.Marshal(simpleStorage.Abi)
	if err != nil {
		log.Printf("error Marshal: %v", err)
		return 0, err
	}

	abi, err := abi.JSON(strings.NewReader(string(abiJSON)))
	if err != nil {
		log.Printf("error parsing abi: %v", err)
		return 0, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Printf("error connecting to eth client: %v", err)
		return 0, err
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x42699A7612A82f1d9C36148af9C77354759b210b")
	caller := bind.CallOpts{
		Pending: false,
		Context: ctx,
	}

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	var output []interface{}
	err = boundContract.Call(&caller, &output, "get")
	if err != nil {
		log.Printf("error calling contract: %v", err)
		return 0, err
	}
	result = output[0].(*big.Int)

	fmt.Println("Successfully called contract!", result)
	return result.Int64(), nil
}
