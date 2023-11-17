package harpocrates

import (
	"encoding/json"
	"log"
	"os"

	"github.com/khelechy/harpocrates/utils"

	"github.com/google/uuid"
)

type Harpocrates struct {
}

var initializeHashKey = "tholos_initialize_key"
var unsealHashKey = "tholos_unseal_key"

func Initialize(part int) {
	// Create keys and split ( into parts )

	if part <= 0 {
		part = 3
	}
	seedingSecret := uuid.New().String()
	secrets, err := utils.SplitSecret(seedingSecret, part, part)

	if err != nil {
		log.Fatalf("Error splitting secrets: %s", err)
	}

	// Save keys in a Json file
	data := map[string]interface{}{"keys": secrets}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
		return
	}

	err = os.WriteFile("keys.json", jsonData, 0644)
	if err != nil {
		log.Fatal("Error writing to file:", err)
		return
	}
}

func Unseal() {

	// Combine Keys

	// Set localStorageUnseal value to true
}

func GetItem() {

	// Check localStorage is unsealed

	// Get Item
}

func SetItem() {

	// Check localStorage is unsealed

	// Set Item
}
