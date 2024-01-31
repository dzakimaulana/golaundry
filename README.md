<p align="center">
  <a href="" rel="noopener">
<<<<<<< HEAD
 <img width=317.6px height=178.7px src="https://www.filepicker.io/api/file/O8dz87hXSheB05h3nO4M" alt="Project logo"></a>
</p>

<h3 align="center">Golaundry</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![Golang](https://img.shields.io/badge/Go-%2300ADD8.svg?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![Docker](https://img.shields.io/badge/Docker-%230db7ed.svg?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)
[![Fiber](https://img.shields.io/badge/Fiber-%2320232a.svg?style=flat&logo=fiber&logoColor=%2361DAFB)](https://github.com/gofiber/fiber)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-%23336791.svg?style=flat&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Gorm](https://img.shields.io/badge/Gorm-%2300ADD8.svg?style=flat&logo=gorm&logoColor=white)](https://gorm.io/)

</div>

## 🤔 About

Hi, folks! this is a backend application for laundry bussines who can make the owner more easier to manage. In this app I using golang with fiber framework to build API service and gorm to get data from Postgres. The owner will be have admin account who have previledge to all endpoint and his employee only doing add transaction. I use JWT token and session to create authorization and authentication.

## 🚦Getting Started

### 🐋 Using Docker

You can pull the latest Docker image for GoLaundry from Docker Hub:

```
docker pull kymmsenpai/golaundry:latest
```

### ✨ Clone Github

1. **Clone the Repository**
```
git clone https://github.com/dzakimaulana/golaundry.git
```
2. **Docker Compose**
```
docker compose up 
```

## 📄 Api Documentation
### Authorization
Will be saved in a ```cookie``` with name ```my-session``` 
### Customer
| Method |        Endpoint        | Admin |
|:-------|:----------------------:|------:|
| POST   | `/api/customer/add`    |  YES  |
| GET    | `/api/customer/get/:id`|   NO  |
| GET    | `/api/customer/get`    |   NO  |
### Item
| Method |        Endpoint      | Admin |
|:-------|:--------------------:|------:|
| POST   | `/api/item/add`      |  YES  |
| GET    | `/api/item/get/:id`  |   NO  |
| GET    | `/api/item/get`      |   NO  |
| UPDATE | `/api/item/update`   |  YES  |
| DELETE | `/api/item/delete`   |  YES  |
### Transaction
| Method |        Endpoint           | Admin |
|:-------|:-------------------------:|------:|
| POST   | `/api/transaction/add`    |   NO  |
| GET    | `/api/transaction/get/:id`|   NO  |
| GET    | `/api/transaction/get`    |   NO  |
### User
| Method |        Endpoint            | Admin |
|:-------|:--------------------------:|------:|
| POST   | `/api/item/login`          |  YES  |
| POST   | `/api/item/create-user`    |   NO  |
| GET    | `/api/item/get`            |   NO  |
| GET    | `/api/item/get/:id`        |  YES  |
| PUT    | `/api/item/reset-password` |  YES  |
| POST   | `/api/item/reset-password` |  YES  |

### 🗄️File Structure
In my project development practices, the adoption of the Repository Pattern plays a pivotal role in enhancing the modularity, maintainability, and testability of our codebase. 
```
golaundry/
├───.github
│   └───workflows
├───bin
├───cmd
│   └───golaundry
│   
├───internal
│   ├───customers
│   ├───items
│   ├───transactions
│   ├───transitems
│   └───users
├───pkg
│   ├───database
│   ├───middlewares
│   ├───models
│   ├───routes
│   └───utils
└───servers.json
```
=======
 <img width=200px height=200px src="https://www.filepicker.io/api/file/O8dz87hXSheB05h3nO4M" alt="Project logo"></a>
</p>

<h3 align="center">Golaundry</h3>
>>>>>>> 09fe03598d789674c670925a6e06b54349517575
