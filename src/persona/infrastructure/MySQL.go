package infrastructure

import (
	"api_recu_corte1/src/core"
	"api_recu_corte1/src/persona/domain"
	"fmt"
	"log"
)

// alamacena la conexión a la bd
type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(person domain.Person) (uint, error) {
	query := "INSERT INTO persons (name, age, gender) VALUES (?, ?, ?)"
	res, err := mysql.conn.ExecutePreparedQuery(query, person.Name, person.Age, person.Gender)
	if err != nil {
		fmt.Println("Error al preparar la consulta: %v", err)
		return 0, err
	}
	id, _ := res.LastInsertId()
	fmt.Println("Persona creada")
	return uint(id), nil
}