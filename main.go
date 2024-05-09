package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

func getMacSysThemeHex() string {
	// Modern mac has low opacity by default on system applications.
	// This means the actual color of app windows and status bars is heavily infualced by the background (desktop picture).
	// These are the closest hard coded values I could find. Looks like no native mac API will give you this info.

	// Ignore error. If not darkmode, this errors since there is nothing to read. Meaning its light mode
	res, _ := exec.Command("defaults", "read", "-g", "AppleInterfaceStyle").Output()
	if strings.Contains(string(res), "Dark") {
		return "#373039"
	}
	return "#E8E4EB"
}

func getWinSysThemeHex() string {
	// Dont have a Windows device to test this against. LMK!
	res, err := exec.Command("cmd", "/c", "reg", "query", "HKEY_CURRENT_USER\\Software\\Microsoft\\Windows\\DWM", "/v", "ColorizationColor").Output()
	if err != nil {
		panic("Error running shell comand")
	}

	// Grab hex code with regex
	re := regexp.MustCompile(`0x([a-fA-F0-9]+)`)
	matches := re.FindStringSubmatch(string(res))

	if len(matches) > 1 {
		hexColor := matches[1]
		return hexColor
	} else {
		return "Error parsing hexcode"
	}

}

func main() {
	switch runtime.GOOS {
	case "darwin":
		color := getMacSysThemeHex()
		fmt.Println("Operating System: macOS")
		fmt.Printf("You system color theme is %v !\n", color)
	case "windows":
		color := getWinSysThemeHex()
		fmt.Println("Operating System: Windows")
		fmt.Printf("You system color theme is %v !\n", color)
	case "linux":
		fmt.Println("Operating System: Linux")
		fmt.Println("You system color theme is Linus Torvalds :-)")
	default:
		fmt.Println("What is this, an OS for ants!?")
	}
}
