# Affiliate Backend

### 1. Prerequisites:

- Golang

### 2. Install dependencies:

```bash
$ go mod download
```

### 3. Create file `app.env` from `app.env.example`

```bash
  cp .env.example .env.local
```

### 4. Update `.env.local` file

### 5. Database Migration Guide

```bash
 (for local env)
$ make migrate_database
```

### 5.1 create migration file

```bash
 (for local env)
$ migrate create -ext sql -dir db/migrations/ -seq {migration name}
```

### 6. Running the app

```bash
# development
$ go run cmd/main.go
```

Navigate to your [host](http://localhost:8000) to check the server is online.
