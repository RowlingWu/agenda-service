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
	"encoding/json"
	"fmt"

	"github.com/RowlingWu/agenda-service/entity"
	"github.com/spf13/cobra"
)

// quCmd represents the qu command
var quCmd = &cobra.Command{
	Use:   "qu",
	Short: "to find user infomation ",
	Long:  `All user's information.`,
	Run: func(cmd *cobra.Command, args []string) {

		getUsersURL := testAllUsersURL
		b := query(getUsersURL)
		var users []entity.User
		json.Unmarshal(b, &users)
		fmt.Println("\n------------------------------------\nAll users' infomation:")
		for _, u := range users {
			fmt.Println("{")
			fmt.Print("\tID: ", u.Id, ",\n\tName: ", u.Name, ",\n\tEmail: ", u.Email, ",\n\tTel: ", u.Tel)
			fmt.Print("\n}\n")
		}
	},
}

func init() {
	RootCmd.AddCommand(quCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
