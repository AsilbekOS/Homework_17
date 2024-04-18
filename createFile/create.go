package createfile

import (
	"log"
	"os"
)

func Createfile() os.File {
	file, err := os.Create("employees_new.json")
	if err != nil {
		log.Fatal("Fayl yaratishda xatolik bor - ", err)
	}
	return *file
}
