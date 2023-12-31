package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Run() {
	//todo validation

	addBlockCommand := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCommand := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCommand.String("data", "", "Block Data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCommand.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCommand.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCommand.Parsed() {
		if *addBlockData == "" {
			addBlockCommand.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCommand.Parsed() {
		cli.printChain()
	}
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addblock -data BLOCK_DATA - add a block to the blockchain")
	fmt.Println("  printchain - print all the blocks of the blockchain")
}
