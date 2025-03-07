# Обзор API

### Получение списка задач

`GET /tasks`  
**Статус ответа**: 200 OK | 500 Internal Server Error  
**Тело ответа**: 

```json
[
    {
        "id": 1,
        "title": "My task",
        "description": "Task description",
        "status": "new",
        "created_at": "2025-01-01T00:00:00.000000Z",
        "updated_at": "2025-01-01T00:00:00.000000Z",
    }
]
```

### Создание новой задачи

`POST /tasks`  
**Заголовки запроса**: `"Content-Type": "application/json"`  
**Тело запроса**:

```json
{
    "title": "My task",
    "description": "Task description",
}
```

**Статус ответа**: 201 Created | 400 Bad Request | 500 Internal Server Error  
**Тело ответа**: 

```json
{
    "id": 1,
    "title": "My task",
    "description": "Task description",
    "status": "new",
    "created_at": "2025-01-01T00:00:00.000000Z",
    "updated_at": "2025-01-01T00:00:00.000000Z",
}
```

### Изменение задачи

`PUT /tasks/:task_id`  
**Заголовки запроса**: `"Content-Type": "application/json"`  
**Тело запроса**:

```json
{
    "title": "My updated task",
    "description": "Updated task description",
    "status": "done",
}
```

**Статус ответа**: 200 OK | 400 Bad Request | 404 Not Found | 500 Internal Server Error  
**Тело ответа**: 

```json
{
    "id": 1,
    "title": "My updated task",
    "description": "Updated task description",
    "status": "done",
    "created_at": "2025-01-01T00:00:00.000000Z",
    "updated_at": "2025-03-08T09:37:12.982716Z",
}
```

### Удаление задачи

`DELETE /tasks/:task_id`  
**Статус ответа**: 204 No Content | 500 Internal Server Error  

