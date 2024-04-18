package upgrade

import (
	"employe/structs"
)

func Upgrade(employes []structs.Employe) {
	for i := 0; i < len(employes); i++ {
		if employes[i].ID == 3 {
			employes[i].Age = 50
		}
	}
}
