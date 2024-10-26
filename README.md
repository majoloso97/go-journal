# Go Journal
## Simple journaling CLI tool to learn the basics of the Go programming language
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) ![SQLite](https://img.shields.io/badge/Sqlite-003B57?style=for-the-badge&logo=sqlite&logoColor=white)

### About the project
Go Journal is a basic command-line interface (CLI) tool focused on Journaling. The CLI currently supports two main flags:
- `go-journal --setup`: Sets up database (runs migrations to ensure database state). Must be run initially before other commands are run.
- `go-journal --retrieve-all`: Retrieves all of your Journal entries
- `go-journal --create-new`: Allows to add a new Journal entry

The built project is stored in the root of the repository, so it can be easily downloaded.

### Learning path
Coming from a Python background, I find Go's syntax initially a bit unfamiliar, but it's simple enough to grasp the basics quickly. Though initially the language felt somewhat verbose, I've found it confortable to read/write, more than initially expected.

Go's documentation is a charm to work with. It's been easy to find things I don't know, and in most cases the official documentation is very helpful. Still in progress of getting familiar with Go's overall programming style, this time by reading the language specification.

As a next step with Go-Journal, I'm thinking to extend this project by building a web version (API only).
