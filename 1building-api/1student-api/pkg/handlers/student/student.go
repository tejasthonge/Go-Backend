package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/utils/response"
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/types"
)

func New() http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		var newStudent types.Student
		err := json.NewDecoder(req.Body).Decode(&newStudent) // here we decode the req.body and store(assing values) in adreess of newStuctnt struct
		//if error gotin we will write the error to resonpsse and retun

		if errors.Is(err, io.EOF) {
			response.WriteJson(res, http.StatusBadRequest, response.GerneralError(fmt.Errorf("body is empty"))) //response.GerneralError(err) this function return the struct but it will autmaticlay convet into json
			// in replase of fmt.Errorf("body is empty") we also pass the err
			return
		}

		//if above errror is not getting or geting anthor genalral error at time of Decoding the req.Body
		if err != nil {
			response.WriteJson(res, http.StatusBadRequest, err)
			return
		}
		//now we have to validate the request
		//this is callad as zero stust pollicy
		//to validate exact field are passad or not by user we use the valitaor pakage
		/*
			stape 1
				-fist get package buy ->go get github.com/go-playground/validator/v10
			stape 2
				- add the stags on our struct that you want
				like
					-
					type Student struct {
						Id    int
						Name  string `validate:"required"`
						Email string `validate:"required"`
					}
			stape3
			call the method New() on validate package like
			;// it will retun the error if gotted
		*/
		if err := validator.New().Struct(newStudent); err != nil {
			//we have to type cast the error in
			validationErros := err.(validator.ValidationErrors) //err type cast into validator.ValidationErrors by this way
			response.WriteJson(res, http.StatusBadRequest, response.ValidationError(validationErros))
			return
		}
		

		slog.Info("Creating the Student")
		response.WriteJson(res, http.StatusCreated, map[string]any{
			"success":     "OK",
			"new student": newStudent,
		})
	}
}
