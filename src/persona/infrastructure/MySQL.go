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
	lastCount int
	
}

func NewMySQL() *MySQL {
	var count int
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn, lastCount: count}
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

func (sql *MySQL)GetnewPersonIsAdded() (bool, error){

	var count int
	err := sql.conn.DB.QueryRow("SELECT COUNT(*) FROM persons").Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error obteniendo el conteo de personas: %v", err)
	}

	if count > sql.lastCount {
		sql.lastCount = count
		return true, nil
	}

	return false, nil
}


func (sql *MySQL) CountGender(sexo bool) (int, error) {
	var count int
	// Usamos el sexo (true para hombres, false para mujeres) como filtro
	err := sql.conn.DB.QueryRow("SELECT COUNT(*) FROM personas WHERE sexo = ?", sexo).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error obteniendo el conteo de personas de sexo %v: %v", sexo, err)
	}
	return count, nil
}