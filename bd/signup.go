package bd

import (
	"fmt"

	"github.com/Rolas444/gambituser/models"
	"github.com/Rolas444/gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES('" + sig.UserEmail + "', '" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	fmt.Println("SignUp > Ejecución exitosa")
	// _, err =
	Db.Exec(sentencia)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }

	return nil
}
