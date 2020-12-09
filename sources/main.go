package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func usage() {
	format := `Glitter Git User Management Cli

Developed by Timo Sarkar: https://github.com/timo-cmd
Latest build version: 0.1.2dev

For bug issueing please visit: https://github.com/timo-cmd/glitter
`
	fmt.Fprintln(os.Stderr, format)
}

var (
	isGlobal = flag.Bool("global", false, "Define user as global Git user")

	// these are set in build step
	version = "no version"
	commit  = "no commit"
	date    = "no date"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	os.Exit(run())
}

const (
	sel = "Select existing Git user"
	add = "Add new Git user"
	del = "Delete existing Git user"
)

func run() int {
	action := promptui.Select{
		Label: "Select action",
		Items: []string{sel, add, del},
	}

	_, actionType, err := action.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to select action: %v\n", err)
		return 1
	}

	switch actionType {
	case sel:
		if err := selectUser(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to select existing user: %v\n", err)
			return 1
		}
	case add:
		if err := addUser(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to add new user: %v\n", err)
			return 1
		}
	case del:
		if err := deleteUser(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to delete existing user: %v\n", err)
			return 1
		}
	default:
		fmt.Fprintf(os.Stderr, "Non-existing action type\n")
		return 1
	}

	return 0
}

func selectUser() error {
	users, err := ListUser()
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("No users selected")
		return nil
	}

	user := promptui.Select{
		Label: "Please select an existing Git user",
		Items: UsersToString(users),
	}
	selectedUserIndex, _, err := user.Run()
	if err != nil {
		return err
	}

	option := "--local"
	if *isGlobal {
		option = "--global"
	}

	cmdName := exec.Command("git", "config", option, "user.name", users[selectedUserIndex].Name)
	if err := cmdName.Run(); err != nil {
		return err
	}
	cmdMail := exec.Command("git", "config", option, "user.email", users[selectedUserIndex].Email)
	if err := cmdMail.Run(); err != nil {
		return err
	}

	return nil
}

func addUser() error {
	name := promptui.Prompt{
		Label: "Please input the Git user name",
	}
	resultName, err := name.Run()
	if err != nil {
		return err
	}

	email := promptui.Prompt{
		Label:    "Please input the Git email address",
		Validate: ValidateEmail,
	}
	resultEmail, err := email.Run()
	if err != nil {
		return err
	}

	if err := CreateUser(resultName, resultEmail); err != nil {
		return err
	}

	return nil
}

func deleteUser() error {
	users, err := ListUser()
	if err != nil {
		return err
	}

	if len(users) == 0 {
		fmt.Println("No users selected")
		return nil
	}

	user := promptui.Select{
		Label: "Please select a Git user",
		Items: UsersToString(users),
	}
	selectedUserIndex, _, err := user.Run()
	if err != nil {
		return err
	}

	if err := RemoveUser(selectedUserIndex, users); err != nil {
		return err
	}

	return nil
}
