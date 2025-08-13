# TestTask

REST API на Go для управления сущностями Task с асинхронным логированием.
time-work - 3 часа
## Запуск

1. Клонировать репозиторий:
git clone https://github.com/Obamakryt/TestTask.git
cd TestTask
2. Запустить сервер:
   go run main.go
Сервер будет доступен на http://localhost:8080.


3. Endpoints

- Получить все задачи
  GET /tasks
- Пример http://localhost:8080/tasks

- Получить задачу по ID
- GET /tasks?id=<task_id>
- Пример http://localhost:8080/tasks?id=1

- Создать новую задачу
- POST /tasks  Content-Type: application/json  Тело запроса:
- {
  "id": "1",
  "title": "Новая задача"
  }
- Для graceful shutdown нажмите Ctrl+C. В течение 10 секунд все операции завершаться
