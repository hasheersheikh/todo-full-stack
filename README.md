# React + Go Todo App

This repository contains a full-stack Todo application built with React, TypeScript, Vite, Chakra UI on the client side, and Go with Fiber and MongoDB on the server side.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
  - [Clone the Repository](#clone-the-repository)
  - [Install Dependencies](#install-dependencies)
- [Running the Client](#running-the-client)
- [Running the Server](#running-the-server)
  - [Using Air for Live Reloading (Optional)](#using-air-for-live-reloading-optional)
- [Deployment](#deployment)
  - [Client Deployment](#client-deployment)
  - [Server Deployment](#server-deployment)
- [Environment Variables](#environment-variables)
- [License](#license)

## Features
- Create, update, and delete todos
- Mark todos as completed
- Responsive UI with Chakra UI
- Backend API with Go Fiber
- MongoDB for data storage

## Prerequisites
Before you begin, ensure you have the following installed:
- **Node.js** (v14 or higher)
- **Go** (v1.23.4 or higher)
- **MongoDB**
- **Air** (optional, for live reloading the Go server)

## Installation

### Clone the Repository
Clone this repository to your local machine:

```bash
git clone https://github.com/your-username/react-go-todo-app.git
cd react-go-todo-app
```

### Install Dependencies
Client
Navigate to the client directory and install the dependencies:

```bash
cd client
npm install
```

### Server
Navigate to the server directory and install the dependencies:

```bash
cd server
go mod tidy
```



## Running the Client
### Navigate to the client directory:

```bash
cd client
npm run dev
```
The client application will be available at http://localhost:5173.

## Running the Server
### Navigate to the server directory:

``` bash
cd server
```
Create a .env file in the server directory with the following content: env

```
MONGODB_URL=mongodb://localhost:27017/todo-app
PORT=4000
GO_ENV=development
Replace MONGODB_URL with your MongoDB connection string if necessary.
```

### Start the server:

```bash
go run main.go

# The server will be available at http://localhost:4000.

# Using Air for Live Reloading (Optional)
# If you have Air installed, you can use it for live reloading of the Go server during development.
# Install Air if you haven't already:

go install github.com/cosmtrek/air@latest

```

### Key Notes:
- Replace `your-username` with your GitHub username or the appropriate repository URL.
- If you have any custom instructions for Air or deployment, feel free to adjust accordingly.
