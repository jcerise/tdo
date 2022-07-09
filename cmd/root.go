package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"tdo/todo"
)

var dataFile string
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tdo",
	Short: "A simple Go CLI Todo app",
	Long:  `With Tdo, all your dreams will come true`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigName(".tdo_config")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	//viper.AutomaticEnv()

	//viper.SetEnvPrefix("tdo")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("config file not found")
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println("Datafile defined as:", viper.GetString("datafile"))
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(todo.OpenOrCreateDataFile)

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+string(os.PathSeparator)+".tdo.json", "data file to store Todos")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tdo.yaml")
}
