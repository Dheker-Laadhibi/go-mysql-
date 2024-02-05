package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"log"
)

type Employee struct {
	ID   int
	Name string
	City string
}

var db *sql.DB
var err error

// connection
func dbConn() {
	dbDrive := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"
	db, err = sql.Open(dbDrive, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	// Vérification de la connexion à la base de données
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Définit les options de la connexion à la base de données
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	// Vérifie si la connexion à la base de données a réussi
	if err == nil {
		fmt.Println("Connexion à la base de données réussie.")
	} else {
		fmt.Println("Échec de la connexion à la base de données.")
	}
}

func getAllEmployees() []Employee {
	row, err := db.Query("select * FROM employee")
	if err != nil {
		log.Fatal(err)
	}
	// create instance of employee
	emp := Employee{}
	employees := []Employee{}
	for row.Next() {
		//scan copy the colomuns in the current row
		err := row.Scan(&emp.ID, &emp.Name, &emp.City)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}
	return employees
}

func insert(name string, city string) {
	stmt, err := db.Prepare("INSERT INTO employee(Name,City) values(?,?)")

	if err != nil {

		log.Fatal()
	}
	r, err := stmt.Exec(name, city)
	if err != nil {
		log.Fatal(err)
	}
	// number of rows affected
	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("statement affected  %d rows\n ", affectedRows)
}

func update(id int, name string, city string) {
	stmt, err := db.Prepare("UPDATE   employee SET Name=? , City=? WHERE ID=?")

	if err != nil {

		log.Fatal()
	}
	r, err := stmt.Exec(name, city, 1)
	if err != nil {
		log.Fatal(err)
	}
	// number of rows affected
	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("statement affected  %d rows\n ", affectedRows)
}

func delete(id int) {
	stmt, err := db.Prepare("DELETE FROM employee WHERE ID=?")

	if err != nil {

		log.Fatal()
	}
	r, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	// number of rows affected
	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Printf("statement affected  %d rows\n ", affectedRows)
}

func main() {
	dbConn()
	fmt.Println(getAllEmployees())
	//insert("zizou ", "algeria ")
	//update(1, "test", "america")
	//delete(2)
	fmt.Println(getAllEmployees())
}
