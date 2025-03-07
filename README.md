# Tasks

Приложение для управления задачами. Тестовое задание, выполненное для компании [SkillsRock](https://skillsrock.ru/).

## Возможности

- Создание новых задач
- Получение списка задач
- Изменение задач
- Удаление задач

## Особенности

- RESTful API
- Контейнеризация с Docker и Docker Compose
- Отказоустойчивость в конкурентной среде
- Чтение переменных среды из файла
- Использован `Fiber` роутер
- Миграции базы данных с [`goose`](https://github.com/pressly/goose)
- Типобезопасные запросы для взаимодействия с базой данных с [`sqlc`](https://github.com/sqlc-dev/sqlc)
- Линтинг с GitHub Actions
- Bash-скрипт [`tasks.sh`](tasks.sh) для упрощённого запуска, остановки и тестирования приложения

## Требования

- Go (версия 1.22 и выше)
- Docker

## Начало работы

### Клонирование репозитория

Сначала клонируйте репозиторий:

```bash
git clone https://github.com/chtozamm/skillsrock-tasks.git
cd skillsrock-tasks
```

### Настройка переменных среды

В корне проекта создайте файл `.env` с переменными среды:

```env
DB_URL="postgres://skillsrock:secret@localhost:5432/tasks"
PORT="8080"
```

### Запуск приложения

Запустите контейнеры с приложением и сервером PostgreSQL, используя **Docker Compose**:

```bash
docker compose up -d
```

### Применение миграций

Примените миграции для создания таблиц в базе данных:

```bash
goose postgres -dir sql/schema postgres://skillsrock:secret@localhost:5432/tasks up
```

### Создание новой задачи

Создайте новую задачу:

```bash
curl http://localhost:8080/tasks -X POST \
	-H "Content-Type: application/json" \
	-d '{"title": "Тестовое задание", "description": "Выполнить тестовое задание для компании SkillsRock", "status": "in_progress"}'
```

### Изменение задачи

Измените статус выполнения задачи

```bash
task_id=1
curl http://localhost:8080/tasks/$task_id -X PUT \
	-H "Content-Type: application/json" \
	-d '{"status": "done"}'
```

### Получение списка задач

```bash
curl http://localhost:8080/tasks
```

### Удаление задачи

```bash
curl http://localhost:8080/tasks/1 -X DELETE
```

## Использованные технологии

- Go
- [Fiber](https://github.com/gofiber/fiber)
- Docker & Docker Compose
- PostgreSQL
- [`sqlc`](https://github.com/sqlc-dev/sqlc) для генерации типобезопасного кода для взаимодействия с базой данных из SQL-запросов
- [`goose`](https://github.com/pressly/goose) для управления миграциями

## Заключение

Спасибо за внимание! С удовольствием готов ответить на любые вопросы и обсудить возможные доработки.
