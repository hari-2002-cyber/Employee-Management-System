
package handler

import (
    "database/sql"
    "encoding/json"
    "errors"
    "net/http"
	"io"
    "strconv"
    "strings"
    "time"

    "employee-management-system/model"
    "employee-management-system/service"
)

type EmployeeHandler struct {
	Service *service.EmployeeService
}

func NewEmployeeHandler(s *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{Service: s}
}

func (h *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var input struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Department  string `json:"department"`
		Designation string `json:"designation"`
		JoiningDate string `json:"joining_date"`
		Status      string `json:"status"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		http.Error(w, "Request must contain one JSON object", http.StatusBadRequest)
		return
	}

	joiningDate, err := time.Parse("2006-01-02", strings.TrimSpace(input.JoiningDate))
	if err != nil {
		http.Error(w, "joining_date must use YYYY-MM-DD format", http.StatusBadRequest)
		return
	}

	employee := model.Employee{
		Name:        input.Name,
		Email:       input.Email,
		Department:  input.Department,
		Designation: input.Designation,
		JoiningDate: joiningDate,
		Status:      input.Status,
	}

	if err := h.Service.Create(r.Context(), &employee); err != nil {
		if errors.Is(err, service.ErrInvalidEmployee) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id <= 0 {
        http.Error(w, "Invalid employee ID", http.StatusBadRequest)
        return
    }

    employee, err := h.Service.GetByID(r.Context(), id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            http.Error(w, "Employee not found", http.StatusNotFound)
            return
        }

        http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(employee)
}

func (h *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var input struct {
		Name        string `json:"name"`
		Email       string `json:"email"`
		Department  string `json:"department"`
		Designation string `json:"designation"`
		JoiningDate string `json:"joining_date"`
		Status      string `json:"status"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&input); err != nil {
		http.Error(w, "Invalid JSON request body", http.StatusBadRequest)
		return
	}

	var extra any
	if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
		http.Error(w, "Request must contain one JSON object", http.StatusBadRequest)
		return
	}

	joiningDate, err := time.Parse("2006-01-02", strings.TrimSpace(input.JoiningDate))
	if err != nil {
		http.Error(w, "joining_date must use YYYY-MM-DD format", http.StatusBadRequest)
		return
	}

	employee := &model.Employee{
		ID:          id,
		Name:        input.Name,
		Email:       input.Email,
		Department:  input.Department,
		Designation: input.Designation,
		JoiningDate: joiningDate,
		Status:      input.Status,
	}

	if err := h.Service.Update(r.Context(), employee); err != nil {
		if errors.Is(err, service.ErrInvalidEmployee) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(employee)
}


func (h *EmployeeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id <= 0 {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	if err := h.Service.Delete(r.Context(), id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
