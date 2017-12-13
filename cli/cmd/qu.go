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
	"fmt"
	"github.com/RowlingWu/agenda/entity"
	"github.com/spf13/cobra"
)

// quCmd represents the qu command
var quCmd = &cobra.Command{
	Use:   "qu",
	Short: "to find user infomation ",
	Long: `All user's information.`,
	Run: func(cmd *cobra.Command, args []string) {
		                                                                                                                                                                                                            
		nuser := entity.GetAllUsers();
		for _,i:= range nuser {
				fmt.Println("----------------")
					fmt.Println("Username: ", i.Name)
					fmt.Println("Telephone number: ", i.Tel)
					fmt.Println("Email: ", i.Email)
					fmt.Println("----------------")
				}
	},
}

func init() {
	RootCmd.AddCommand(quCmd)
	quCmd.Flags().StringP("userinfo", "q", "", "user info display")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
