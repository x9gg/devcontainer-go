# devcontainer-go

This is a minimalist Go web application that serves as a proof of concept to verify that Go is installed and running correctly on your system.

## Features

- Basic HTTP server with "Hello World" message
- Endpoint to display Go version information
- Simple codebase for easy understanding and modification

## Requirements

- Go (any recent version), in our image `h3xsh/image:go` it is 1.24.1

## How to Run

1. Clone this repository or download the files

2. Navigate to the project directory

3. Run the application:
   ```
   go run main.go
   ```

4. The server will start on port 8080 by default (you can change this by setting the PORT environment variable)

5. Open a web browser and visit:
   - http://localhost:8080/ - Shows the Hello World message
   - http://localhost:8080/version - Shows Go version information

### Running with Auto-Reload (Development Mode)

To use Air for hot-reloading without requiring a global installation, you can use Go modules to run it directly:

1. Run the application with Air via Go:
   ```
   go run github.com/cosmtrek/air@latest
   ```

This will download Air as a dependency and run it in the current project. Any changes you make to your Go files will automatically trigger a rebuild and restart of the server!

Note: The Air configuration is stored in the `.air.toml` file.

## Building an Executable

To build a standalone executable:

```
go build -o goapp
```

Then run the executable:

```
./goapp    # On Linux/Mac
goapp.exe  # On Windows
```

## Project Structure

- `main.go` - Contains the server code and route handlers
- `README.md` - This documentation file
- `.air.toml` - Configuration for the Air hot-reload tool

## Customizing

- Edit the message in the `handleRoot` function to change the greeting
- Add more routes and handlers to extend functionality
- Modify the port by setting the PORT environment variable

## Verifying Go Installation

This application provides a quick way to verify:

1. That Go is properly installed on your system
2. Which version of Go you're running
3. That your Go environment can compile and run web applications

If the application runs successfully, your Go installation is working correctly!