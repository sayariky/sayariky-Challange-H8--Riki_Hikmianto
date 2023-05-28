package cli

import (
	"bufio"
	"fmt"
	"go-trial-class/config"
	"go-trial-class/entity"
	"go-trial-class/helpers"
	"os"
)

func ListOrder() {
	helpers.ClearScreen()
	consoleReader := bufio.NewReader(os.Stdin)
	var order []entity.Order

	err := config.DB.Preload("Product").Find(&order).Error

	if err != nil {
		ErrorHandler(err.Error())
		return
	}

	fmt.Println("---List Order---")
	for _, order := range order {
		order.PrintDetail()
	}

	var input string
	fmt.Println("Tekan (any key) untuk kembali ke menu utama")
	fmt.Println("Tekan (q) untuk keluar dari aplikasi")

	input, err = consoleReader.ReadString('\n')

	switch input {
	case "q\n":
		fmt.Println("Terima kasih telah menggunakan aplikasi ini")
		// 1 is true, 0 is false
		os.Exit(1)
	default:
		MainMenu()
	}
}
