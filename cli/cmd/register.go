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
	"log"
	"github.com/RowlingWu/agenda/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "to build an account",
	Long: `build an account, include username, password, email, telephone number`,
	Run: func(cmd *cobra.Command, args []string) {
		name,_:= cmd.Flags().GetString("username")
		pw, _ := cmd.Flags().GetString("password")
		em, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		pass, err := entity.MyRegister(name, pw, em, phone)
		if pass ==false {
			log.Fatal("username has been registered, please choose another one")
		} else {
			if err != nil {
				log.Fatal("some unexpected errors occur")
			} else {
				log.Fatal("register successfully")
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "username")
	registerCmd.Flags().StringP("password", "p", "", "password")
	registerCmd.Flags().StringP("email", "e", "", "user's email")
	registerCmd.Flags().StringP("phone", "t", "", "user's phone number")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
