package sqlite

import (
	"database/sql"
	"fmt"
	// "log/slog"

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


func(s *Sqlite) GetStudents() ([]types.Student,error){
	statement,err := s.Db.Prepare("SELECT * FROM students")

	if err!=nil{
		return nil , err
	}

	defer statement.Close();

	rows,err := statement.Query()
	if err!= nil{
		return nil, err
	}

	defer rows.Close()

	var students []types.Student

	for rows.Next(){
		var student types.Student

		err := rows.Scan(&student.Id,&student.Name,&student.Email,&student.Age)
		if err!= nil{
			return nil, err
		}

		students = append(students, student)
	}

	return students,nil
}

func(s *Sqlite) UpdateStudent(id int64, name string, email string, age int) (types.Student, error){

	statement,err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")

	if err!=nil{
		return types.Student{},err
	}

	defer statement.Close()

	_, err = statement.Exec(name, email, age, id)
	if err != nil{
		return types.Student{},err
	}

	return s.GetStudentById(id)
}

func(s *Sqlite) DeleteStudent(id int64) (string, error) {
	
	statement, err := s.Db.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		return "ERROR ENCOUNTERED", err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return "ERROR ENCOUNTERED", err
	}

	// slog.Info("DELETED SUCCESSFULLY")

	return "DELETED SUCCESSFULLY",nil
}


