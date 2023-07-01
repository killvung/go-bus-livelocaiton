# go-bus-livelocaiton

Speedrun to create a simple web application written in Go using the Gin framework and Gorm library. It ingests bus live location data from a specified endpoint and stores it in an SQLite database. (any %, any times)

## Requirements

- Go 1.16 or higher
- SQLite

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/go-web-app.git
   ```
2. Change into the project directory:
   ``` bash
   cd go-web-app
   ```
3. Install the dependencies:
   ```bash
   go mod download
   ```
4. Build the application:
   ```bash
   go build
   ```

5. Run the application:
    ```bash
    ./go-web-app
   The application should now be running on http://localhost:8080.
   ```

## Configuration

The application does not require any additional configuration. By default, it uses an SQLite database named bus.db for storing the bus location data.

## API Endpoint

/ingest: Triggers the ingestion of bus live location data from the specified endpoint into the SQLite database.

## Directory Structure

```bash
.
├── main.go             # Main entry point of the application
├── go.mod              # Go module file
├── go.sum              # Go module checksum file
├── handlers/           # Directory for HTTP request handlers
│   └── handlers.go
├── models/             # Directory for data models or structures
│   └── vehicle.go
├── routes/             # Directory for managing routes
│   └── routes.go
├── database/           # Directory for handling database setup and connection
│   └── database.go
├── templates/          # Directory for HTML templates
│   └── index.html
└── static/             # Directory for static assets
    ├── css/
    │   └── style.css
    └── js/
        └── main.js
```

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

Feel free to modify the README.md file to suit your specific project requirements and provide more details about your application, such as additional features, deployment instructions, or usage examples.
