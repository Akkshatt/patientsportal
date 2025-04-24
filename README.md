# Patient Portal Backend

This is a simple backend implementation of a Patient Portal for managing patient data with role-based access. It contains two main portals:
- **Receptionist Portal**: For managing patient information.
- **Doctor Portal**: For viewing and updating patient details.

## Tech Stack
- **Golang**: Backend framework.
- **Gin**: Web framework for Golang.
- **PostgreSQL**: Database.
- **JWT**: For user authentication.

## Setup Instructions

### 1. Install PostgreSQL
You'll need to have PostgreSQL installed on your machine. You can download it from the official site: [Download PostgreSQL](https://www.postgresql.org/download/)

Once installed, make sure the PostgreSQL server is running and accessible. To verify, run:
```bash
psql --version
```

### 2. Clone the Repository
Clone the repository to your local machine:
```bash
git clone https://github.com/yourusername/patient-portal.git
```

Navigate to the project directory:
```bash
cd patient-portal
```

### 3. Set Up Database
Create a database and user in PostgreSQL:
```sql
CREATE DATABASE patient_portal;
CREATE USER portal_user WITH ENCRYPTED PASSWORD 'password123';
GRANT ALL PRIVILEGES ON DATABASE patient_portal TO portal_user;
```

### 4. Create `.env` File
Create a `.env` file in the root directory of the project with the following content:
```env
DB_HOST=localhost
DB_USER=portal_user
DB_PASSWORD=password123
DB_NAME=patient_portal
DB_PORT=5432
JWT_SECRET=supersecretkey
```

### 5. Install Dependencies
Run the following command to install all the required dependencies:
```bash
go mod tidy
```

### 6. Run the Application
Start the server with the following command:
```bash
go run cmd/main.go
```

The server will be available at http://localhost:8080.

## API Endpoints

### 1. `POST /login`
**Description**: Login for both receptionists and doctors to receive a JWT token.

**Request Body**:
```json
{
  "username": "user1",
  "password": "password123"
}
```

**Response**:
```json
{
  "token": "your-jwt-token-here"
}
```

### 2. `GET /patients`
**Description**: Get a list of all patients (Doctors only).

**Headers**:
```text
Authorization: Bearer <JWT_TOKEN>
```

**Response**:
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "age": 30,
    "gender": "Male",
    "address": "123 Main St",
    "notes": "Some notes about the patient"
  }
]
```

### 3. `POST /patients`
**Description**: Add a new patient (Receptionists only).

**Request Body**:
```json
{
  "name": "Jane Doe",
  "age": 28,
  "gender": "Female",
  "address": "456 Oak St",
  "notes": "Some important notes"
}
```

**Response**:
```json
{
  "message": "Patient added successfully"
}
```

### 4. `PUT /patients/{id}`
**Description**: Update patient details (Doctors only).

**Request Body**:
```json
{
  "name": "Jane Doe Updated",
  "age": 29,
  "gender": "Female",
  "address": "456 Oak St Updated",
  "notes": "Updated notes"
}
```

**Response**:
```json
{
  "message": "Patient details updated successfully"
}
```

## Unit Testing
To run unit tests, use the following command:
```bash
go test ./...
```

## API Documentation
You can document your APIs using Postman or Swagger for better understanding and easier use.
* Postman Collection
* Swagger Documentation

## Contributing
1. Fork the repo.
2. Create a new branch: `git checkout -b feature/your-feature`.
3. Make your changes.
4. Commit the changes: `git commit -m "Add new feature"`.
5. Push to the branch: `git push origin feature/your-feature`.
6. Create a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
