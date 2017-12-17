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
	"os"

	"github.com/RowlingWu/agenda-service/entity"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the system using an identity",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		err := login(username, password, entity.Localhost+"/v1/user/login")
		checkError(err)
		log.Println("Log in success.")

		f, err := os.OpenFile(entity.CurUser, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		checkError(err)
		f.WriteString(username)
	},
}

func init() {
	RootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("user", "u", "Anonymous", "Help message for username")
	loginCmd.Flags().StringP("password", "p", "Anonymous", "Help message for password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
