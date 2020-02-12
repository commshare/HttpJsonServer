package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"html"
	"net/http"
	"time"
)
import "fmt"

//By adding struct tags you can control exactly what an how your struct will be marshalled to JSON.
//type Todo struct {
//	Name      string
//	Completed bool
//	Due       time.Time
//}
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
//Note that in the last line we create another type, called Todos, which is a slice (an ordered collection) of Todo.
type Todos []Todo
//http://localhost:8080/todos
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
	//Send Back Some JSON
	//a static slice of Todos
	/* 数组 或者叫做 slice
	[
	    {
	        "Name": "Write presentation",
	        "Completed": false,
	        "Due": "0001-01-01T00:00:00Z"
	    },
	    {
	        "Name": "Host meetup",
	        "Completed": false,
	        "Due": "0001-01-01T00:00:00Z"
	    }
	]
	*/
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}
// http://localhost:8080/todos/{todoId}
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func main() {

	//http.HandleFunc("/echo", func(w http.ResponseWriter, r * http.Request){
	//	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//
	//	print(" =====server echo===>")
	//	s := r.FormValue("data") //form data
	//	b1 := [] byte(s)
	//	w.Write(b1)
	//
	//	print(" server echo ",s)
	//	body,err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		//error
	//	} else {
	//		w.Write(body)
	//	}
	//})

	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)
	err := http.ListenAndServe(":8080", router )
	if err != nil {
		fmt.Println("http server create error: ",err.Error())
		return
	}
}
