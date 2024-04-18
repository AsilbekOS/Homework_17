package marshalindent

import (
	"employe/structs"
	"encoding/json"
	"log"
)

func MarshalIndent(employes []structs.Employe) []byte {
	encodeData, err := json.MarshalIndent(employes, "", "\t")
	if err != nil {
		log.Fatal("Json faylga yozishda xatolik bot -  ", err)
	}
	return encodeData
}
