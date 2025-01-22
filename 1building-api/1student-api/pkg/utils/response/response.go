package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(res http.ResponseWriter, status int, data interface{}) error {

	res.Header().Set("Content-Type", "application/json") //here we setting the content-type
	res.WriteHeader(status)                              // here we passing what we want the status code
	return json.NewEncoder(res).Encode(data)  //
}
