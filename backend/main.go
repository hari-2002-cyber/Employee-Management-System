
package main

import (
	"log"
	"net/http"

	"employee-management-system/config"
	"employee-management-system/handler"
	"employee-management-system/repository"
	"employee-management-system/service"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	h := handler.NewEmployeeHandler(svc)

	mux := http.NewServeMux()

	// View all employees
	mux.HandleFunc("GET /api/employees", h.GetAll)

	// View one employee
	mux.HandleFunc("GET /api/employees/{id}", h.GetByID)

	// Add an employee
	mux.HandleFunc("POST /api/employees", h.Create)

	// Update an employee
	mux.HandleFunc("PUT /api/employees/{id}", h.Update)

	// Delete an employee
	mux.HandleFunc("DELETE /api/employees/{id}", h.Delete)

	log.Println("Employee Management System API running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", enableCORS(mux)))
}
