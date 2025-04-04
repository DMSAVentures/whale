# Matrix Operations API

A Go-based HTTP API that processes CSV matrix files and performs basic matrix operations: **invert**, **sum**, **multiply**, and **flatten**.

---

## 🚀 Features

- Upload a CSV file containing integers
- Perform the following matrix operations:
   - **Invert**: Transpose the matrix
   - **Sum**: Calculate the sum of all elements (with overflow detection)
   - **Multiply**: Calculate the product of all elements (with overflow detection)
   - **Flatten**: Output a comma-separated list of all elements
- Well-tested API with table-driven integration tests
- Graceful handling of invalid or malformed input

---

## 📦 Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── main.go                # Starts the HTTP server
├── internal/
│   ├── api/               # HTTP handlers
│   ├── matrixoperations/  # Core matrix logic and safety utils
│   └── utils/             # CSV parsing utilities
├── test/                  # API tests
```

---

## 🔧 Usage

### Run locally

```bash
go run .
```

Then test with:

```bash
curl -F 'file=@test/matrix.csv' http://localhost:8080/invert
```

### Run with Docker

```bash
make run
```

### Run API tests

```bash
make integration-tests
```

---

## API Endpoints

| Endpoint     | Description                       | Method |
|--------------|-----------------------------------|--------|
| `/invert`    | Transposes the matrix             | `GET`  |
| `/sum`       | Sums all matrix elements          | `GET`  |
| `/multiply`  | Multiplies all matrix elements    | `GET`  |
| `/flatten`   | Flattens matrix into CSV string   | `GET`  |

---

## 📁 Example Matrix (matrix.csv)

```csv
1,2,3
4,5,6
7,8,9
```

---

## 📋 Example Responses

- **Invert**:
  ```
  1,4,7
  2,5,8
  3,6,9
  ```

- **Sum**:
  ```
  45
  ```

- **Multiply**:
  ```
  362880
  ```

- **Flatten**:
  ```
  1,2,3,4,5,6,7,8,9
  ```

---

## Unit Testing

```bash
make test
```

---

## Future Improvements

- Support for floating-point
- Health check and `/status` endpoint
- More detailed validation and error messages
