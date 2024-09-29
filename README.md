# task-management

## Running the application
`docker compose up --build -d`

## Restart/Rerun the application
Delete docker containers and volumes
and then run the `docker compose up --build -d`

## DB Schema
```
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL
    );
```

``` CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        description TEXT,
        user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
        status VARCHAR(50) CHECK (status IN ('Pending', 'In Progress', 'Completed')) DEFAULT 'Pending',
        due_date DATE,
        created_at TIMESTAMP DEFAULT NOW(),
        updated_at TIMESTAMP DEFAULT NOW()
    );
```

## Swagger Document

http://localhost:8080/docs/index.html

## Task Management APIs

- Create Task: Endpoint to create a new task. A task must have a title, description, assigned user, status (Pending, In Progress,     Completed), and due date.
- Get Tasks: Endpoint to retrieve all tasks or filter tasks by status, user, or due date.
- Update Task: Endpoint to update task details (e.g., change status, description, etc.).
- Delete Task: Endpoint to delete a task.

### Get All Tasks

```curl -X 'GET'
  'http://localhost:8080/v1/tasks'
  -H 'accept: application/json'
```

### Get Task By Id

```curl -X 'GET'
  'http://localhost:8080/v1/tasks/1'
  -H 'accept: application/json'
```

### Create Task

```curl -X 'POST'
  'http://localhost:8080/v1/tasks'
  -H 'accept: application/json'
  -H 'Content-Type: application/json'
  -d '{
  "description": "test",
  "due_date": "2024-09-29T18:25:43.511Z",
  "status": "Pending",
  "title": "test",
  "user_id": 1
}'
```

### Update Task

```curl -X 'PUT'
  'http://localhost:8080/v1/tasks'
  -H 'accept: application/json'
  -H 'Content-Type: application/json'
  -d '{
  "description": "Description",
  "id": 1,
  "status": "Pending",
  "title": "Title"
}'
```

### Delete Task

```curl -X 'DELETE'
  'http://localhost:8080/v1/tasks/1'
  -H 'accept: application/json'
```

## User Management APIs
 - Create User: Endpoint to create a new user. A user has a name and email

### GetAllUsers

```
curl -X 'GET'
  'http://localhost:8080/v1/user'
  -H 'accept: application/json'
```

### Get User By EmailId

```
curl -X 'GET'
  'http://localhost:8080/v1/user/rkpundhir90%40gmail.com'
  -H 'accept: application/json'
```


### Create User

```
curl -X 'POST' \
  'http://localhost:8080/v1/user'
  -H 'accept: application/json'
  -H 'Content-Type: application/json'
  -d '{
  "email_id": "rkpundhir90@gmail.com",
  "name": "Rohit Pundhir"
}'
```