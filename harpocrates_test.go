package harpocrates

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/khelechy/harpocrates/utils"
	"github.com/stretchr/testify/assert"
)

func TestMount(t *testing.T){

	// Override the file writing path
	oldPath := "keys.json"
	defer func() {
		// Clean up: restore the original file writing path
		_ = os.Rename("keys.json", oldPath)
	}()

	// Call the Mount function
	Mount(5)

	// Check if keys.json file is created
	_, err := os.Stat("keys.json")
	assert.NoError(t, err, "keys.json file should be created")

	// Check if harpocrates_db folder is created
	_, err = os.Stat("harpocrates_db")
	assert.NoError(t, err, "harpocrates_db folder should be created")

	// Read the content of the keys.json file
	fileContent, err := os.ReadFile("keys.json")
	assert.NoError(t, err, "Error reading keys.json file")

	// Unmarshal the JSON content
	var jsonData map[string]interface{}
	err = json.Unmarshal(fileContent, &jsonData)
	assert.NoError(t, err, "Error unmarshaling JSON")

	// Check if the "keys" key exists in the JSON data
	keys, ok := jsonData["keys"].([]interface{})
	assert.True(t, ok, "keys should be a slice in the JSON data")

	// Check if the number of keys is correct
	assert.Equal(t, 5, len(keys), "Number of keys should be 5")

	stripTestFiles()
}


func TestUnseal(t *testing.T) {

	keys := prepareTestEnv()

	Unseal([]string{keys[0].(string), keys[1].(string), keys[2].(string)})

	want := "1"

	result := utils.GetItem("tholos_seal_key")

	assert.Equal(t, want, result)

	stripTestFiles()
}

func TestSeal(t *testing.T) {


	keys := prepareTestEnv()

	Unseal([]string{keys[0].(string), keys[1].(string), keys[2].(string)})

	Seal([]string{keys[0].(string), keys[1].(string), keys[2].(string)})

	want := "0"

	result := utils.GetItem("tholos_seal_key")

	assert.Equal(t, want, result)

	stripTestFiles()
}

func TestGet(t *testing.T){

	keys := prepareTestEnv()

	Unseal([]string{keys[0].(string), keys[1].(string), keys[2].(string)})

	want := "kelechi"

	utils.SetItem("name", want)

	result := Get("name")

	assert.Equal(t, want, result)

	stripTestFiles()
}

func TestSet(t *testing.T){
	keys := prepareTestEnv()

	Unseal([]string{keys[0].(string), keys[1].(string), keys[2].(string)})

	want := "kelechi"

	Set("name", "kelechi")

	result := utils.GetItem("name")

	assert.Equal(t, want, result)

	stripTestFiles()
}





func prepareTestEnv() []interface{}{
	Mount(5)
	_, _ = os.Stat("harpocrates_db")

	// Read the content of the keys.json file
	fileContent, _ := os.ReadFile("keys.json")

	// Unmarshal the JSON content
	var jsonData map[string]interface{}
	_ = json.Unmarshal(fileContent, &jsonData)

	// Check if the "keys" key exists in the JSON data
	keys, _ := jsonData["keys"].([]interface{})

	return keys
}

func stripTestFiles(){
	os.RemoveAll("harpocrates_db")
	os.Remove("keys.json")
}