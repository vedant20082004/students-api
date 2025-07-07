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

		response.WriteJson(w, http.StatusCreated, map[string]int64{"id": lastId})
	}
}

func GetById(storage storage.Storage) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		
		id := r.PathValue("id")
		slog.Info("Getting A student", slog.String("id", id))


		// STEP 2
		intId,e:= strconv.ParseInt(id, 10,64)
		if e!=nil{
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(e))
			return 
		}

		// STEP 1
		student, err := storage.GetStudentById(intId)
		if err != nil{
			slog.Error("error getting user", slog.String("id",id))
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return 	
		}
		
		response.WriteJson(w, http.StatusOK, student)
		
	}

}

func GetList(storage storage.Storage) http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request) {
		
		slog.Info("GETTING ALL STUDENTS")

		students,err := storage.GetStudents()
		if err!=nil{
			response.WriteJson(w,http.StatusInternalServerError,err)
			return 
		}

		response.WriteJson(w,http.StatusOK,students)
	}

}

func UpdateList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("UPDATING STUDENT")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student); 
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}

		updatedStudent, err := storage.UpdateStudent(student.Id, student.Name, student.Email, student.Age)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, updatedStudent)
	}
}


func DeleteList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("DELETING STUDENT")

		idStr := r.PathValue("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid id: %w", err)))
			return
		}

		deletedStudent, err := storage.DeleteStudent(id)
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, deletedStudent)
	}
}