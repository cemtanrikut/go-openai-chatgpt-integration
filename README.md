# Go OpenAI ChatGPT Integration

A simple REST API built with Go that integrates with OpenAI's ChatGPT API to process user prompts and return AI-generated responses.

---

## Test Coverage Summary

The project maintains a high test coverage for all key components:

- handlers/: 100% coverage
- services/: 95% coverage
- config/: 100% coverage
- utils/: 100% coverage

Overall test coverage is 98%.

## Features

- Accepts user prompts via a REST API endpoint.
- Connects to OpenAI's ChatGPT API for AI-generated responses.
- Fully tested with a high test coverage.
- Containerized using Docker for easy deployment.

---

## Requirements

- Go 1.20 or later
- OpenAI API key
- Docker (optional, for containerized deployment)

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/cemtanrikut/go-openai-chatgpt-integration.git
cd go-openai-chatgpt-integration
```

### 2. Set Up Environment Variables

Create a .env file in the root of the project and add the following:

```bash
OPENAI_API_KEY=your_openai_api_key
OPENAI_API_URL=https://api.openai.com/v1/chat/completions
```

### 3. Run the Application

Locally (without Docker)

```bash
go run main.go
```

The application will run at http://localhost:8080 .

## API Endpoints

### POST /chat

Description: Accepts a user prompt and returns the AI-generated response.

- Request Body:

```json
{
  "prompt": "Your question or text here"
}
```

- Response:

```json
{
  "response": "AI-generated response here"
}
```

- Error Responses:

* 400 Bad Request: Invalid request payload.

* 500 Internal Server Error: Issues communicating with the OpenAI API.

## Testing

Run the following command to execute all tests:

```bash
go test ./... -v
```

To generate a test coverage report:

```bash
make coverage
```

To view an HTML report of the coverage:

```bash
make coverage-html
```

## Docker Usage

### 1. Build the Docker Image

```bash
docker build -t go-openai-chatgpt-integration .
```

### 2. Run the Docker Container

```bash
docker run -p 8080:8080 --env-file .env go-openai-chatgpt-integration
```

The application will be accessible at http://localhost:8080.

## Docker Compose (Optional)

If you prefer using Docker Compose, create a docker-compose.yml file:

```yaml
version: "3.8"
services:
  app:
    build:
      context: .
    ports:
      - "8080:8080"
    env_file:
      - .env

```

Run the following command:

```bash
docker-compose up --build
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (feature/new-feature).
3. Commit your changes.
4. Push to the branch.
5. Submit a pull request.

## License

This project is licensed under the MIT License.