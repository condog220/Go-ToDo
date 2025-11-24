# Go ToDo CLI

A simple command-line ToDo list application written in Go. This tool allows you to manage your daily tasks efficiently

## Features

- **Interactive Menu**: Easy-to-use command-line interface.
- **Task Management**:
  - **Add**: Create new tasks.
  - **Remove**: Delete unwanted tasks by index.
  - **Update**: Edit the description of existing tasks.
  - **Complete**: Mark tasks as done and track completion time.
  - **List**: View all tasks with status and timestamps.
- **Persistence**: Save your task list to a JSON file for backup or future use.

## Project Structure

- [main.go](cci:7://file:///d:/Go/ToDo/main.go:0:0-0:0): The entry point of the application. Handles the CLI menu and user input.
- [todo.go](cci:7://file:///d:/Go/ToDo/todo.go:0:0-0:0): Defines the [Task](cci:2://file:///d:/Go/ToDo/todo.go:8:0-13:1) structure and core logic (add, remove, update, etc.).
- [json.go](cci:7://file:///d:/Go/ToDo/json.go:0:0-0:0): Handles saving and loading tasks to/from JSON files.

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) installed on your machine.

### Installation

1. Clone the repository:
   
   ```bash
   git clone <https://github.com/condog220/ToDo.git>
   ```

2. Navigate to the directory
   ```bash
   cd ToDo
   ```

3. Run the application
   ```bash
   go run main.go
   ```
