[![wakatime](https://wakatime.com/badge/github/alexanderilyin/coursera.org.svg)](https://wakatime.com/badge/github/alexanderilyin/coursera.org)

# Coursera

https://www.coursera.org/user/e469b67aca8383a1a9bf55544137cba5

## Courses

* [Getting Started with Go](https://www.coursera.org/learn/golang-getting-started/home)
* [Functions, Methods, and Interfaces in Go](https://www.coursera.org/learn/golang-functions-methods/home/welcome)

## Configuration

```bash
go mod init coursera.org
```

## Know Problems

```
$ git status
fatal: detected dubious ownership in repository at '/workspaces/coursera.org'
To add an exception for this directory, call:

        git config --global --add safe.directory /workspaces/coursera.org
```

## TODO

* Configure GH Actions to run tests and some linters.
* Configure code coverage.
  * Enable by default when running test via VS Code.
  * Create VS commands to view coverage and generate HTML.