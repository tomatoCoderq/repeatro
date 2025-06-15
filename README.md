# Repeatro – Anki-Style Vocabulary Learning App

**Repeatro** is a modern web-based vocabulary learning tool inspired by Anki, built using Go and PostgreSQL

## Features

- Spaced repetition for efficient vocabulary retention (using SM2 algorithm)
- JWT-based user authentication
- Decks to organize your vocabulary by topic or language
- [In progress...] Language detection using [lingua-go](https://github.com/pemistahl/lingua-go)
- [In progress...] RESTful API with [Swaggo](https://github.com/swaggo/swag) auto-generated Swagger docs
- PostgreSQL backend with [Goose](https://github.com/pressly/goose) for migrations

## Getting Started

### Prerequisites

- Go 1.18+ && PostgreSQL 15+
- [Goose](https://github.com/pressly/goose)

### Installation

```bash
git clone https://github.com/tomatoCoderq/repeatro.git
cd repeatro

go mod tidy
```

### Starting
1. CREATE DATABASE repeatro;
2. Set up goose (Check out https://github.com/pressly/goose for setup information)
3. Set up config.toml (Check out example in the root)
4. go run cmd/app/main.go
N. Additionally air can be used for auto server restart

## RoadMap

### Near future tasks
1. Add import/export via csv, json
2. Enhance current stats

### Future tasks
1. Add simple front