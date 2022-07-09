/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sort"
	"strconv"
	"tdo/todo"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing Todo item",
	Run:   editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	items, _ := todo.ReadItems(viper.GetString("datafile"))
	if item > 0 && item <= len(items) {
		updatedItem := &items[item-1]

		if newPriority != 0 {
			updatedItem.Priority = newPriority
		}

		updatedItem.Text = message

		fmt.Printf("Item %v %v\n", strconv.Itoa(item), " has been updated")

		sort.Sort(todo.ByPriority(items))
		todo.SaveItems(viper.GetString("datafile"), items)

	} else {
		fmt.Fprintln(cmd.OutOrStdout(), item, "doesn't match any items")
	}
}

var message string
var item int
var newPriority int

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringVarP(&message, "message", "m", "", "Updated Todo item message")
	editCmd.Flags().IntVarP(&item, "item", "i", 0, "Todo item index to edit")
	editCmd.Flags().IntVarP(&newPriority, "priority", "p", 0, "Priority: 1, 2, 3")

	editCmd.MarkFlagRequired("message")
	editCmd.MarkFlagRequired("item")
}
