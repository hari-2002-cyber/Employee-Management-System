
import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule, DatePipe } from '@angular/common';
import { FormsModule } from '@angular/forms';
import {
  HttpClient,
  HttpClientModule,
  HttpErrorResponse
} from '@angular/common/http';

interface Employee {
  id?: number;
  name: string;
  email: string;
  department: string;
  designation: string;
  joining_date: string;
  status: 'Active' | 'Inactive';
}

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, FormsModule, HttpClientModule, DatePipe],
  templateUrl: './app.html',
  styleUrl: './app.css'
})
export class App implements OnInit {
  employees: Employee[] = [];

  // API loading and error states
  loading = false;
  errorMessage = '';
  formErrorMessage = '';
  deleteErrorMessage = '';

  searchTerm = '';

  // Department and status filters
  departmentFilter = '';
  statusFilter = '';

  // Add/Edit employee form
  showForm = false;
  editingId: number | null = null;
  saving = false;

  // Custom delete confirmation dialog
  showDeleteWarning = false;
  employeeToDelete: Employee | null = null;
  deleting = false;

  form: Employee = this.emptyForm();

  private readonly apiUrl =
    'http://localhost:8080/api/employees';

  constructor(
    private readonly http: HttpClient,
    private readonly cdr: ChangeDetectorRef
  ) {}

  ngOnInit(): void {
    this.loadEmployees();
  }

  emptyForm(): Employee {
    return {
      name: '',
      email: '',
      department: '',
      designation: '',
      joining_date: '',
      status: 'Active'
    };
  }

  // Convert API errors into readable messages
  private getApiErrorMessage(
    error: HttpErrorResponse,
    operation: string
  ): string {
    if (error.status === 0) {
      return 'Cannot connect to the server. Please check that the Go backend is running on port 8080.';
    }

    if (error.status === 400) {
      return `Invalid employee information. Please check the entered details and try again.`;
    }

    if (error.status === 401 || error.status === 403) {
      return `You are not authorized to ${operation}.`;
    }

    if (error.status === 404) {
      return `The requested employee or API endpoint was not found.`;
    }

    if (error.status === 409) {
      return 'This employee information conflicts with an existing record. Check whether the email address is already registered.';
    }

    if (error.status >= 500) {
      return 'The server encountered an error. Please try again later.';
    }

    const serverMessage = error.error;

    if (
      typeof serverMessage === 'string' &&
      serverMessage.trim() !== ''
    ) {
      return serverMessage;
    }

    if (
      serverMessage &&
      typeof serverMessage.message === 'string'
    ) {
      return serverMessage.message;
    }

    return `Unable to ${operation}. Please try again.`;
  }

  // Automatically generate department filter options
  get departments(): string[] {
    return [
      ...new Set(
        this.employees
          .map(employee => employee.department.trim())
          .filter(department => department !== '')
      )
    ].sort((a, b) => a.localeCompare(b));
  }

  // Search, department and status filters
  get filteredEmployees(): Employee[] {
    const term = this.searchTerm.trim().toLowerCase();

    return this.employees.filter(employee => {
      const matchesSearch =
        !term ||
        [
          employee.name,
          employee.email,
          employee.department,
          employee.designation
        ].some(value =>
          (value ?? '').toLowerCase().includes(term)
        );

      const matchesDepartment =
        !this.departmentFilter ||
        employee.department.toLowerCase() ===
          this.departmentFilter.toLowerCase();

      const matchesStatus =
        !this.statusFilter ||
        employee.status === this.statusFilter;

      return (
        matchesSearch &&
        matchesDepartment &&
        matchesStatus
      );
    });
  }

  // Dashboard statistics
  get activeCount(): number {
    return this.employees.filter(
      employee => employee.status === 'Active'
    ).length;
  }

  get inactiveCount(): number {
    return this.employees.filter(
      employee => employee.status === 'Inactive'
    ).length;
  }

  // Clear search and filters
  clearFilters(): void {
    this.searchTerm = '';
    this.departmentFilter = '';
    this.statusFilter = '';
  }

  // Load employees from the Go API
  loadEmployees(): void {
    if (this.loading) {
      return;
    }

    this.loading = true;
    this.errorMessage = '';
    this.cdr.markForCheck();

    this.http.get<Employee[]>(this.apiUrl).subscribe({
      next: data => {
        this.employees = (data ?? []).map(employee => ({
          ...employee,
          joining_date: employee.joining_date
            ? employee.joining_date.substring(0, 10)
            : ''
        }));
      },

      error: (error: HttpErrorResponse) => {
        console.error('Failed to load employees:', error);

        this.errorMessage = this.getApiErrorMessage(
          error,
          'load employees'
        );

        this.loading = false;
        this.cdr.markForCheck();
      },

      complete: () => {
        this.loading = false;
        this.cdr.markForCheck();
      }
    });
  }

  // Open Add Employee form
  openAddForm(): void {
    if (this.saving || this.deleting) {
      return;
    }

    this.editingId = null;
    this.form = this.emptyForm();
    this.formErrorMessage = '';
    this.showForm = true;
  }

  // Open Edit Employee form
  openEditForm(employee: Employee): void {
    if (
      employee.id === undefined ||
      this.saving ||
      this.deleting ||
      this.loading
    ) {
      return;
    }

    this.editingId = employee.id;

    this.form = {
      ...employee,
      joining_date: employee.joining_date
        ? employee.joining_date.substring(0, 10)
        : ''
    };

    this.formErrorMessage = '';
    this.showForm = true;
  }

  // Close Add/Edit form
  closeForm(): void {
    if (this.saving) {
      return;
    }

    this.showForm = false;
    this.editingId = null;
    this.form = this.emptyForm();
    this.formErrorMessage = '';
  }

  // Add or update employee
  saveEmployee(): void {
    if (this.saving || this.deleting) {
      return;
    }

    this.saving = true;
    this.formErrorMessage = '';
    this.errorMessage = '';

    const payload = {
      name: this.form.name.trim(),
      email: this.form.email.trim(),
      department: this.form.department.trim(),
      designation: this.form.designation.trim(),
      joining_date: this.form.joining_date.substring(0, 10),
      status: this.form.status
    };

    const request = this.editingId === null
      ? this.http.post<Employee>(this.apiUrl, payload)
      : this.http.put<Employee>(
          `${this.apiUrl}/${this.editingId}`,
          payload
        );

    request.subscribe({
      next: employee => {
        console.log('Employee saved successfully:', employee);

        this.saving = false;
        this.showForm = false;
        this.editingId = null;
        this.form = this.emptyForm();
        this.formErrorMessage = '';

        this.cdr.markForCheck();

        // Refresh the employee list after saving
        this.loadEmployees();
      },

      error: (error: HttpErrorResponse) => {
        console.error('Failed to save employee:', error);

        this.formErrorMessage = this.getApiErrorMessage(
          error,
          'save the employee'
        );

        this.saving = false;
        this.cdr.markForCheck();
      }
    });
  }

  // Open custom delete warning dialog
  deleteEmployee(employee: Employee): void {
    if (
      employee.id === undefined ||
      this.deleting ||
      this.saving ||
      this.loading
    ) {
      return;
    }

    this.employeeToDelete = employee;
    this.deleteErrorMessage = '';
    this.showDeleteWarning = true;
    this.cdr.markForCheck();
  }

  // Cancel deletion
  cancelDelete(): void {
    if (this.deleting) {
      return;
    }

    this.showDeleteWarning = false;
    this.employeeToDelete = null;
    this.deleteErrorMessage = '';
    this.cdr.markForCheck();
  }

  // Delete only after user confirms
  confirmDelete(): void {
    const employee = this.employeeToDelete;

    if (
      !employee ||
      employee.id === undefined ||
      this.deleting ||
      this.saving
    ) {
      return;
    }

    this.deleting = true;
    this.deleteErrorMessage = '';
    this.cdr.markForCheck();

    this.http.delete(
      `${this.apiUrl}/${employee.id}`
    ).subscribe({
      next: () => {
        console.log('Employee deleted:', employee.id);

        this.deleting = false;
        this.showDeleteWarning = false;
        this.employeeToDelete = null;
        this.deleteErrorMessage = '';

        this.cdr.markForCheck();

        // Refresh the employee list after deletion
        this.loadEmployees();
      },

      error: (error: HttpErrorResponse) => {
        console.error('Failed to delete employee:', error);

        this.deleteErrorMessage = this.getApiErrorMessage(
          error,
          'delete the employee'
        );

        // Keep the dialog open so the error is visible
        // and the user can retry or cancel.
        this.deleting = false;
        this.cdr.markForCheck();
      }
    });
  }
}
