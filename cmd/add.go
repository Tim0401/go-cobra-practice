/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: add,
}

func add(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)
	var result int
	for {
		scanner.Scan()
		in := scanner.Text()
		i, _ := strconv.Atoi(in)
		result += i
		fmt.Printf("result %v\n", result)
	}
	// var result int
	// for _, v := range args {
	// 	i, _ := strconv.Atoi(v)
	// 	result += i
	// }
	// fmt.Printf("hello, world!: %v", result)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
