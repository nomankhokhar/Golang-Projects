# 50 Golang

## Basic Go HTTP Server

This is a simple HTTP server implemented in Go. The server serves static files from the `./src` directory, and provides two routes: `/hello` and `/form`.

- /form
- /hello
- / server static html files to the frontend

## Build a CRUD API with Golang

Frontend build with tailwinds simple no any framework, Table show the Movies within the table there are Buttons for Edit and Delete the Movie Records

Backend is build with Golang Gorilla Mux framework that have 5 API. Other info is given below.

- Edit the Movie
- Delete the Movie
- Update the Movie
- Create the Movie
- Get All the Movie

## Email Verifier

This is program written in Golang that use its own package net to check Email is verified or not like gmail.com, yahoo.com, facebook.com etc

Run the program First with following command

```
go run main.go
```

then, pass the email you want to verify.

## Slack Bot File Uploader

This program is written in Golang that is use to upload file and bunch of files

## Basic Golang API with MongoDB

This program is Written in Golang and Database is MongoDB

Below are the Endpoint

- Get User
- Create User
- Delete User

## Weather Tracker

This program is Written in Golang with Http package and frontend is build with Tailwind CSS

Below are the Endpoint

- Hello
- weather

## LoadBalancer in Golang

This is a **Load Balancer** implementation written in Go that distributes incoming HTTP requests across multiple backend servers using a **round-robin algorithm**.

### Features

- **Round-Robin Load Balancing**: Distributes requests evenly across servers in a cyclic manner
- **Reverse Proxy**: Uses Go's `net/http/httputil.ReverseProxy` to forward requests to backend servers
- **Health Checking**: Built-in server availability checking (IsAlive method)
- **Interface-Based Design**: Clean abstraction using Go interfaces for extensibility

### How It Works

1. **Startup**: Initializes 3 backend servers (Facebook, Bing, DuckDuckGo)
2. **Listen**: Listens for HTTP requests on `http://localhost:8000`
3. **Route**: Each incoming request is forwarded to the next available server in round-robin order
4. **Proxy**: Uses reverse proxy to transparently forward requests and responses

### Request Flow

```
Request 1 → Facebook
Request 2 → Bing
Request 3 → DuckDuckGo
Request 4 → Facebook (cycles back)
Request 5 → Bing
...
```

### Run the Program

```bash
go run main.go
```

Then access: `http://localhost:8000`

### Components

- **Server Interface**: Defines behavior for backend servers (Address, IsAlive, Serve)
- **simpleServer Struct**: Implements the Server interface with reverse proxy
- **LoadBalancer Struct**: Manages multiple servers and distributes traffic using round-robin
- **Round-Robin Algorithm**: Uses modulo operator to cycle through servers

### Skill Level

Intermediate - Covers interfaces, structs, methods, HTTP handling, and algorithm implementation.
