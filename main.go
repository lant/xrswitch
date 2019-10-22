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

  var optionsSize = len(commands)

	var selected int
	fmt.Print("Select Option > ")
	_, inputError := fmt.Scanf("%d", &selected)
	if inputError != nil {
		fmt.Print(inputError)
		os.Exit(-1)
	}

  if selected > optionsSize || selected <= 0 {
		fmt.Println("Invalid option")
		os.Exit(-1)
  }

  execRandr(commands[selected])
  execWallpaper()
}

func execRandr(randr []string) {
	command := exec.Command("xrandr", randr...)
	err := command.Run()
	if err != nil {
		fmt.Printf("xrandr failed with: %v",err)
	}
}

func execWallpaper() {
  command := exec.Command("nitrogen", "--set-scaled", "/home/marc/wallpapers/amps.jpg")
	err := command.Run()
	if err != nil {
		fmt.Printf("Could not set the wallpaper: %v",err)
	}
}
