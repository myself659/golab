package main

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// 产生私钥
	key, _ := crypto.GenerateKey()
	fmt.Println("key:", key)
	// key的内容是什么

	// 根据私钥产生签名
	auth := bind.NewKeyedTransactor(key)
	fmt.Println("auth:", auth)
	//
	// alloc记录balance情况

	alloc := make(core.GenesisAlloc)
	// 给地址发放token
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(133700000)}
	// 模拟后端，正常可以调用以太坊公链
	sim := backends.NewSimulatedBackend(alloc)

	// deploy contract
	addr, _, contract, err := DeployWinnerTakesAll(auth,
		sim,
		big.NewInt(10),
		big.NewInt(time.Now().Add(2*time.Minute).Unix()),
		big.NewInt(time.Now().Add(5*time.Minute).Unix()))
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}

	// interact with contract
	fmt.Printf("Contract deployed to %s\n", addr.String())
	deadlineCampaign, _ := contract.DeadlineCampaign(nil)
	fmt.Printf("Pre-mining Campaign Deadline: %s\n", deadlineCampaign)

	fmt.Println("Mining...")
	// simulate mining
	sim.Commit()

	postDeadlineCampaign, _ := contract.DeadlineCampaign(nil)
	fmt.Printf("Post-mining Campaign Deadline: %s\n", time.Unix(postDeadlineCampaign.Int64(), 0))

	// create a project
	numOfProjects, _ := contract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects before: %d\n", numOfProjects)

	fmt.Println("Adding new project...")
	contract.SubmitProject(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: big.NewInt(2381623),
		Value:    big.NewInt(10),
	}, "test project", "http://www.example.com")

	fmt.Println("Mining...")
	sim.Commit()

	numOfProjects, _ = contract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects after: %d\n", numOfProjects)
	info, _ := contract.GetProjectInfo(nil, auth.From)
	fmt.Printf("Project Info: %v\n", info)

	// instantiate deployed contract
	fmt.Printf("Instantiating contract at address %s...\n", auth.From.String())
	instContract, err := NewWinnerTakesAll(addr, sim)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}
	numOfProjects, _ = instContract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects of instantiated Contract: %d\n", numOfProjects)
}
