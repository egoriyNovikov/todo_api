# Todo API с Go и PostgreSQL

Простое REST API для управления задачами, построенное на Go и PostgreSQL с использованием Docker.

## Структура проекта

```
.
├── cmd/
│   └── main.go          # Главный файл приложения
├── internal/            # Внутренние пакеты
├── pkg/                 # Публичные пакеты
├── Dockerfile           # Docker образ для Go приложения
├── docker-compose.yaml  # Docker Compose конфигурация
├── init.sql            # SQL скрипт инициализации базы данных
├── .dockerignore       # Исключения для Docker
└── go.mod              # Go модули

```

## Быстрый старт

### Предварительные требования

- Docker
- Docker Compose

### Запуск приложения

1. Клонируйте репозиторий:
```bash
git clone <repository-url>
cd second_weack
```

2. Запустите приложение с помощью Docker Compose:
```bash
docker-compose up --build
```

3. Приложение будет доступно по адресу:
   - API: http://localhost:8080
   - PostgreSQL: localhost:5432

### Остановка приложения

```bash
docker-compose down
```

### Остановка с удалением данных

```bash
docker-compose down -v
```

## Переменные окружения

| Переменная | Описание | Значение по умолчанию |
|------------|----------|----------------------|
| `DB_HOST` | Хост базы данных | postgres |
| `DB_PORT` | Порт базы данных | 5432 |
| `DB_USER` | Пользователь БД | todo_user |
| `DB_PASSWORD` | Пароль БД | todo_password |
| `DB_NAME` | Имя базы данных | todo_db |
| `PORT` | Порт приложения | 8080 |

## API Endpoints

- `GET /` - Главная страница
- `GET /health` - Проверка здоровья приложения

## Разработка

### Локальная разработка без Docker

1. Установите Go 1.24+
2. Установите PostgreSQL
3. Создайте базу данных и пользователя согласно `init.sql`
4. Запустите приложение:
```bash
go run cmd/main.go
```

### Сборка Docker образа

```bash
docker build -t todo-api .
```

### Запуск только базы данных

```bash
docker-compose up postgres
```

## База данных

PostgreSQL 15 с предустановленной схемой:
- Таблица `todos` с полями: id, title, description, completed, created_at, updated_at
- Тестовые данные для демонстрации

## Полезные команды

```bash
# Просмотр логов
docker-compose logs -f

# Подключение к базе данных
docker-compose exec postgres psql -U todo_user -d todo_db

# Перезапуск только приложения
docker-compose restart app

# Просмотр статуса контейнеров
docker-compose ps
```

```bash
# Проверить статус контейнеров
docker-compose ps

# Подключиться к базе данных
docker exec -it todo_postgres psql -U todo_user -d todo_db

# Посмотреть все таблицы
docker exec -it todo_postgres psql -U todo_user -d todo_db -c "\dt"

# Выполнить SQL запрос
docker exec -it todo_postgres psql -U todo_user -d todo_db -c "SELECT * FROM todos;"

# Посмотреть логи PostgreSQL
docker-compose logs postgres
```