# Corona Today

This service sends useful news about COVID-19 to your inbox!

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

```
go
git
```

### Installing

1. Clone this repository

2. Go to ```cmd/web/mail.go``` and fix this code.
    ```go
   auth := smtp.PlainAuth("", "your gmail user id", "your gmail user password", "smtp.gmail.com")
    ```

### Run the app

1. Start the server
    ```bash
    $ cd coronatoday
    $ go run cmd/web/*
    ```

2. Browse to `http://localhost:4000/`
