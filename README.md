# Pill Reminder Application

Приложение для управления напоминаниями о приеме таблеток (лекарств).

## Требования

- Docker
- Docker Compose

## Быстрый старт

### 1. Клонирование репозитория

```bash
git clone <repository-url>
cd pillReminder
```

### 2. Конфигурация окружения

Создайте файл `.env` на основе `.env.example`:

```bash
cp .env.example .env
```

Измените значения в `.env` если необходимо:

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=reminder_user
DB_PASSWORD=reminder_password
DB_NAME=medicine-reminder-db
APP_PORT=8080
GIN_MODE=release
```

### 3. Запуск приложения

#### С помощью Docker Compose:

```bash
docker-compose up --build
```

Приложение будет доступно по адресу: `http://localhost:8080`

#### Только база данных:

```bash
docker-compose up -d postgres
```

Это запустит только PostgreSQL контейнер. Приложение можно запустить локально:

```bash
go run main.go
```

### 4. Остановка приложения

```bash
docker-compose down
```

Для удаления также данных БД:

```bash
docker-compose down -v
```

## API Endpoints

Полная документация API доступна в файле [API_DOCUMENTATION.md](./API_DOCUMENTATION.md).

### Краткий список

**User API:**
- **POST** `/api/user/register` - Регистрация нового пользователя

**Reminder API:**
- **POST** `/api/reminder` - Создать напоминание
- **GET** `/api/reminder/{:id}` - Получить напоминание по ID
- **GET** `/api/reminder/findByUserId` - Получить все напоминания пользователя
- **DELETE** `/api/reminder/{:id}` - Удалить напоминание

## Переменные окружения

| Переменная | Значение по умолчанию | Описание |
|-----------|----------------------|---------|
| DB_HOST | postgres | Хост базы данных |
| DB_PORT | 5432 | Порт базы данных |
| DB_USER | reminder_user | Пользователь БД |
| DB_PASSWORD | reminder_password | Пароль БД |
| DB_NAME | medicine-reminder-db | Название БД |
| APP_PORT | 8080 | Порт приложения |
| GIN_MODE | release | Режим Gin (debug/release) |

## Структура проекта

```
.
├── Dockerfile           # Docker образ для приложения
├── docker-compose.yml   # Конфиг Docker Compose
├── .env.example         # Пример переменных окружения
├── main.go              # Главная точка входа
├── models.go            # Модели данных
├── entity.go            # Сущности
├── service.go           # Бизнес-логика
├── go.mod               # Go модули
└── go.sum               # Go зависимости
```

## Разработка

### Локальное развертывание

Установите Go 1.26 и PostgreSQL, затем:

```bash
go mod download
go run main.go
```

### Обновление зависимостей

```bash
go mod tidy
```

## Решение проблем

### Приложение не может подключиться к БД

1. Убедитесь, что PostgreSQL контейнер запущен: `docker-compose ps`
2. Проверьте логи: `docker-compose logs postgres`
3. Убедитесь, что переменные окружения установлены правильно в `.env`

### Порт уже занят

Измените `APP_PORT` и `DB_PORT` в `.env`:

```env
APP_PORT=8081
DB_PORT=5433
```

Затем пересоздайте контейнеры:

```bash
docker-compose down
docker-compose up --build
```

## Автор

Medicine Reminder Team
