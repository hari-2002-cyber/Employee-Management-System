
package service

import (
	"context"
	"errors"
	"net/mail"
	"strings"
	"time"

	"employee-management-system/model"
	"employee-management-system/repository"
)

var ErrInvalidEmployee = errors.New("invalid employee data")

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{Repo: repo}
}

// GetAll retrieves all employees.
func (s *EmployeeService) GetAll(ctx context.Context) ([]model.Employee, error) {
	return s.Repo.GetAll(ctx)
}

// GetByID retrieves one employee by ID.
func (s *EmployeeService) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	if id <= 0 {
		return nil, ErrInvalidEmployee
	}

	return s.Repo.GetByID(ctx, id)
}

// validateEmployee checks and normalizes employee data.
func (s *EmployeeService) validateEmployee(e *model.Employee) error {
	e.Name = strings.TrimSpace(e.Name)
	e.Email = strings.ToLower(strings.TrimSpace(e.Email))
	e.Department = strings.TrimSpace(e.Department)
	e.Designation = strings.TrimSpace(e.Designation)
	e.Status = strings.TrimSpace(e.Status)

	if e.Name == "" || e.Email == "" ||
		e.Department == "" || e.Designation == "" {
		return ErrInvalidEmployee
	}

	address, err := mail.ParseAddress(e.Email)
	if err != nil || address.Address != e.Email || !strings.Contains(e.Email, "@") {
		return ErrInvalidEmployee
	}

	if e.Status != "Active" && e.Status != "Inactive" {
		return ErrInvalidEmployee
	}

	if e.JoiningDate.IsZero() ||
		e.JoiningDate.After(time.Now().AddDate(10, 0, 0)) {
		return ErrInvalidEmployee
	}

	return nil
}

// Create validates and saves a new employee.
func (s *EmployeeService) Create(ctx context.Context, e *model.Employee) error {
	e.Status = strings.TrimSpace(e.Status)
	if e.Status == "" {
		e.Status = "Active"
	}

	if err := s.validateEmployee(e); err != nil {
		return err
	}

	return s.Repo.Create(ctx, e)
}

// Update validates and updates an existing employee.
func (s *EmployeeService) Update(ctx context.Context, e *model.Employee) error {
	if e.ID <= 0 {
		return ErrInvalidEmployee
	}

	if err := s.validateEmployee(e); err != nil {
		return err
	}

	return s.Repo.Update(ctx, e)
}

// Delete removes an employee by ID.
func (s *EmployeeService) Delete(ctx context.Context, id int) error {
	if id <= 0 {
		return ErrInvalidEmployee
	}

	return s.Repo.Delete(ctx, id)
}
