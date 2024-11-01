
# Go URL Shortener Program

This program is a URL shortener written in Go. It reads URL mappings from both hard-coded paths and a YAML configuration, redirects users based on these mappings, and serves content through HTTP.

## Features

- **Path-to-URL Mapping**: Define URL redirections through a map of paths to URLs or by loading a YAML configuration.
- **Configurable YAML Input**: Add URL mappings using a YAML file for greater flexibility.
- **Multiplexer Setup**: A default multiplexer handles requests not specified in the mappings.
- **Graceful Error Handling**: The program logs errors for missing configurations or server failures, ensuring clean exits.

## Usage

1. **Run the Program**:
   Use Go to run the program. Start the server and make sure it listens on port 8080:
   ```bash
   go run main.go
   ```

2. **Default URL Mappings**:
   - You can add hardcoded URL mappings in the `pathsToUrls` map within `main.go`.
   - You can also specify YAML mappings in the `yaml` variable for additional paths.

3. **YAML Configuration Format**:
   - Each mapping in YAML should specify a path and a URL:
     ```yaml
     - path: /learn
       url: https://go.dev/learn/
     ```
   - Paths not found in `pathsToUrls` or the YAML configuration will fall back to the default multiplexer.

4. **HTTP Server**:
   The server listens on port 8080 by default. Open a browser and navigate to `http://localhost:8080/{path}` (e.g., `/doc` or `/learn`) to test redirections.

## Packages Used

- **`net/http`**: Provides the server functionality, routing, and request handling.
- **`gopkg.in/yaml.v2`**: Parses YAML configuration to dynamically load paths and URLs.
- **`fmt` and `os`**: For printing, error handling, and program termination.

## Code Structure

- **`main`**: Initializes the multiplexer, loads URL mappings, and starts the HTTP server on port 8080.
- **`MapHandler`**: Redirects requests based on the hardcoded `pathsToUrls` map and defaults to the multiplexer for unmapped paths.
- **`YAMLHandler`**: Parses and maps YAML configuration paths, creating a handler that falls back to the `MapHandler` if a path is not found.
- **`parseYAML`**: Converts YAML data into structured mappings.
- **`buildPathMap`**: Builds a fast-access map from parsed YAML data.
- **`defaultMux`**: Sets up a basic HTTP multiplexer with a root path handler displaying a welcome message.

## Example Execution

1. **Run the Server**:
   ```bash
   go run main.go
   ```

2. **Access URL Mappings**:
   - Navigate to `http://localhost:8080/doc` to be redirected to the Go documentation.
   - Go to `http://localhost:8080/pkg` to access the Go package documentation.
   - Visit `http://localhost:8080/learn` for the Go learning resources as defined in YAML.

## Possible Improvements

- **Dynamic YAML Reloading**: Enable dynamic reloading of YAML without restarting the server.
- **Load YAML from External File**: Modify the program to load the YAML configuration from an external file, allowing users to easily adjust mappings without editing the code.
- **JSON Configuration Support**: Extend configuration options to allow JSON or other formats.
- **Advanced Path Matching**: Implement pattern matching for more flexible redirection rules.

## License

This project is open source and available under the MIT license.