# Getting Started with Golang App

This project aim to Test the Skills of prospective candidate for **PT. Sharing Vision Indonesia** at the position of **Full Stack Engineer**. The creator, build these Apps and running on Ubuntu environment as well.

## About This Apps

1. Using Go version 1.18.2

## Prerequisites

At the first, this is required before you running the application. At default OS, I'm using Ubuntu 20.04.

1. **Install Go**

   Linux installer console version.
    * Make sure you're on home directory\
      `cd ~`
    * Download file Go installer with curl\
      `curl -OL https://golang.org/dl/go1.18.2.linux-amd64.tar.gz`
    * Extract file and written to the /usr/local/
      directory\
      `sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz`
    * Open .profile\
      `sudo nano ~/.profile`
    * Add this script to the end of your file\
      `export PATH=$PATH:/usr/local/go/bin`
    * Refresh your profile\
      `source ~/.profile`
    * Check Go\
      `go version`
    * If success installation then it will show like this\
      `go version go1.18.2 linux/amd64`

   If Windows or Mac user, you can click [this page](https://go.dev/doc/install) for installation.


2. **Download dependencies**

     `go mod download`

## How to Run

### `go run main.go`

Runs the app main.go

### `nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go`

**Notes: This script can only run if you have installed node.js and npm. And also you should have nodemon package which install as global.**

This way is watch mode.\
The app will reload when you make changes.
