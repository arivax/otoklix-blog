package common

import (
	"fmt"
)

// ShowBanner ...
func ShowBanner() {
	fmt.Print("\033[H\033[2J")
	fmt.Println(" ")

	fmt.Println(" ______ _             ")
	fmt.Println(" | ___ \\ |            ")
	fmt.Println(" | |_/ / | ___   __ _ ")
	fmt.Println(" | ___ \\ |/ _ \\ / _` |")
	fmt.Println(" | |_/ / | (_) | (_| |")
	fmt.Println(" \\____/|_|\\___/ \\__, |")
	fmt.Println("                 __/ |")
	fmt.Println("                |___/ ")

	fmt.Println("\n Ver.2021.12 Created by Ariawan Satia Graha")
	//fmt.Println(" ")
}
