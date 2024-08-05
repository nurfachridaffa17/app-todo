
# App To Do

The **To-Do List Application** is a user-friendly and efficient task management tool designed to help individuals and teams organize their daily activities. With this application, users can create, manage, and track tasks with ease, ensuring productivity and effective time management.



## Tech Stack

- **Backend**: Go Programming Language(Golang) Echo Framework
- **Database**: MySQL or PostgreSQL for data storage and management
- **Authentication**: JSON Web Tokens (JWT) for secure authentication
- **Version Control**: Git and GitHub for code management and collaboration

## Features

- **User Authentication**: Secure login and registration system to protect user data and ensure privacy.
- **Task Management**: Create, update, delete, and manage tasks efficiently.
- **Categorization**: Organize tasks into different categories for better management and retrieval.
- **Status Tracking**: Track task progress with statuses like Pending, In Progress, and Completed.
- **Responsive Design**: Accessible on various devices, including desktops, tablets, and mobile phones.
- **User-Specific Data**: Each user's tasks and categories are stored independently, providing a personalized experience.

## Installation

Follow these steps to set up and run the application locally:

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/nurfachridaffa17/app-todo.git
   cd todo-list-app
   ```
2. **Setup Env**:
   ```bash
    cp .env.development.test .env.development
   ```

3. **Running my project**:

```bash
  go run main.go
```

4. **Add Swagger**

```bash
  swag init
```
    