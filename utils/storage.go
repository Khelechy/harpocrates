package utils

import (
	"bytes"
	"fmt"
	"log"

	"github.com/sopherapps/go-scdb/scdb"
)

var tholos = "harpocrates_db"

var maxKeys uint64 = 1_000_000
var redundantBlocks uint16 = 1
var poolCapacity uint64 = 10
var compactionInterval uint32 = 1_800

func OpenStorage() *scdb.Store {

	vault, err := scdb.New(
		tholos,
		&maxKeys,
		&redundantBlocks,
		&poolCapacity,
		&compactionInterval,
		true, // isSearchEnabled EDIT(v0.2.0)16/01/2023
	)

	if err != nil {
		log.Fatalf("error opening store: %s", err)
	}

	return vault
}

func Set(key, value string) {

	vault := OpenStorage()
	defer func() {
		_ = vault.Close()
	}()

	byteKey := []byte(key)
	byteValue := []byte(value)

	err := vault.Set(byteKey, byteValue, nil)

	if err != nil {
		log.Fatalf("error inserting: %s", err)
	}
}

func Get(key string) {

	vault := OpenStorage()
	defer func() {
		_ = vault.Close()
	}()
	
	byteKey := []byte(key)

	byteValue, err := vault.Get(byteKey)

	if err != nil {
		log.Fatalf("error getting: %s", err)
	}

	value := bytes.NewBuffer(byteValue).String()

	fmt.Printf("Key: %s, Value: %s", key, value)
}
