package writefile

import (
	"fmt"
	"log"
	"os"
)

func WriteFile(file os.File, encodeData []byte) {
	_, err := file.Write(encodeData)
	if err != nil {
		_, err := file.Write(encodeData)
		if err != nil {
			log.Fatal("Faylga yozishda xatolik bor - ", err)
		}
	}
	fmt.Println("Faylga muvaffaqiyatli yozildi.")
}
