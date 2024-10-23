package utils

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadDotenv(){
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env dosyası yüklenmedi.")
	}
}