package harpocrates

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/khelechy/harpocrates/utils"

	"github.com/google/uuid"
	"github.com/fatih/color"
)

var mountHashKey = "tholos_mount_key"
var unsealHashKey = "tholos_unseal_key"
var seedingHashKey = "tholos_seeding_key"

//Errors 
var errNotMounted = "Harpocrates vault is not mounted"
var errIsSealed = "Harpocrates vault is sealed"

func Mount(part int) {
	// Create keys and split ( into parts )

	if part <= 5 {
		part = 5
	}
	seedkey := uuid.New().String()
	seedingSecret := strings.ReplaceAll(seedkey, "-", "")
	secrets, err := utils.SplitSecret(seedingSecret, part, 3)

	if err != nil {
		color.Red("Error splitting secrets: %s", err)
		return
	}

	// Save keys in a Json file
	data := map[string]interface{}{"keys": secrets}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		color.Red("Error marshaling JSON:", err)
		return
	}

	if err = os.WriteFile("keys.json", jsonData, 0644); err != nil {
		color.Red("Error writing to file:", err)
		return
	}

	// Set localStorageUnseal value to true
	utils.SetItem(mountHashKey, "1")
	utils.SetItem(seedingHashKey, seedingSecret)

	color.Green("Harpocrates vault is succesfully mounted")
}

func Unseal(secrets []string) {

	// Combine Keys
	combinedSeedingSecret, err := utils.CombineSecret(secrets)

	if err != nil {
		color.Red("Error unsealing vault:", err)
		return
	}

	seedingSecret := utils.GetItem(seedingHashKey)

	ValidateSharedKeys(combinedSeedingSecret, seedingSecret)

	if isUnsealed() {
		color.Yellow("Harpocrates vault is already unsealed")
		return
	}

	// Set localStorageUnseal value to true
	utils.SetItem(unsealHashKey, "1")

	color.Green("Harpocrates vault is unsealed successfully")
}

func Seal(secrets []string) {

	// Combine Keys
	combinedSeedingSecret, err := utils.CombineSecret(secrets)
	if err != nil {
		color.Red("Error sealing harpocrates vault:", err)
		return
	}

	seedingSecret := utils.GetItem(seedingHashKey)

	ValidateSharedKeys(combinedSeedingSecret, seedingSecret)

	// Set localStorageUnseal value to false
	utils.SetItem(unsealHashKey, "0")

}

func Get(key string) string {

	if isMounted() == false {
		color.Red(errNotMounted)
	}

	// Check localStorage is unsealed
	if isUnsealed() {
		return utils.GetItem(key)
	}

	color.Red(errIsSealed)
	return ""
}

func Set(key, value string) {

	if isMounted() == false {
		color.Red(errNotMounted)
		return
	}

	// Check localStorage is unsealed
	if isUnsealed() {
		utils.SetItem(key, value)
		return
	}

	color.Red(errIsSealed)
	return
}

func ValidateSharedKeys(combinedSeedingSecret, seedingSecret string) {
	if combinedSeedingSecret != seedingSecret {
		color.Red("Error unsealing vault: ", errors.New("Mismatched keys"))
		return
	}
}

func isMounted() bool {
	isMounted := utils.GetItem(mountHashKey)
	if isMounted != "" && isMounted == "1" {
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
