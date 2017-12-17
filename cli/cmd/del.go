// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/RowlingWu/agenda/entity"
	"log"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete current user",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		// read username from curUser.txt
		log.Println("read info of the current user...")
		name, ok := entity.ReadCur()
		if ok == 1 {
			log.Fatal("fatal: please log in first")
		}
		if ok == 2 {
			log.Fatal("failed to delete current user")
		}

		// clear curUser.txt
		log.Println("delete current user...")
		// seek userInfo in userInfo.txt and delete it
		f := entity.SeekUsr(name)
		if !f {
			log.Fatal("failed to delete current user")
		}
		log.Println("success in deleting current user")
	},
}

func init() {
	RootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
