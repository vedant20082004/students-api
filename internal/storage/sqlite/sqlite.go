package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/vedant20082004/students-api/internal/config"
	"github.com/vedant20082004/students-api/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {

	db ,err := sql.Open("sqlite3", cfg.StoragePath)

	if err!=nil{
		return nil, err
	}

	_, err =db.Exec(`CREATE TABLE IF NOT EXISTS students (

		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT,
		age INTEGER
 
	)` )

	if err != nil{
		return nil, err
	}

	return &Sqlite{
		Db:db,
	},nil

}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error){

	statement,err := s.Db.Prepare("INSERT INTO STUDENTS (name,email,age) VALUES(?,?,?)")

	if err != nil{
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(name, email, age)
	if err != nil{
		return 0, err
	}

	lastId,err :=	result.LastInsertId()
	if err != nil{
		return 0, err
	} 

	return lastId, nil



	return 0, nil

}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error){

	statement, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	
	if err!=nil{
		return  types.Student{} ,err
	}

	defer statement.Close()

	var student types.Student

	if err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age); err!=nil{
		if err==sql.ErrNoRows {
			return types.Student{},fmt.Errorf("no student found with id %s",fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("QUERY ERROR : %w", err)
	}

	return student,nil

}