package cmd

/*
REED - Remote Envelope Encryption Dictionary
*/
import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yasseraboudkhil/reed/conf"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "reed",
}

func RootCommand() *cobra.Command {

	rootCmd.Flags().StringP("config", "c", "", `File to be used
	as the config. Include the full path and extension of the file.`)

	return rootCmd
}

func init() {
	cobra.OnInitialize(loadConfig)
}

func loadConfig() {

	configFile, err := rootCmd.Flags().GetString("config")
	if err != nil {
		log.Println("Call to rootCmd.Flags().GetString failed: " + err.Error())
		configFile = ""
	}

	if err := conf.LoadConfig(configFile); err != nil {
		log.Fatal("Failed to load config: " + err.Error())
	}

	fmt.Println("Started the Corba CLI")
	//fmt.Println(conf.C.Region)
	fmt.Println(conf.C)
}
