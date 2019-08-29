package http

import (
	"github.com/gorilla/mux"

	c "../http/controllers"
)

//New creates handlers for controllers
func New(r *mux.Router) error {

	r.HandleFunc("/", c.SayHello)
	r.HandleFunc("/user/{name}", c.ReturnCertainUserRecords)
	r.HandleFunc("/phone/{phone}", c.ReturnCertainPhoneRecords)

	return nil
}
