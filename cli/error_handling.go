package cli

import (
	"fmt"
	"os"
)

func ErrorHandler(message string) {
	fmt.Println("Terjadi kesalahan dalam aplikasi")
	fmt.Println("Error:", message)

	var input string
	fmt.Println("Tekan (m) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err.Error())
	}

	switch input {
	case "l":
		MainMenu()
	case "q":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini")
		// 1 is true, 0 is false
		os.Exit(1)
	default:
		ErrorHandler(message)
	}
}
