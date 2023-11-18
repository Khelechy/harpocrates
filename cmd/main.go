package main

import (
	"fmt"
	"log"
	"os"
	"encoding/json"
	"math/rand"

	"github.com/khelechy/harpocrates"

	"github.com/spf13/cobra"
)

type KeysData struct {
	Keys []string `json:"keys"`
}

var rootCmd = &cobra.Command{
	Use:   "harpocrates",
	Short: "A lightweight CLI tool for generating .gitignore files",
}

// Harpocrates mount

var mountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mounts Harpocrates vault and Generates distributed secrets to be shared",
	Run: func(cmd *cobra.Command, args []string) {
		partsFlag, _ := cmd.Flags().GetInt("parts")
		harpocrates.Mount(partsFlag)
	},
}

var getItemCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a string value using it key from the harpocrates vault",
	Run: func(cmd *cobra.Command, args []string) {
		keyFlag, _ := cmd.Flags().GetString("key")
		var key string

		if keyFlag != "" {
			key = keyFlag
			item := harpocrates.GetItem(key)
			fmt.Println(item)
		} else {
			log.Fatalln("Key flag is empty")
		}

	},
}

var setItemCmd = &cobra.Command{
	Use:   "set",
	Short: "Stores a string value against a key in the harpocrates vault",
	Run: func(cmd *cobra.Command, args []string) {

		keyFlag, _ := cmd.Flags().GetString("key")
		valueFlag, _ := cmd.Flags().GetString("value")

		var key string
		var value string

		if keyFlag != "" || valueFlag != "" {
			key = keyFlag
			value = valueFlag
			harpocrates.SetItem(key, value)
		} else {
			log.Fatalln("Key or Value flag is empty")
		}
	},
}

var sealCmd = &cobra.Command{
	Use: "seal",
	Short: "Seals the vault so that Set and Get operations would be blocked",
	Run: func(cmd *cobra.Command, args []string) {
		jsonFilePathFlag, _ := cmd.Flags().GetString("path")

		if jsonFilePathFlag != "" {
			keys := decodeJsonAndSelectKeys(jsonFilePathFlag)
			harpocrates.Seal(keys)
		}else{
			log.Fatalln("Json file path is empty")
		}
	},
}

var unSealCmd = &cobra.Command{
	Use: "unseal",
	Short: "Unseals the vault so that Set and Get operations would be unblocked",
	Run: func(cmd *cobra.Command, args []string) {
		jsonFilePathFlag, _ := cmd.Flags().GetString("path")

		if jsonFilePathFlag != "" {
			keys := decodeJsonAndSelectKeys(jsonFilePathFlag)
			harpocrates.Unseal(keys)
		}else{
			log.Fatalln("Json file path is empty")
		}
	},
}

func decodeJsonAndSelectKeys(jsonFilePath string) []string{

	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalln("Error reading JSON file:", err)
		return nil
	}

	// Unmarshal JSON data into KeysData struct
	var keysData KeysData
	if err = json.Unmarshal(jsonData, &keysData); err != nil{
		log.Fatalln("Error unmarshalling JSON:", err)
		return nil
	}

	// Randomly select 3 keys from the array
	selectedKeys := make([]string, 3)
	for i := 0; i < 3; i++ {
		randomIndex := rand.Intn(len(keysData.Keys))
		selectedKeys[i] = keysData.Keys[randomIndex]
	}

	return selectedKeys
}

func init() {
	rootCmd.AddCommand(mountCmd)
	rootCmd.AddCommand(sealCmd)
	rootCmd.AddCommand(unSealCmd)
	rootCmd.AddCommand(getItemCmd)
	rootCmd.AddCommand(setItemCmd)

	mountCmd.Flags().Int("parts", 5, "The number of parts the secret should be shared into")
	getItemCmd.Flags().String("key", "", "The key of the associated value to retrieve")
	setItemCmd.Flags().String("key", "", "The key of the associated value to store")
	setItemCmd.Flags().String("value", "", "The value of the associated key to store")
}

func main() {

	// if err := rootCmd.Execute(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	//harpocrates.Mount(5)

	var keys = []string{"5105c79604545210e5fd30c1269ac1189e748aea93bae50db2d36bde8886adc143", "db206494f721b82e307417f54d18fc9396f12c067fd9d117a1f2e30fc5fb99f429", "94c7bcdf37cd0c2d48b19bf334b6e4c352a2729c55721fc9ef18aebfd94efe3ff8"}

	harpocrates.Unseal(keys)

	//harpocrates.Seal(keys)

	//harpocrates.SetItem("name", "kelechi")
	//ss := harpocrates.GetItem("name")

	//fmt.Println(ss)
}
