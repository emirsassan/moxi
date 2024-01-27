# Moxi - Usage Guide

Moxi is a lightweight build tool designed to be a simple alternative to CMake, focusing on simplicity and efficiency. This guide provides detailed information on using Moxi for your projects.

## Table of Contents

- [Tasks and Moxifile](#tasks-and-moxifile)
- [Running Tasks](#running-tasks)
- [Task Groups](#task-groups)
- [Running All Tasks in a Group](#running-all-tasks-in-a-group)
- [Advanced Usage](#advanced-usage)

## Tasks and Moxifile

Moxi uses a straightforward syntax for defining tasks in a `Moxifile`. A task is defined as follows:

```mo
task: <task-name> => <command>
```

- `<task-name>`: A unique identifier for the task.
- `<command>`: The command or set of commands to execute for the task.

Here's an example `Moxifile`:

```mo
task: build => go build -o myapp main.go
task: test => go test ./...
```

## Running Tasks

To execute a task, use the following command:

```bash
moxi <task-name>
```

Replace `<task-name>` with the actual task you want to run. If no task is specified, Moxi will run all tasks defined in the `Moxifile`.

Example:

```bash
moxi build
```

## Task Groups

Moxi supports task grouping for better organization. Groups are defined as follows:

```
group <group-name> {
    task: <task-name> => <command>
    task: <task-name> => <command>
}
```

Here's an example:

```
group build {
    task: compile => go build -o myapp main.go
    task: test => go test ./...
}
```

To run a task within a group:

```bash
moxi <group-name>:<task-name>
```

Example:

```bash
moxi build:compile
```

## Running All Tasks in a Group
To run all tasks within a group, use the following command:

```bash
moxi <group-name>
```
Example:

```bash
moxi build
```
This command executes all tasks defined within the specified group.

## Advanced Usage

For any questions or issues open an [issue](https://github.com/emirsassan/moxi/issues).

Happy building with Moxi!