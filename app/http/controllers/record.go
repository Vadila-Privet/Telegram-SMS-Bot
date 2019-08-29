package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"../../repository"
	"github.com/gorilla/mux"
)

const tmpl = `<html>
<h3>{{range .}}</h3>
<h3>Name: {{.Name}}</h3>
<h3>Phone: {{.Phone}}</h3>
<h3>Message: {{.Message}}</h3>
<h3>Error: {{.Error}}</h3>
<br>
{{end}}
</html>`

//SayHello ...
func SayHello(w http.ResponseWriter, r *http.Request) {

	list, _ := repository.GetAllRecords()

	t := template.New("Records template")

	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := t.Execute(w, list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//ReturnCertainUserRecords returns list of the certain user records
func ReturnCertainUserRecords(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	list, _ := repository.GetRecordsByName(vars["name"])

	t := template.New("Records template")

	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := t.Execute(w, list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

//ReturnCertainPhoneRecords returns list of the certain records by phone
func ReturnCertainPhoneRecords(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	list, _ := repository.GetRecordsByPhone(vars["phone"])

	t := template.New("Records template")

	t, err := t.Parse(tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := t.Execute(w, list); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
