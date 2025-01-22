package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/pkg/utils/response"
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/types"
	"github.com/tejasthonge/Go-Backend/1building-api/1student-api/types/types"
)

func New() http.HandlerFunc {

	return func(res http.ResponseWriter, req *http.Request) {
		var newStudent types.Student
		err := json.NewDecoder(req.Body).Decode(&newStudent) // here we decode the req.body and store in adreess of newStuctnt struct

		if errors.Is(err, io.EOF) {
			response.WriteJson(res, http.StatusBadRequest, err.Error())
			return
		}

		slog.Info("Creating the Student")
		response.WriteJson(res, http.StatusCreated, map[string]string{
			"success": "OK",
		})
	}
}
