package harpocrates

import (
	"encoding/json"
	"log"
	"os"

	"github.com/khelechy/harpocrates/utils"

	"github.com/google/uuid"
)

var mountHashKey = "tholos_mount_key"
var unsealHashKey = "tholos_unseal_key"

func Mount(part int) {
	// Create keys and split ( into parts )

	if part <= 5 {
		part = 5
	}
	seedingSecret := uuid.New().String()
	secrets, err := utils.SplitSecret(seedingSecret, part, 3)

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

	if err = os.WriteFile("../keys.json", jsonData, 0644); err != nil {
		log.Fatal("Error writing to file:", err)
		return
	}

	// Set localStorageUnseal value to true
	utils.SetItem(mountHashKey, "1")
}

func Unseal(secrets []string) {

	// Combine Keys
	_, err := utils.CombineSecret(secrets)
	if err != nil {
		log.Fatal("Error unsealing vault:", err)
		return
	}

	if isUnsealed() {
		log.Println("Vault has already been unsealed")
	}

	// Set localStorageUnseal value to true
	utils.SetItem(unsealHashKey, "1")
}

func Seal(secrets []string) {
	// Combine Keys
	_, err := utils.CombineSecret(secrets)
	if err != nil {
		log.Fatal("Error sealing vault:", err)
		return
	}

	// Set localStorageUnseal value to false
	utils.SetItem(unsealHashKey, "0")

}

func GetItem(key string) string {

	if isInitialized() == false {
		log.Fatal("Vault has not been initialized")
	}

	// Check localStorage is unsealed
	if isUnsealed() {
		return utils.GetItem(key)
	}

	log.Fatal("Vault has not been unsealed")
	return ""
}

func SetItem(key, value string) {

	if isInitialized() == false {
		log.Fatal("Vault has not been initialized")
	}

	// Check localStorage is unsealed
	if isUnsealed() {
		utils.SetItem(key, value)
		return
	}

	log.Fatal("Vault has not been unsealed")
}

func isInitialized() bool {
	isInitialized := utils.GetItem(mountHashKey)
	if isInitialized != "" && isInitialized == "1" {
		return true
	}

	return false
}

func isUnsealed() bool {
	isUnsealed := utils.GetItem(unsealHashKey)
	if isUnsealed != "" && isUnsealed == "1" {
		return true
	}

	return false
}
