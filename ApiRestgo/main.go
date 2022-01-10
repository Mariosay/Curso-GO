package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"strconv"
)
type task struct{
	ID int  `json:ID`
	Name string `json:Name`
	Content string  `json:Content`
}

type allTask []task

var tasks =allTask {
	{
		ID:1,
		Name:"Task One",
		Content:"Some Content",
	},
}
func getTasks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type", "aplication/json");
  json.NewEncoder(w).Encode(tasks)
}


func createTask(w http.ResponseWriter, r *http.Request){
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil{
		fmt.Fprintf(w, "Insert a valid Task")
	}
	json.Unmarshal(reqBody,&newTask)
	newTask.ID= len(tasks)+1
	tasks=append(tasks, newTask)

	w.Header().Set("content-Type", "aplication/json");
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
func getTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	
	taskID,err :=strconv.Atoi(vars["id"])
	if err!=nil{
		fmt.Fprintf(w,"Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.ID==taskID{
			w.Header().Set("content-Type", "aplication/json");
			json.NewEncoder(w).Encode(task)
		}
	}
}
func deleteTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskID,err:=strconv.Atoi(vars["id"])
	if err!=nil{
		fmt.Fprintf(w,"Invalid ID")
		return
	}
	for i,t := range tasks{
		if t.ID == taskID{
			tasks=append(tasks[:i], tasks[i + 1:]...)
			fmt.Fprintf(w,"The task ID %v has been remove succesfuly",taskID)
		}
	}
 }

func updateTask(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r)
	taskID,err := strconv.Atoi(vars["id"])
	var UpdatedTask task
	if err != nil {
		fmt.Fprintf(w,"Invalid ID")
	}

	reqBody,err := ioutil.ReadAll(r.Body)	
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data" )
	}

	json.Unmarshal(reqBody,&UpdatedTask)

	for i, t := range tasks{
		if t.ID == taskID{
			tasks=append(tasks[:i], tasks[i + 1:]...)
			UpdatedTask.ID = taskID
			tasks=append(tasks,UpdatedTask)
			fmt.Fprintf(w, "The task with ID %v has been Updated successfully",taskID)
		}
	}

 }



func indexRoute(w http.ResponseWriter, r *http.Request){
 fmt.Fprintf(w, "welcome to my API");
}

func main() {
	router:=mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")





	
	log.Fatal(http.ListenAndServe(":3000",router))
	log.Println("Servidor corriendo..")
}

