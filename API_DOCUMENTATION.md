# API Документация - Pill Reminder

Полная документация всех API endpoints приложения для управления напоминаниями о приеме таблеток.

## Базовая информация

- **Base URL**: `http://localhost:8080/api`
- **Content-Type**: `application/json`
- **Формат ID**: UUID v4

---

## 📋 Таблица содержания

1. [User API](#user-api) - Управление пользователями
2. [Reminder API](#reminder-api) - Управление напоминаниями
3. [Модели данных](#модели-данных)
4. [Коды ошибок](#коды-ошибок)
5. [Примеры использования](#примеры-использования)

---

## User API

### Регистрация пользователя

Создает новый аккаунт пользователя в системе.

**Endpoint**:
```
POST /user/register
```

**Описание**: Регистрирует нового пользователя с заданным ID

**Request Body**:
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Parameters**:
| Поле | Тип | Обязательное | Описание |
|------|-----|-------------|---------|
| id | string (UUID) | Да | Уникальный идентификатор пользователя |

**Response** (200 OK):
```json
null
```

**Response** (500 Internal Server Error):
```json
{
  "error": "description of error"
}
```

**Примеры запроса**:

=== "cURL"
    ```bash
    curl -X POST http://localhost:8080/api/user/register \
      -H "Content-Type: application/json" \
      -d '{"id": "550e8400-e29b-41d4-a716-446655440000"}'
    ```

=== "JavaScript"
    ```javascript
    fetch('http://localhost:8080/api/user/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        id: '550e8400-e29b-41d4-a716-446655440000'
      })
    })
    .then(response => response.json())
    .then(data => console.log(data));
    ```

=== "Python"
    ```python
    import requests
    
    url = 'http://localhost:8080/api/user/register'
    payload = {
        'id': '550e8400-e29b-41d4-a716-446655440000'
    }
    response = requests.post(url, json=payload)
    print(response.status_code)
    ```

---

## Reminder API

### Создание напоминания

Создает новое напоминание о приеме таблетки с расписанием дат.

**Endpoint**:
```
POST /reminder
```

**Описание**: Создает напоминание с указанным названием лекарства, дозировкой и датами напоминания

**Request Body**:
```json
{
  "medicine_name": "Аспирин",
  "dosage": "500mg",
  "description": "Принимать после еды",
  "date": [
    "2026-08-02T09:00:00Z",
    "2026-08-02T21:00:00Z",
    "2026-08-03T09:00:00Z"
  ]
}
```

**Parameters**:
| Поле | Тип | Обязательное | Описание |
|------|-----|-------------|---------|
| medicine_name | string | Да | Название лекарства (макс. 20 символов) |
| dosage | string | Да | Дозировка (макс. 20 символов), например "500mg" |
| description | string | Да | Дополнительное описание (макс. 20 символов) |
| date | array[datetime] | Да | Массив дат и времени напоминания (ISO 8601 format) |

**Response** (200 OK):
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "medicine_name": "Аспирин",
  "dosage": "500mg",
  "description": "Принимать после еды",
  "date": [
    {
      "id": "223e4567-e89b-12d3-a456-426614174001",
      "date": "2026-08-02T09:00:00Z",
      "isCompleted": false,
      "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
    },
    {
      "id": "323e4567-e89b-12d3-a456-426614174002",
      "date": "2026-08-02T21:00:00Z",
      "isCompleted": false,
      "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
    },
    {
      "id": "423e4567-e89b-12d3-a456-426614174003",
      "date": "2026-08-03T09:00:00Z",
      "isCompleted": false,
      "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
    }
  ]
}
```

**Примеры запроса**:

=== "cURL"
    ```bash
    curl -X POST http://localhost:8080/api/reminder \
      -H "Content-Type: application/json" \
      -d '{
        "medicine_name": "Аспирин",
        "dosage": "500mg",
        "description": "Принимать после еды",
        "date": [
          "2026-08-02T09:00:00Z",
          "2026-08-02T21:00:00Z"
        ]
      }'
    ```

=== "JavaScript"
    ```javascript
    const reminder = {
      medicine_name: 'Аспирин',
      dosage: '500mg',
      description: 'Принимать после еды',
      date: [
        new Date('2026-08-02T09:00:00Z'),
        new Date('2026-08-02T21:00:00Z')
      ]
    };
    
    fetch('http://localhost:8080/api/reminder', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(reminder)
    })
    .then(response => response.json())
    .then(data => console.log(data));
    ```

=== "Python"
    ```python
    import requests
    from datetime import datetime
    
    url = 'http://localhost:8080/api/reminder'
    payload = {
        'medicine_name': 'Аспирин',
        'dosage': '500mg',
        'description': 'Принимать после еды',
        'date': [
            '2026-08-02T09:00:00Z',
            '2026-08-02T21:00:00Z'
        ]
    }
    response = requests.post(url, json=payload)
    print(response.json())
    ```

---

### Получить напоминание по ID

Получает полную информацию о напоминании по его ID.

**Endpoint**:
```
GET /reminder/{id}
```

**Описание**: Возвращает напоминание и все связанные с ним даты

**Path Parameters**:
| Параметр | Тип | Описание |
|----------|-----|---------|
| id | string (UUID) | ID напоминания |

**Query Parameters**: Нет

**Response** (200 OK):
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "medicine_name": "Аспирин",
  "dosage": "500mg",
  "description": "Принимать после еды",
  "date": [
    {
      "id": "223e4567-e89b-12d3-a456-426614174001",
      "date": "2026-08-02T09:00:00Z",
      "isCompleted": false,
      "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
    }
  ]
}
```

**Response** (500 Internal Server Error):
```json
{
  "error": "Reminder not found"
}
```

**Примеры запроса**:

=== "cURL"
    ```bash
    curl -X GET http://localhost:8080/api/reminder/123e4567-e89b-12d3-a456-426614174000 \
      -H "Content-Type: application/json"
    ```

=== "JavaScript"
    ```javascript
    const reminderId = '123e4567-e89b-12d3-a456-426614174000';
    
    fetch(`http://localhost:8080/api/reminder/${reminderId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(response => response.json())
    .then(data => console.log(data));
    ```

=== "Python"
    ```python
    import requests
    
    reminder_id = '123e4567-e89b-12d3-a456-426614174000'
    url = f'http://localhost:8080/api/reminder/{reminder_id}'
    response = requests.get(url)
    print(response.json())
    ```

---

### Получить все напоминания пользователя

Получает список всех напоминаний для конкретного пользователя.

**Endpoint**:
```
GET /reminder/findByUserId
```

**Описание**: Возвращает массив всех напоминаний пользователя со всеми датами

**Query Parameters**:
| Параметр | Тип | Обязательное | Описание |
|----------|-----|-------------|---------|
| userId | string (UUID) | Да | ID пользователя |

**Response** (200 OK):
```json
[
  {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "medicine_name": "Аспирин",
    "dosage": "500mg",
    "description": "Принимать после еды",
    "date": [
      {
        "id": "223e4567-e89b-12d3-a456-426614174001",
        "date": "2026-08-02T09:00:00Z",
        "isCompleted": false,
        "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
      }
    ]
  },
  {
    "id": "423e4567-e89b-12d3-a456-426614174100",
    "medicine_name": "Витамин D",
    "dosage": "1000 IU",
    "description": "Один раз в день",
    "date": [
      {
        "id": "523e4567-e89b-12d3-a456-426614174101",
        "date": "2026-08-02T08:00:00Z",
        "isCompleted": true,
        "reminder_id": "423e4567-e89b-12d3-a456-426614174100"
      }
    ]
  }
]
```

**Response** (500 Internal Server Error):
```json
{
  "error": "Invalid userId format"
}
```

**Примеры запроса**:

=== "cURL"
    ```bash
    curl -X GET "http://localhost:8080/api/reminder/findByUserId?userId=550e8400-e29b-41d4-a716-446655440000" \
      -H "Content-Type: application/json"
    ```

=== "JavaScript"
    ```javascript
    const userId = '550e8400-e29b-41d4-a716-446655440000';
    
    fetch(`http://localhost:8080/api/reminder/findByUserId?userId=${userId}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(response => response.json())
    .then(data => console.log(data));
    ```

=== "Python"
    ```python
    import requests
    
    user_id = '550e8400-e29b-41d4-a716-446655440000'
    url = 'http://localhost:8080/api/reminder/findByUserId'
    params = {'userId': user_id}
    response = requests.get(url, params=params)
    print(response.json())
    ```

---

### Удалить напоминание

Удаляет напоминание и все связанные с ним даты.

**Endpoint**:
```
DELETE /reminder/{id}
```

**Описание**: Удаляет напоминание по ID. Также удаляются все связанные записи дат напоминания.

**Path Parameters**:
| Параметр | Тип | Описание |
|----------|-----|---------|
| id | string (UUID) | ID напоминания для удаления |

**Query Parameters**: Нет

**Response** (200 OK):
```json
"deleted"
```

**Response** (500 Internal Server Error):
```json
{
  "error": "Failed to delete reminder"
}
```

**Примеры запроса**:

=== "cURL"
    ```bash
    curl -X DELETE http://localhost:8080/api/reminder/123e4567-e89b-12d3-a456-426614174000 \
      -H "Content-Type: application/json"
    ```

=== "JavaScript"
    ```javascript
    const reminderId = '123e4567-e89b-12d3-a456-426614174000';
    
    fetch(`http://localhost:8080/api/reminder/${reminderId}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(response => response.json())
    .then(data => console.log(data));
    ```

=== "Python"
    ```python
    import requests
    
    reminder_id = '123e4567-e89b-12d3-a456-426614174000'
    url = f'http://localhost:8080/api/reminder/{reminder_id}'
    response = requests.delete(url)
    print(response.json())
    ```

---

## Модели данных

### User

Модель пользователя системы.

```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000"
}
```

| Поле | Тип | Описание |
|------|-----|---------|
| id | string (UUID) | Уникальный идентификатор пользователя |

---

### ReminderEntity

Основная сущность напоминания о приеме таблетки.

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "medicine_name": "Аспирин",
  "dosage": "500mg",
  "description": "Принимать после еды",
  "user_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

| Поле | Тип | Описание |
|------|-----|---------|
| id | string (UUID) | Уникальный идентификатор напоминания |
| medicine_name | string | Название лекарства (макс. 20 символов) |
| dosage | string | Дозировка (макс. 20 символов) |
| description | string | Описание (макс. 20 символов) |
| user_id | string (UUID) | ID пользователя-владельца |

---

### ReminderDate

Дата и статус одного напоминания.

```json
{
  "id": "223e4567-e89b-12d3-a456-426614174001",
  "date": "2026-08-02T09:00:00Z",
  "isCompleted": false,
  "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
}
```

| Поле | Тип | Описание |
|------|-----|---------|
| id | string (UUID) | Уникальный идентификатор записи даты |
| date | datetime | Дата и время напоминания (ISO 8601) |
| isCompleted | boolean | Статус выполнения (принял ли пользователь таблетку) |
| reminder_id | string (UUID) | ID связанного напоминания |

---

### CreateReminderReq

Request модель для создания напоминания.

```json
{
  "medicine_name": "Аспирин",
  "dosage": "500mg",
  "description": "Принимать после еды",
  "date": [
    "2026-08-02T09:00:00Z",
    "2026-08-02T21:00:00Z"
  ]
}
```

| Поле | Тип | Описание |
|------|-----|---------|
| medicine_name | string | Название лекарства |
| dosage | string | Дозировка |
| description | string | Описание |
| date | array[datetime] | Массив дат напоминания |

---

### ReminderResponse

Response модель при возврате напоминания.

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "medicine_name": "Аспирин",
  "dosage": "500mg",
  "description": "Принимать после еды",
  "date": [
    {
      "id": "223e4567-e89b-12d3-a456-426614174001",
      "date": "2026-08-02T09:00:00Z",
      "isCompleted": false,
      "reminder_id": "123e4567-e89b-12d3-a456-426614174000"
    }
  ]
}
```

| Поле | Тип | Описание |
|------|-----|---------|
| id | string (UUID) | ID напоминания |
| medicine_name | string | Название лекарства |
| dosage | string | Дозировка |
| description | string | Описание |
| date | array[ReminderDate] | Массив дат с статусами |

---

## Коды ошибок

| Код | Описание |
|-----|---------|
| 200 OK | Запрос успешно обработан |
| 400 Bad Request | Ошибка в формате запроса или невалидные данные |
| 404 Not Found | Ресурс не найден |
| 500 Internal Server Error | Внутренняя ошибка сервера или БД |

---

## Примеры использования

### Сценарий 1: Регистрация и создание первого напоминания

```bash
# 1. Регистрируем пользователя
USER_ID="550e8400-e29b-41d4-a716-446655440000"
curl -X POST http://localhost:8080/api/user/register \
  -H "Content-Type: application/json" \
  -d "{\"id\": \"$USER_ID\"}"

# 2. Создаем напоминание для приема таблеток
curl -X POST http://localhost:8080/api/reminder \
  -H "Content-Type: application/json" \
  -d '{
    "medicine_name": "Аспирин",
    "dosage": "500mg",
    "description": "После еды",
    "date": [
      "2026-08-02T09:00:00Z",
      "2026-08-02T21:00:00Z"
    ]
  }'

# 3. Получаем все напоминания пользователя
curl -X GET "http://localhost:8080/api/reminder/findByUserId?userId=$USER_ID" \
  -H "Content-Type: application/json"
```

### Сценарий 2: Поиск и удаление напоминания

```bash
REMINDER_ID="123e4567-e89b-12d3-a456-426614174000"

# 1. Получаем информацию о напоминании
curl -X GET http://localhost:8080/api/reminder/$REMINDER_ID \
  -H "Content-Type: application/json"

# 2. Если нужно, удаляем напоминание
curl -X DELETE http://localhost:8080/api/reminder/$REMINDER_ID \
  -H "Content-Type: application/json"
```

---

## Планируемые улучшения

- [ ] Метод обновления статуса напоминания (`UpdateReminderStatus`)
- [ ] Фильтрация напоминаний по дате
- [ ] Пагинация для больших списков
- [ ] Авторизация и аутентификация
- [ ] Дополнительная валидация входных данных

---

## Контакты и поддержка

Для вопросов или сообщений об ошибках, пожалуйста, свяжитесь с командой разработки.
