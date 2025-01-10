package main

import (
	"log"
	"net/http"

	"pantela/internal/database"
	"pantela/internal/handlers"
	"pantela/internal/taskServise"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskServise.Task{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
	repo := taskServise.NewRepository(database.DB)
	service := taskServise.NewService(repo)
	taskHandler := handlers.NewTaskHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PATCH")
	router.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//package main
//
//import (
//	"encoding/json"
//	"net/http"
//
//	"github.com/gorilla/mux"
//)
//
//func main() {
//	InitDB()
//	DB.AutoMigrate(&Message{})
//
//	router := mux.NewRouter()
//	router.HandleFunc("/api/messages", CreateTask).Methods("GET")
//	router.HandleFunc("/api/messages", GetAllTasks).Methods("POST")
//	router.HandleFunc("/api/messages/{id}", UpdateTaskByID).Methods("PATCH")
//	router.HandleFunc("/api/messages/{id}", DeleteTaskByID).Methods("DELETE")
//	http.ListenAndServe(":8080", router)
//}
//
//func CreateTask(w http.ResponseWriter, r *http.Request) {
//	var messages []Message
//	if err := DB.Find(&messages).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(messages)
//}
//
//func GetAllTasks(w http.ResponseWriter, r *http.Request) {
//	var message Message
//	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if err := DB.Create(&message).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(message)
//}
//
//func UpdateTaskByID(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	var message Message
//	if err := DB.First(&message, id).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//
//	var updateData map[string]interface{}
//	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if err := DB.Model(&message).Updates(updateData).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	json.NewEncoder(w).Encode(message)
//}
//
//func DeleteTaskByID(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	var message Message
//	if err := DB.First(&message, id).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusNotFound)
//		return
//	}
//
//	if err := DB.Delete(&message).Error; err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
