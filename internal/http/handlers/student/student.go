package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/vedant20082004/students-api/internal/storage"
	"github.com/vedant20082004/students-api/internal/types"
	"github.com/vedant20082004/students-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a Student")

		var student types.Student


		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF){
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return 
		}

		if err != nil{
			response.WriteJson(w,http.StatusBadRequest, response.GeneralError(err))
			return 
		}


		

		// request valdation 

		if err := validator.New().Struct(student); err != nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest, response.ValidationError(validateErrs))
			return 
		}

		lastId,err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("User Created seuccessfully", slog.String("userId", fmt.Sprint(lastId)))

		if err !=nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return 
		}





		// "name":"Rakesh",
		// "email" : "r@gmail.com",
		// "age": "30",



		// w.Write([]byte("WELCOME TO STUDENTS API"))

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		
		id := r.PathValue("id")
		slog.Info("Getting A student", slog.String("id", id))


		intId,e1:= strconv.ParseInt(id, 10,64)
		if e1!=nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(e1))
			return 
		}

		student, e2 := storage.GetStudentById(intId)
		if e2 != nil{
			slog.Error("error getting user", slog.String("id",id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(e2))
			return 	
		}
		
		response.WriteJson(w, http.StatusOK, student)
		
	}

}