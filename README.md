# Task Tracker CLI

A simple and efficient command-line application for managing your tasks. Task Tracker CLI allows you to add, update, delete, and mark tasks as complete or in progress. It also lets you list all tasks or filter them by their status.

**To get started, clone the repository and build the application:**

git clone [https://github.com/toemoe/task-tracker-cli.git](https://github.com/toemoe/task-tracker-cli.git)
cd task-tracker-cli
go build -o task-cli ./cmd

**After building, you can use the following commands to manage your tasks:**

Add a new task:
./task-cli add "Task Name" "Task Description"

Delete a task by its ID:
./task-cli delete <task_id>

Update a task's name and description by its ID:
./task-cli update <task_id> "New Name" "New Description"

Mark a task as done or in progress by its ID:
./task-cli mark <task_id>

List all tasks:
./task-cli list

List completed tasks:
./task-cli list-complete

List tasks that are in progress:
./task-cli list-in-progress

List tasks that are not started yet:
./task-cli list-todo

**The application is designed to be simple and intuitive, making task management from the command line efficient and straightforward.**

https://roadmap.sh/projects/task-tracker
