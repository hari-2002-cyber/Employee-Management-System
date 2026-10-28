# Employee Management System

## 1. Project Overview

The Employee Management System is a web-based application developed to manage employee information in an organized and efficient way.

The application allows users to view employee records, add new employees, update existing employee details, search for employees, filter records, and delete employee information.

The project uses Angular for the frontend, Go (Golang) for the backend REST API, and PostgreSQL as the database.

This project demonstrates full-stack web development, REST API integration, database connectivity, and CRUD (Create, Read, Update, Delete) operations.

## 2. Project Objectives

- To maintain employee information in a centralized database.
- To provide a user-friendly interface for employee management.
- To perform CRUD operations through REST APIs.
- To make it easier to search and filter employee records.
- To separate the frontend, backend, and database into distinct layers.

## 3. Technologies Used

 Technology  Purpose 
------
 Angular  Frontend user interface 
 TypeScript  Frontend application logic 
 HTML and CSS  Page structure and styling 
 Go (Golang)  Backend development 
 PostgreSQL  Relational database 
 REST API  Communication between frontend and backend 
 Visual Studio Code  Development environment 

## 4. Features Implemented

### Employee Management
- Add a new employee.
- View all employee records.
- Edit existing employee information.
- Delete employee records with a confirmation dialog.
- Store employee information in PostgreSQL.

### Search and Filtering
- Search employees using the available search field.
- Filter employees by department.
- Filter employees by employment status.
- Clear filters to display the full employee list.

### Dashboard
- Display the total number of employees.
- Display the number of active employees.
- Display the number of inactive employees.
- Present employee information in a structured table.

### Employee Information

Each employee record contains the following fields

- Employee ID
- Name
- Email
- Department
- Designation
- Joining Date
- Employment Status (Active or Inactive)

## 5. Project Structure

The project is organized into frontend and backend folders.

```text
employee-management-system
│
├── backend
│   ├── config
│   ├── model
│   ├── repository
│   ├── service
│   ├── handler
│   ├── main.go
│   ├── go.mod
│   └── go.sum
│
├── frontend
│   ├── src
│   │   └── app
│   │       ├── app.ts
│   │       ├── app.html
│   │       └── app.css
│   ├── package.json
│   └── angular.json
│
└── README.md
```

Note The exact files and folders may vary slightly depending on the final project setup.

### Frontend

The Angular frontend provides the dashboard, employee table, search and filter controls, and forms for adding and editing employees.

### Backend

The Go backend handles HTTP requests, processes employee data, communicates with the PostgreSQL database, and returns API responses.

### Database

PostgreSQL stores employee records persistently.

## 6. Prerequisites

Install the following software before running the project

1. Visual Studio Code
2. Go
3. Node.js and npm
4. Angular CLI
5. PostgreSQL

Verify the installations in the VS Code terminal

```powershell
go version
node -v
npm -v
ng version
```

To check PostgreSQL, run

```powershell
& CProgram FilesPostgreSQL18binpsql.exe --version
```

If PostgreSQL is installed in a different location or version, adjust the path accordingly.

## 7. Database Setup

### Step 1 Open PostgreSQL

Open SQL Shell (psql) or connect to PostgreSQL using your preferred database tool.

### Step 2 Create the database

```sql
CREATE DATABASE employee_management;
```

Connect to the database

```sql
c employee_management
```

### Step 3 Create the employees table

Run the following SQL statement

```sql
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    department VARCHAR(100) NOT NULL,
    designation VARCHAR(100) NOT NULL,
    joining_date DATE NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'Active'
        CHECK (status IN ('Active', 'Inactive'))
);
```

The database is now ready to store employee information.

Important Keep your PostgreSQL username and password private. Do not upload database credentials to a public repository.

## 8. How to Run the Project

The backend and frontend must run simultaneously in two separate terminals.

### Step 1 Open the project

Open Visual Studio Code.

Select File → Open Folder and open

```text
Demployee-management-system
```

### Step 2 Start the Go backend

Open the first terminal in VS Code.

Navigate to the backend folder

```powershell
cd Demployee-management-systembackend
```

Set the PostgreSQL password for the current PowerShell session

```powershell
$secure = Read-Host Enter PostgreSQL password -AsSecureString
$envDB_PASSWORD = [System.Net.NetworkCredential]new(, $secure).Password
```

Enter the password for your PostgreSQL user when prompted.

Start the backend

```powershell
go run .
```

Keep this terminal open. The API should start on port `8080`.

Backend API URL

httplocalhost8080apiemployees

### Step 3 Start the Angular frontend

Open a second terminal in VS Code.

Navigate to the frontend folder

```powershell
cd Demployee-management-systemfrontend
```

If dependencies have not been installed, run

```powershell
npm install
```

Start the Angular development server

```powershell
ng serve
```

Wait for Angular to finish compiling.

### Step 4 Open the application

Open your browser and visit

httplocalhost4200

The Employee Management System dashboard should now be visible.

Keep both terminals running while using the application.

## 9. API Endpoints

The backend exposes REST API endpoints for employee management.

 Method  Endpoint  Purpose 
---------
 GET  `apiemployees`  Retrieve all employees 
 GET  `apiemployees{id}`  Retrieve a specific employee 
 POST  `apiemployees`  Create a new employee 
 PUT  `apiemployees{id}`  Update an existing employee 
 DELETE  `apiemployees{id}`  Delete an employee 

The endpoints are available through

`httplocalhost8080apiemployees`

## 10. How to Use the Application

### Add an Employee

1. Open the dashboard.
2. Click the Add Employee button.
3. Enter the employee's name, email, department, designation, and joining date.
4. Select the employment status.
5. Submit the form.
6. Verify that the employee appears in the table.

### View Employees

The dashboard displays the employee records retrieved from the backend and PostgreSQL database.

### Search for Employees

Enter a name, email, or other supported search text in the search field to find matching records.

### Filter Employees

Select a department or employment status from the filter controls. Use Clear Filters to reset the selections.

### Edit an Employee

1. Select the edit action for an employee.
2. Modify the required fields.
3. Save the changes.
4. Verify the updated information in the table.

### Delete an Employee

1. Select the delete action for an employee.
2. Review the confirmation dialog.
3. Confirm the deletion to remove the record, or cancel to keep the employee.

## 11. Troubleshooting

### Backend cannot connect to PostgreSQL

- Confirm that the PostgreSQL service is running.
- Verify the database name, username, and password.
- Confirm that the `DB_PASSWORD` environment variable is set in the backend terminal.
- Make sure the `employee_management` database exists.

### Port 8080 is already in use

The backend may already be running in another terminal. Check the existing terminal before starting another instance.

### Port 4200 is already in use

Angular may already be running. Open the existing application URL or stop the previous development server before starting another one.

### Frontend cannot retrieve employee data

- Confirm that the Go backend is running.
- Check that the API URL is correct.
- Check the browser developer console for errors.
- Verify that the backend and frontend are configured to communicate correctly.

### Employee records are not visible

Verify that the database contains records and that the search, department, and status filters are not excluding them.

## 12. Security and Submission Notes

- Do not include database passwords or other secrets in the submitted source code.
- Do not include `node_modules`, generated build output, or unnecessary cache folders in the project ZIP.
- Include `go.mod`, `go.sum`, `package.json`, and the application source files.
- Include screenshots of the working dashboard and important application features.
- For a production deployment, configure secure credentials, appropriate CORS settings, and HTTPS.

## 13. Conclusion

The Employee Management System provides a structured solution for managing employee records using Angular, Go, and PostgreSQL. It demonstrates frontend development, backend REST API implementation, relational database operations, and the integration of multiple technologies into a single web application.

The project can be extended in the future with authentication, role-based access control, pagination, employee profile images, and reporting features.
