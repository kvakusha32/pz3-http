# Проверочные запросы практики №3

Сначала запустить сервер командой `go run ./cmd/server`.
Команды ниже выполнять из корня проекта `pz3-http` во втором терминале.
Команду создания задачи выполнить до запросов по ID.

1. Проверка сервера: `curl.exe -i http://localhost:8080/health`
2. Создание задачи: `curl.exe -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" --data-binary "@requests/create-task.json"`
3. Список задач: `curl.exe -i http://localhost:8080/tasks`
4. Фильтр списка: `curl.exe -i "http://localhost:8080/tasks?q=milk"`
5. Задача по ID: `curl.exe -i http://localhost:8080/tasks/1`
6. Пустой title: `curl.exe -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" --data-binary "@requests/empty-title.json"`
7. Некорректный ID: `curl.exe -i http://localhost:8080/tasks/abc`
8. Несуществующий ID: `curl.exe -i http://localhost:8080/tasks/9999`