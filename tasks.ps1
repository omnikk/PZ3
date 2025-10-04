param(
    [string]$target = "run"
)

switch ($target) {
    "run" {
        Write-Host "Запуск сервера..."
        go run ./cmd/server
    }
    "build" {
        Write-Host "Сборка exe..."
        go build -o .\bin\server.exe ./cmd/server
    }
    "test" {
        Write-Host "Запуск юнит-тестов..."
        go test ./...
    }
    default {
        Write-Host "Неизвестная цель. Используйте run/build/test"
    }
}
