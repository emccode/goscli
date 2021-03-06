package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v1"
)

var userCmdV *cobra.Command

func init() {
	addCommandsUser()
	// userCmd.Flags().StringVar(&username, "username", "", "GOSCALEIO_TEMP")
	userCmd.Flags().StringVar(&systemid, "systemid", "", "GOSCALEIO_SYSTEMID")
	usergetCmd.Flags().StringVar(&systemid, "systemid", "", "GOSCALEIO_SYSTEMID")

	userCmdV = userCmd

	// initConfig(userCmd, "goscli", true, map[string]FlagValue{
	// 	"endpoint": {endpoint, true, false, ""},
	// 	"insecure": {insecure, false, false, ""},
	// })

	userCmd.Run = func(cmd *cobra.Command, args []string) {
		setGobValues(cmd, "goscli", "")
		cmd.Usage()
	}
}

func addCommandsUser() {
	userCmd.AddCommand(usergetCmd)
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user",
	Long:  `user`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var usergetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user",
	Long:  `Get a user`,
	Run:   cmdGetUser,
}

func cmdGetUser(cmd *cobra.Command, args []string) {
	client, err := authenticate()
	if err != nil {
		log.Fatalf("error authenticating: %v", err)
	}

	initConfig(cmd, "goscli", true, map[string]FlagValue{
		"systemhref": {&systemhref, true, false, ""},
	})

	systemhref = viper.GetString("systemhref")

	system, err := client.FindSystem("", systemhref)
	if err != nil {
		log.Fatalf("err: problem getting system: %v", err)
	}

	users, err := system.GetUser()
	if err != nil {
		log.Fatalf("error getting statistics: %v", err)
	}

	yamlOutput, err := yaml.Marshal(&users)
	if err != nil {
		log.Fatalf("error marshaling: %s", err)
	}
	fmt.Println(string(yamlOutput))

}
