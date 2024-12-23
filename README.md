 # Todo List CLI

A simple command-line todo list application written in Go.

## Features

- Add a new todo item
- List all todo items
- Mark a todo item as completed
- Delete a todo item
- Edit a todo item

## Installation

### Clone the repository:
```sh
git clone https://github.com/koushik0329/CLI-todo.git
cd CLI-todo
```
### Initialize the Go module:

```sh
go mod init github.com/koushik0329/CLI-todo
```
### Install the required dependencies:
```sh
go get github.com/aquasecurity/table
```
### Usage

You can run the application using go run ./ without building an executable file.

### Commands
#### Add a new todo item:
```sh
go run ./ -add "Your todo item"
```
This command adds a new todo item with the specified title.

#### List all todo items:

```sh
go run ./ -list
```
This command lists all the todo items in a table format.

#### Mark a todo item as completed:

```sh
go run ./ -toggle <item-id>
```
This command toggles the completion status of the todo item with the specified ID.

#### Delete a todo item:

```sh
go run ./ -del <item-id>
```
This command deletes the todo item with the specified ID.

#### Edit a todo item:
```sh
go run ./ -edit <item-id>:<new-title>
```
This command edits the todo item with the specified ID to have the new title.

### Example
#### Add a new todo item:

```sh
go run ./ -add "Buy groceries"
```
#### List all todo items:

```sh
go run ./ -list
```
#### Mark a todo item as completed:
```sh
go run ./ -toggle 1
```
#### Delete a todo item:
```sh
go run ./ -del 2
```
#### Edit a todo item:
```sh
go run ./ -edit 3:"Update todo title"
```
#### Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
