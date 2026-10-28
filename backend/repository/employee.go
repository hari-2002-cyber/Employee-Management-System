
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"employee-management-system/model"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository(db *sql.DB) *EmployeeRepository {
	return &EmployeeRepository{DB: db}
}

// GetAll retrieves all employees.
func (r *EmployeeRepository) GetAll(ctx context.Context) ([]model.Employee, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, name, email, department, designation, joining_date, status
		FROM employees
		ORDER BY id
	`)
	if err != nil {
		return nil, fmt.Errorf("query employees: %w", err)
	}
	defer rows.Close()

	employees := []model.Employee{}

	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(
			&e.ID, &e.Name, &e.Email, &e.Department,
			&e.Designation, &e.JoiningDate, &e.Status,
		); err != nil {
			return nil, fmt.Errorf("scan employee: %w", err)
		}
		employees = append(employees, e)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate employees: %w", err)
	}

	return employees, nil
}

// GetByID retrieves one employee.
func (r *EmployeeRepository) GetByID(ctx context.Context, id int) (*model.Employee, error) {
	var e model.Employee

	err := r.DB.QueryRowContext(ctx, `
		SELECT id, name, email, department, designation, joining_date, status
		FROM employees WHERE id = $1
	`, id).Scan(
		&e.ID, &e.Name, &e.Email, &e.Department,
		&e.Designation, &e.JoiningDate, &e.Status,
	)
	if err != nil {
		return nil, fmt.Errorf("get employee by ID: %w", err)
	}

	return &e, nil
}

// Create inserts a new employee.
func (r *EmployeeRepository) Create(ctx context.Context, e *model.Employee) error {
	err := r.DB.QueryRowContext(ctx, `
		INSERT INTO employees
			(name, email, department, designation, joining_date, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`,
		e.Name, e.Email, e.Department, e.Designation,
		e.JoiningDate.Format("2006-01-02"), e.Status,
	).Scan(&e.ID)

	if err != nil {
		return fmt.Errorf("create employee: %w", err)
	}

	return nil
}

// Update modifies an existing employee.
func (r *EmployeeRepository) Update(ctx context.Context, e *model.Employee) error {
	result, err := r.DB.ExecContext(ctx, `
		UPDATE employees
		SET name = $1, email = $2, department = $3,
		    designation = $4, joining_date = $5, status = $6
		WHERE id = $7
	`,
		e.Name, e.Email, e.Department, e.Designation,
		e.JoiningDate.Format("2006-01-02"), e.Status, e.ID,
	)
	if err != nil {
		return fmt.Errorf("update employee: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check updated rows: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete removes an employee by ID.
func (r *EmployeeRepository) Delete(ctx context.Context, id int) error {
	result, err := r.DB.ExecContext(ctx,
		`DELETE FROM employees WHERE id = $1`, id,
	)
	if err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("check deleted rows: %w", err)
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
