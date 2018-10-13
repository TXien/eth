package infolib

import (
        "context"
        //"crypto/elliptic"
        //"encoding/hex"
        //"fmt"
        "log"
        //"math/big"
        //"github.com/ethereum/go-ethereum/crypto/secp256k1"
        "github.com/ethereum/go-ethereum/common"
        //"github.com/ethereum/go-ethereum/crypto"
        "github.com/ethereum/go-ethereum/ethclient"
        //"os"
        //"time"
        //"reflect"
)

func GetBalance(client *ethclient.Client, address string)(string){
        header, err := client.BalanceAt(context.Background(),common.HexToAddress(address), nil)
        if err != nil {
                //log.Fatal(err)
        }
        return header.String()
}

func BlockNumber(client *ethclient.Client)(string){
        header, err := client.HeaderByNumber(context.Background(), nil)
        if err != nil {
                log.Fatal(err)
        }
        return header.Number.String()
}

func NonceAt(client *ethclient.Client, addresses string)(uint64){
	address := common.HexToAddress(addresses)
        header, err := client.NonceAt(context.Background(), address, nil)
        if err != nil {
                log.Fatal(err)
        }
        return header
}

