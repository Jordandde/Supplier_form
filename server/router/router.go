package router

import (
	"github.com/gorilla/mux"
	"github.com/jordandde/supplier-form/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/supplier", middleware.GetSuppliers).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/supplier", middleware.CreateSuppliers).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteSupplier/{id}", middleware.DeleteSupplier).Methods("DELETE", "OPTIONS")

	return router
}
