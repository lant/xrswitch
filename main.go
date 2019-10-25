package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type tomlConfig struct {
	Profiles map[string]profile
}

type profile struct {
	Name    string
	Command string
}

var configDirName = "xrswitch"

func main() {

	var config tomlConfig
	var configDir, error = GetDefaultConfigDir()

	if error != nil {
		fmt.Print(error)
		os.Exit(-1)
	}

	var configFile = filepath.Join(configDir, "xrswitch.toml")
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Select display setup:")
	for profile := range config.Profiles {
		fmt.Println(profile)
	}
	fmt.Println("---------------------------")

	var selected string
	fmt.Print("Select Option > ")
	_, inputError := fmt.Scanf("%s", &selected)
	if inputError != nil {
		fmt.Print(inputError)
		os.Exit(-1)
	}

	if _, ok := config.Profiles[selected]; !ok {
		fmt.Println("Invalid option")
		os.Exit(-1)
	}

	execRandr(config.Profiles[selected].Command)
	execWallpaper()
}

func GetDefaultConfigDir() (string, error) {
	var configDirLocation string

	homeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}

	// Use the XDG_CONFIG_HOME variable if it is set, otherwise
	// $HOME/.config/example
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		configDirLocation = xdgConfigHome
	} else {
		configDirLocation = filepath.Join(homeDir, ".config", configDirName)
	}

	return configDirLocation, nil
}

func execRandr(randr string) {
	command := exec.Command("xrandr", strings.Fields(randr)...)
	err := command.Run()
	if err != nil {
		fmt.Printf("xrandr failed with: %v", err)
	}
}

func execWallpaper() {
	command := exec.Command("nitrogen", "--set-scaled", "/home/marc/wallpapers/amps.jpg")
	err := command.Run()
	if err != nil {
		fmt.Printf("Could not set the wallpaper: %v", err)
	}
}
