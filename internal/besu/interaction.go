package besu

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ExecContract() {
	abi, err := abi.JSON(strings.NewReader("REPLACE: abi JSON as string goes here")) // found under besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "REPLACE: network URL") // e.g., http://localhost:8545
	if err != nil {
		log.Fatalf("error dialing node: %v", err)
	}

	slog.Info("querying chain id")

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("error querying chain id: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("REPLACE: contract address") // will be returned during startDev.sh execution

	boundContract := bind.NewBoundContract(
		contractAddress,
		abi,
		client,
		client,
		client,
	)

	priv, err := crypto.HexToECDSA("REPLACE: private key") // this can be found in the genesis.json file
	if err != nil {
		log.Fatalf("error loading private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priv, chainId)
	if err != nil {
		log.Fatalf("error creating transactor: %v", err)
	}

	tx, err := boundContract.Transact(auth, "REPLACE: method name")
	if err != nil {
		log.Fatalf("error transacting: %v", err)
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
		log.Fatalf("error waiting for transaction to be mined: %v", err)
	}

	fmt.Printf("transaction mined: %v\n", receipt)
}

func CallContract() {
	var result interface{}

	abi, err := abi.JSON(strings.NewReader("REPLACE: abi JSON as string goes here")) // found under besu/artifacts/contracts/SimpleStorage.sol/SimpleStorage.json
	if err != nil {
		log.Fatalf("error parsing abi: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "REPLACE: network URL") // e.g., http://localhost:8545
	if err != nil {
		log.Fatalf("error connecting to eth client: %v", err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("REPLACE: contract address") // will be returned during startDev.sh execution
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
	err = boundContract.Call(&caller, &output, "REPLACE: method name")
	if err != nil {
		log.Fatalf("error calling contract: %v", err)
	}
	result = output

	fmt.Println("Successfully called contract!", result)
}
