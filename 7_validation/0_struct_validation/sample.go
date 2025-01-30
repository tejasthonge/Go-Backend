package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

/*
	Validation:
		- that data is comming forn clint side is is
		- present in req
		- but actule data is inside the body
		- but we dont know what types of data is comming form the clint side
		- so before the storing the data in database we must have validete all
		- the fild comming form the clint it may be wrong so
		- .. we are all the body form json formate comming  form teh server but in go
		-body type is  io.ReadCloser struct
		- we fuct convet or write in or stuct
		- this is similar as seirlaztion or deserialzation
		- in other oop language like dart ,java ,use convert by using formJson method
		- and this method returns the Our class object that have the varibles with json filds values
		-but in go we use convet the io.ReadCloser to stuct we
		have the
			json.NewDecoder(body).Decode(&userData)
				 body -> is the geted body form the client present in r.Body
				useData -> this is address of the ourt struct type varible where we have to copy the value
						- but for copping in same we have to pass the address of the varible

		- for that we user  the validator package
		go get github.com/go-playground/validator/v10

		validator :-
			by using
*/

type User struct {
	FirstName      string  `validate:"required"`
	LastName       string  `validate:"required"`
	Age            int     `validate:"gte=0,lte=130,required"`
	Email          string  `validate:"required,email"`
	Gender         string  `validate:"oneof=male female prefer_not_to"`
	FavouriteColor string  `validate:"iscolor"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      Address `validate:"required,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func main() {

	router := http.NewServeMux()

	router.HandleFunc("POST /api/user", func(w http.ResponseWriter, r *http.Request) {
		var userData User
		body := r.Body //this body is io.ReadCloser type

		err := json.NewDecoder(body).Decode(&userData)

		if errors.Is(err, io.EOF) { //this error come when the body is complitly empty
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("body is require"))
			return
		}
		fmt.Println(userData)

		// now till that we coppy all the data foming form the  server and store inside the
		//now we have to validate this eacht filds by adding the fackage validtor
		// run -> go get github.com/go-playground/validator/v10
		// create the instance of

		//create the validator varible
		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(userData) //if any error we get it the it well returnthe error
		if err != nil {
			var erroMsgList []string

			for _, err := range err.(validator.ValidationErrors) {
				// fmt.Println(err.Namespace())
				// fmt.Println(err.Field())
				// fmt.Println(err.StructNamespace())
				// fmt.Println(err.StructField())
				// fmt.Println(err.Tag())
				// fmt.Println(err.ActualTag())
				// fmt.Println(err.Kind())
				// fmt.Println(err.Type())
				// fmt.Println(err.Value())
				// fmt.Println(err.Param())
				// fmt.Println()
				switch err.ActualTag() {
				case "required":
					erroMsgList = append(erroMsgList, fmt.Sprintf("field %s is required", err.Field()))
				default:
					erroMsgList = append(erroMsgList, fmt.Sprintf("field %s is Invalid", err.Field()))
				}
			}
			erMsg := strings.Join(erroMsgList, " , ")
			resposne := struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{
				Status:  "faild",
				Message: erMsg,
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resposne)
			return

		}
		//saving the data in database logic
		resposne := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{
			Status:  "success",
			Message: "user data succesfully stored",
		}
		w.Header().Set("Content-Type", "application/json") // this is call fist other wise it not work
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resposne)
	})

	server := http.Server{
		Addr:    ":8055",
		Handler: router,
	}
	fmt.Println("Starting the server ...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Errror for Starting the server...", err)
	}
}

/*  //wrong json
		{
    "FirstName":      "Tejas",
	"LastName":       "Thonge",
	"Age":            23,
	"Gender":         "male",
	"Email":          "amrthonge.com",
	"FavouriteColor": "#000-",
    "Addresses": {
		"Street": "Fc road",
		"Planet": "Earth",
		"Phone":  "+917558667986"
	}
}
*/

/*
//right
{
    "FirstName":      "Tejas",
	"LastName":       "Thonge",
	"Age":            23,
	"Gender":         "male",
	"Email":          "amrthonge@gmail.com",
	"FavouriteColor": "#000000",
    "Addresses": {
		"Street": "Fc road",
         "City": "Pune",
		"Planet": "Earth",
		"Phone":  "+917558667986"
	}
}
*/
