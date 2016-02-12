package main

import(
	"fmt"
	"net/http"
	"html/template"
	"empdb"
)

func home(w http.ResponseWriter, r *http.Request) {
	t , err := template.ParseFiles("emp.html")
	if err != nil{
		fmt.Println(err)	
	}
	t.Execute(w, nil)
}

func save(w http.ResponseWriter, r *http.Request){
	var emp empdb.EmpData
	emp.FirstName = r.FormValue("firstname")
	emp.LastName = r.FormValue("lastname")
	emp.EmpId = r.FormValue("empid")
	emp.Phone = r.FormValue("phone")

	if emp.Dbcommit() {
		fmt.Fprintf(w, "Hello %s, your data is saved ", r.FormValue("firstname"))
	}
}

func main(){
	http.HandleFunc("/" , home)
	http.HandleFunc("/save", save)
	http.ListenAndServe(":8001" , nil)
}
