package router

import (
	"github.com/Keshav-Agrawal/mongoapi/controller"
	"github.com/Keshav-Agrawal/mongoapi/datasource"
	"github.com/gorilla/mux"
)

func Router(ds datasource.IDataSource) *mux.Router {
	router := mux.NewRouter()
	homeworkSVC := controller.NewHomeWorkService(ds)

	router.HandleFunc("/api/tasks", homeworkSVC.GetMyAllTask).Methods("GET")
	router.HandleFunc("/api/task", homeworkSVC.CreateTask).Methods("POST")
	router.HandleFunc("/api/task/{id}", homeworkSVC.MarkAsDone).Methods("PUT")
	router.HandleFunc("/api/task/{id}", homeworkSVC.DeleteATask).Methods("DELETE")
	router.HandleFunc("/api/task", homeworkSVC.DeleteAllTask).Methods("DELETE")

	return router
}
