# Dealls Test

This is a simple web application written in Go using the Gin framework. It provides user authentication functionality with MySQL database integration.

## Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/airlanggatirta/dealls.git
   ```

2. **Install dependencies:**
   Navigate to the project directory and install the required dependencies using Go Modules:
   ```bash
   cd golang-web-app
   go mod tidy
   ```

3. **Set up the database:**
   - Create a MySQL database for the application.
   - Update the database configuration in the `config/config.go` file with your database credentials.

4. **Run the application:**
   Start the Go application by running the following command:
   ```bash
   go run main.go
   ```

5. **Access the application:**
   Once the application is running, you can access it at `http://localhost:8080` in your web browser.

## Usage

- **API Docs:**
  Postman endpoints are located under folder apidocs.

## Configuration

- The application configuration can be found in the `config/config.go` file. You can modify settings such as the server port, database credentials, etc., as needed.


## Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework for Go.
- [GORM](https://gorm.io/): ORM library for Go, used for interacting with the database.
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt): Library for hashing passwords securely.
- [MySQL](https://github.com/go-gorm/mysql): MySQL database driver for Go.

## Contributing

- Contributions are welcome! Feel free to submit bug reports, feature requests, or pull requests to improve the application.

## License

- This project is licensed under the [MIT License](LICENSE).
