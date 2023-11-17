package utils

import (
	"bytes"
	"log"

	"github.com/sopherapps/go-scdb/scdb"
)

var tholos = "../harpocrates_db"

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
		log.Fatalf("Error opening store: %s", err)
	}

	return vault
}

func SetItem(key, value string) {

	vault := OpenStorage()
	defer func() {
		_ = vault.Close()
	}()

	byteKey := []byte(key)
	byteValue := []byte(value)

	if err := vault.Set(byteKey, byteValue, nil); err != nil {
		log.Fatalf("Error inserting: %s", err)
	}
}

func GetItem(key string) string {

	vault := OpenStorage()
	defer func() {
		_ = vault.Close()
	}()

	byteKey := []byte(key)

	byteValue, err := vault.Get(byteKey)

	if err != nil {
		log.Fatalf("Error getting: %s", err)
	}

	value := bytes.NewBuffer(byteValue).String()

	return value
}
