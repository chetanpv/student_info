package database

import (
	"database/sql"
	"fmt"
	"studentDetails/gomodule/models"

	_ "github.com/lib/pq"
)

var (
	name         string
	phone_number int
	email_id     string
	branch       string
	student_id   string
)

type Student struct {
	Id          string `json:"student_id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	EmailId     string `json:"email_id"`
	Branch      string `json:"branch"`
}

// var Dbconn *sql.DB

// func ConnectToDB() {
// 	err := errors.New("")
// 	connStr := "user=postgres dbname=student password=postgres host=127.0.0.1 sslmode=disable port=5433"
// 	Dbconn, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer Dbconn.Close()

// 	err = Dbconn.Ping()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("\nSuccessfully connected to database!\n")
// }

func GetAllFromTable() []models.StudentDetail {
	// ---------------------------------------------------------------------------------------------------------------
	connStr := "user=postgres dbname=student password=postgres host=127.0.0.1 sslmode=disable port=5433"
	Dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer Dbconn.Close()

	err = Dbconn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	// ---------------------------------------------------------------------------------------------------------------
	rows, err := Dbconn.Query("SELECT name, branch, phone_number, email_id, student_id FROM students")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var results = []models.StudentDetail{}
	for rows.Next() {
		err := rows.Scan(&name, &branch, &phone_number, &email_id, &student_id)
		if err != nil {
			panic(err)
		}
		result := models.StudentDetail{Name: name, Branch: branch, Email: email_id, PhoneNumber: phone_number, ID: student_id}
		results = append(results, result)
	}
	fmt.Println("My result from DB: ", results)
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return results
}

func InsertIntoTable(studentdetails []models.StudentDetail) {
	// ---------------------------------------------------------------------------------------------------------------
	connStr := "user=postgres dbname=student password=postgres host=127.0.0.1 sslmode=disable port=5433"
	Dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer Dbconn.Close()

	err = Dbconn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	// ---------------------------------------------------------------------------------------------------------------
	for _, value := range studentdetails {
		sqlStatement := `INSERT INTO students (name, branch, phone_number, email_id, student_id) 
    	VALUES ($1, $2, $3, $4, $5)`
		_, err = Dbconn.Exec(sqlStatement, value.Name, value.Branch, value.PhoneNumber, value.Email, value.ID)
		if err != nil {
			panic(err)
		}
	}

}

func UpdateTable() {
	// ---------------------------------------------------------------------------------------------------------------
	connStr := "user=postgres dbname=student password=postgres host=127.0.0.1 sslmode=disable port=5433"
	Dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer Dbconn.Close()

	err = Dbconn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	// ---------------------------------------------------------------------------------------------------------------
	sqlStatementUpdt := `
	UPDATE students
	SET phone_number = 3001 WHERE name = $1;`
	res, err := Dbconn.Exec(sqlStatementUpdt, "Chetan")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows updated: %v\n", count)
}

func DeleteRowInTable() {
	// ---------------------------------------------------------------------------------------------------------------
	connStr := "user=postgres dbname=student password=postgres host=127.0.0.1 sslmode=disable port=5433"
	Dbconn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer Dbconn.Close()

	err = Dbconn.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	// ---------------------------------------------------------------------------------------------------------------
	sqlStatementDel := `
	DELETE FROM students
	WHERE name = $1;`
	res1, err := Dbconn.Exec(sqlStatementDel, "Chets")
	if err != nil {
		panic(err)
	}
	count, err := res1.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("rows deleted: %v\n", count)
}
