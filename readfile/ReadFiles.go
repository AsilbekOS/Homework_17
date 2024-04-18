package readfile

import (
	"log"
	"os"
)

func Reader() []byte {
	data, err := os.ReadFile("employees.json")
	if err != nil {
		log.Fatal("Fileni o'qishda xatolik bor - ", err)
	}
	return data
}
