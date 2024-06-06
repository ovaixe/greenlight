## Overview
The Greenlight Project is designed to showcase a practical application of backend development using Go. Inspired by the "Let's Go Further" book by Alex Edwards, this project focuses on creating a high-performance, reliable, and easy-to-maintain backend service for managing movie data.

## Getting Started
To get started with the Greenlight Project, follow the instructions below.
### Prerequisites
-    Go (version 1.16 or higher)
-    Git
-    PostgreSQL

### Installation
#### Clone the repository
```bash
git clone https://github.com/ovaixe/greenlight.git
cd greenlight
```
#### Set up environment variables
Create a .envrc file in the root directory and add the following variables:
```bash
export GREENLIGHT_DB_DSN=your_database_connection_string
```
#### Install dependencies
```bash
make vendor
```

### Usage
To run the project locally:
#### Start the application
```bash
make run/api
```

### API Endpoints
-     GET /v1/movies: Retrieve a list of movies
-     POST /v1/movies: Add a new movie
-     GET /v1/movies/{id}: Retrieve a specific movie by ID
-     PATCH /v1/movies/{id}: Update an existing movie
-     DELETE /v1/movies/{id}: Delete an existing movie
-     POST /v1/users: Register a new user
-     PUT /v1/users/activated: Activate a registerd user
-     POST /v1/tokens/authentication: Create a authentication token for user

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.