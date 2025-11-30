# Driver Service (Go)

Невеликий сервіс на **Go**, який демонструє базові принципи **Clean Architecture**.  
Зберігає водіїв у **Postgres** та останню позицію водія у **Redis**.

## Функціонал
- Додати водія (Postgres)
- Додати/оновити позицію водія (Redis)
- Отримати водія + останню позицію
- Health-check

## Технології
- Go 
- Postgres
- Redis
- Pure net/http
- sqlx/pgx
