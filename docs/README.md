# Golang Tutorial

## Installation

Visit this [link](https://go.dev/doc/install) to install `go` on your machine.

> [!IMPORTANT]
> This Markdown is made by my personal researches, so if you want to learn from better source, I strongly recommand you to [Take a tour on Golang](https://go.dev/tour/)

### What is GOPATH?

In Golang, `GOPATH` is an environment variable that specifies the **root directory of your Go workspace**. The Go workspace is where Go code is organized and managed. It includes source code, compiled binaries, and package objects.

#### Structure of `GOPATH`

A typical GOPATH directory contains three main subdirectories:

1. `src`: This directory holds the source code of your Go projects and external packages.
2. `pkg`: This directory stores compiled package files, which Go uses to speed up builds.
3. `bin`: This directory stores compiled binary executables.

For example, if GOPATH is set to `/home/user/go`, the structure looks like:

```
/home/user/go/
├── bin/
├── pkg/
└── src/
```

> [!NOTE]
> Since the introduction of `Go modules` *(starting with Go 1.11)*, **the use of GOPATH has become less critical**. With modules, Go projects can exist outside the GOPATH directory, and dependencies are managed in a `go.mod` file. This change enables Go developers to organize projects more flexibly without relying on the strict workspace layout.
> In module mode, the `GOPATH` is mainly used for caching downloaded dependencies in the `GOPATH/pkg/mod directory`.

**Key Points**

- Before modules, all Go projects had to reside within the GOPATH/src.
- With modules enabled, you can place your projects anywhere, but GOPATH still acts as a default location for dependency caching and storage.
- Use the go env command to check your current GOPATH:
```
go env GOPATH
```
