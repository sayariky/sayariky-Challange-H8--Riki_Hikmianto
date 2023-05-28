package cli

import (
	"fmt"
	"go-trial-class/helpers"
	"os"
)

func MainMenu() {
	helpers.ClearScreen()
	fmt.Println("Selamat datang di aplikasi ini")
	fmt.Println("---------------------------------")

	var input string
	fmt.Println("Tekan (1) untuk register")
	fmt.Println("Tekan (2) untuk login")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err.Error())
	}

	switch input {
	case "1":
		Register()
	case "2":
		Login("")
	case "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini")
		os.Exit(1)
	default:
		MainMenu()
	}
}
