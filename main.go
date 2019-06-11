package main

import (
	"fmt"
	"os"
	"os/exec"
)


func main() {
	commands := map[int][]string {
		1: { "--output",  "eDP-1-1", "--mode", "1600x900", "--output", "DP-1-1", "--off", "--output", "DP-1-3", "--off" },
		2: { "--output",  "eDP-1-1", "--off", "--output", "DP-1-1", "--mode", "2560x1440" },
		3: { "--output",  "eDP-1-1", "--off", "--output", "DP-1-3", "--mode", "2560x1440", "--primary" },
		4: { "--output",  "eDP-1-1", "--off", "--output", "DP-1-3", "--mode", "1920x1080" },
	}

	fmt.Println("Select display setup:")
	fmt.Println("1) Laptop")
	fmt.Println("2) Office")
	fmt.Println("3) Home")
	fmt.Println("4) Parents")
	fmt.Println("---------------------------")

	var selected int
	fmt.Print("Select Option > ")
	_, inputError := fmt.Scanf("%d", &selected)
	if inputError != nil {
		fmt.Print(inputError)
		os.Exit(-1)
	}

	switch selected {
	case 1:
		execRandr(commands[1])
	case 2:
		execRandr(commands[2])
	case 3:
		execRandr(commands[3])
	case 4:
		execRandr(commands[4])
	default:
		fmt.Println("Invalid option")
	}
}

func execRandr(randr []string) {
	command := exec.Command("xrandr", randr...)
	err := command.Run()
	if err != nil {
		fmt.Printf("xrandr failed with: %v",err)
	}
}
