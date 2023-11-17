package main

import (
	"fmt"
	"log"

	"github.com/khelechy/harpocrates"

	"github.com/spf13/cobra"
)

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

func init() {
	rootCmd.AddCommand(mountCmd)
	rootCmd.AddCommand(getItemCmd)
	rootCmd.AddCommand(setItemCmd)

	mountCmd.Flags().Int("parts", 5, "The number of parts the secret should be shared into")
	getItemCmd.Flags().String("key", "", "The key of the associated value to retrieve")
	setItemCmd.Flags().String("key", "", "The key of the associated value to store")
	setItemCmd.Flags().String("value", "", "The value of the associated key to store")
}

func main() {

	//harpocrates.Initialize(5)

	//var keys = []string{"aa6ab8828ae248ab7be63ca731ea0adee52a713f97ef76872fd9a8df8d8012ea4c66c8845a", "3499ac76ef82ae119e5b4be0b45d98da86558ea75326eae05b410e0bd78800ba1e9a42f16e", "9c6ed4f5d10644e86871372dddabb23b33c8a1f87c74acd919da49147859d42c786abee5e8"}

	//harpocrates.Unseal(keys)

	//harpocrates.Seal(keys)

	//harpocrates.SetItem("name", "kelechi")
	ss := harpocrates.GetItem("name")

	fmt.Println(ss)
}
