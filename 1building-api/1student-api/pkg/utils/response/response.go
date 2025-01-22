package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

// type Response struct {
// 	Status string
// 	Error  string
// }
//at time of retiing this above response that ,but in go captital case words not allowed so
/*
{
    "Status": "Error",
    "Error": "EOF"
}

*/
//we use tags for it so how the json is retun as

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error message"`
}

/*
then output is look like as bellow
{
    "status": "Error",
    "error message": "EOF"
}
*/

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson(res http.ResponseWriter, status int, data interface{}) error {

	res.Header().Set("Content-Type", "application/json") //here we setting the content-type
	res.WriteHeader(status)                              // here we passing what we want the status code
	return json.NewEncoder(res).Encode(data)             // encode return error in not error then it retun null
}

func GerneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	//we are lisving list of validation error by err validator.ValidationErrors
	var errMsgList []string
	for _, err := range errs {
		switch err.ActualTag() { //here we cheking due to which tag error is gating
		case "required":
			errMsgList = append(errMsgList, fmt.Sprintf("fild %s is required", err.Field()))
		default:
			errMsgList = append(errMsgList, fmt.Sprintf("Invalid field %s", fmt.Sprintf(err.Field())))
		}
	}
	//now we join all the eror list in string that we was adding in errMsg
	errMsg := strings.Join(errMsgList, ",")

	return Response{
		Status: StatusError,
		Error:  errMsg,
	}

}
