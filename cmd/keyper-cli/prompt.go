package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/manifoldco/promptui"
)

func promptInput(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		os.Exit(1)
	}
	return result
}

func promptInputOptional(label string) string {
	prompt := promptui.Prompt{
		Label:     label,
		Default:   "",
		AllowEdit: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return ""
	}
	return result
}

func promptInt(label string, defaultVal int) int {
	prompt := promptui.Prompt{
		Label:   label,
		Default: fmt.Sprintf("%d", defaultVal),
		Validate: func(input string) error {
			_, err := strconv.Atoi(input)
			return err
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		os.Exit(1)
	}

	port, _ := strconv.Atoi(result)
	return port
}

func promptBool(label string, defaultVal bool) bool {
	prompt := promptui.Prompt{
		Label:   label,
		Default: fmt.Sprintf("%t", defaultVal),
		Validate: func(input string) error {
			_, err := strconv.ParseBool(input)
			return err
		},
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		os.Exit(1)
	}

	b, err := strconv.ParseBool(result)
	if err != nil {
		fmt.Printf("failed to parse bool %s", result)
		os.Exit(1)
	}
	return b
}
