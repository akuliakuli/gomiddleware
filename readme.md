# Middleware Task Project

This project is a simple Go application built using the Echo framework. It includes a middleware to check for the presence of a specific HTTP header and a handler that calculates the number of days left until January 1, 2025.

## Features

1. **Days Until 2025 Handler**:  
   A handler that returns the number of days left until January 1, 2025. The response is sent with an HTTP 200 OK status code.

2. **User-Role Middleware**:  
   Middleware that checks if the `User-Role` HTTP header is present and contains the value `admin`. If so, it logs "red button user detected" to the console.

## Running the Application

To run this application, follow the basic Go instructions:

1. Ensure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).

2. Clone this repository and navigate to the project directory:

   ```bash
   git clone <repository-url>
   cd middleware-task

3. Install dependencies

    ```bash
    go mod tidy

4. Run the application:
    ```bash
    go run cmd/middleware-task/main.go
