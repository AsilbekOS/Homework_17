package userinfo

import (
	"bufio"
	"employe/structs"
	"fmt"
	"os"
)

func InputUser() structs.Employe {
	var users structs.Employe

	var (
		id, age        int
		name, position string
	)
	id = 6
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Ismingizni kiriting: ")
	scan.Scan()
	name = scan.Text()
	fmt.Print("Yoshingizni kiriting: ")
	fmt.Scan(&age)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Lavozimingizni kiriting: ")
	scanner.Scan()
	position = scanner.Text()

	users.ID = id
	users.Name = name
	users.Age = age
	users.Position = position

	return users
}
