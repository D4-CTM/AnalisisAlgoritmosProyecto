# Analisis de Algoritmos proyecto Grupo \#1

> [!IMPORTANT]
> As a prerequisite Golang must be installed in the system.

How to run:
```bash
cd backend
go mod tidy
# This will take some seconds the first time as it will fetch the gin dependency.
go run .
```

By default it runs on `http://localhost:5461` and it has the following placeholders to work upon:
- `http://localhost:5461/api/color-graphs`
- `http://localhost:5461/api/alg/2`
- `http://localhost:5461/api/alg/3`

You can stop the server by pressing `ctrl + c`.
