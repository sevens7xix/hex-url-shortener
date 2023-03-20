# hex-url-shortener

A simple URL Shortener microservice trough Hexagonal architecture in Go.

## Functional requirements

- Shorten a URl: given a specific URL, the microservice should be able to Shorten it and return the short functional URL.
- Resolve a Short URl: given a short URL, the microservice must be able of resolve it and make the proper redirection.

### Posible project structure

```
app/
├─ cmd/
│  ├─ router/
│  │  ├─ main.go
├─ pkg/
│  ├─ utilities/
│  │  ├─ hasher.go
├─ internal/
│  ├─ core/
│  │  ├─ services/
│  │  │  ├─ service.go
│  │  ├─ ports/
│  │  │  ├─ ports.go
│  │  ├─ domain/
│  │  │  ├─ data.go
│  ├─ handlers/
|  |  ├─ router.go
│  ├─ repositores/

```
