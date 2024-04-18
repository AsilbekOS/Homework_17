package main

import (
	createfile "employe/createFile"
	"employe/readfile"
	"employe/readfile/marshalindent"
	"employe/readfile/upgrade"
	"employe/structs"
	"employe/userinfo"
	"employe/writefile"
	"encoding/json"
	"fmt"
	"log"
)

func main() {

	var employes []structs.Employe

	// employees.json faylini o'qish
	data := readfile.Reader()

	// Json faylni stringga o'qish
	err := json.Unmarshal(data, &employes)
	if err != nil {
		log.Fatal("employes Slicega joylashda xatolik bor - ", err)
	}

	// 3 - IDdagi personning yoshini o'zgartishish
	upgrade.Upgrade(employes)

	// Yangi "employees_new.json" faylini yaratish
	file := createfile.Createfile()

	// User tomonidan yangi ma'lumotlar kiritish
	fmt.Println("Yangi ma'lumot kitish:")
	newuser := userinfo.InputUser()

	// Yangi ma'lumotni oldingisig qo'shish
	employes = append(employes, newuser)

	// Strut ma'lumotarini jsonga yozish uchun marshal qilish
	encodeData := marshalindent.MarshalIndent(employes)
	// Yangi faylga barcha ma'lumotlarni yozis
	writefile.WriteFile(file, encodeData)

}
