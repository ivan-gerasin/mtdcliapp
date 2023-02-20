package mtdcli

import (
	"fmt"
	"github.com/ivan-gerasin/mtdcore"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "mtd",
	Short: "Simple todo list",
	Long:  `As I said: simple todo list. MTD stands for "My To Do"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command executed")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of MTD",
	Long:  `All software has versions. This is MTD's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("My To Do (MTD) v0.0")
	},
}

var addItemCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new item into list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		priority, _ := cmd.Flags().GetInt8("priority")
		fmt.Println("Added: " + strings.Join(args, " "))
		mtdCore.AddItem(strings.Join(args, " "), priority)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all todo items as a list",
	Run: func(cmd *cobra.Command, args []string) {
		list := mtdCore.List()
		priority, _ := cmd.Flags().GetBool("priority")
		Render(list, false, priority)
	},
}

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark given item as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		mtdCore.Done(num)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addItemCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	addItemCmd.PersistentFlags().Int8P("priority", "p", 0, "Set priority of item")
	listCmd.PersistentFlags().BoolP("priority", "p", false, "Sort by priority")
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	//rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	//rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	//rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	//viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

//
//func initConfig() {
//	// Don't forget to read config either from cfgFile or from home directory!
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//
//		// Search config in home directory with name ".cobra" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cobra")
//	}
//
//	if err := viper.ReadInConfig(); err != nil {
//		fmt.Println("Can't read config:", err)
//		os.Exit(1)
//	}
//}
