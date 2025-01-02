# Project Name
REST API project menggunakan Golang, Fiber, MongoDB, dan Logrus.

## Table of Contents
- [Overview](#overview)
- [Technologies](#technologies)
- [Features](#features)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Database Migration](#database-migration)
- [Testing](#testing)
- [Directory Structure](#directory-structure)
- [API Documentation](#api-documentation)
- [License](#license)

## Overview
Project ini adalah REST API yang dikembangkan dengan Golang menggunakan framework Fiber untuk membuat server, MongoDB sebagai database, Logrus untuk pencatatan log, dan Golang-Migrate untuk migrasi database. Mockery digunakan untuk mocking dalam pengujian unit, memudahkan pembuatan tes yang terisolasi.

## Technologies
- **Golang**: Bahasa pemrograman yang cepat dan ringan untuk pengembangan aplikasi.
- **Fiber**: Web framework yang sederhana dan cepat, terinspirasi oleh Express.
- **MongoDB**: Database NoSQL untuk penyimpanan data yang fleksibel dan scalable.
- **Logrus**: Logger yang sangat fleksibel untuk menangani berbagai level logging dan output JSON.
- **Golang-Migrate**: Alat migrasi database untuk mengelola skema MongoDB.
- **Mockery**: Alat untuk otomatisasi pembuatan mock objek, berguna untuk pengujian unit.

## Features
- CRUD untuk resource utama
- Logging dengan Logrus untuk pencatatan aktivitas aplikasi
- Migrasi database menggunakan Golang-Migrate
- Pengujian unit dengan mocking menggunakan Mockery

## Installation

### Prerequisites
- **Golang**: Pastikan Golang sudah diinstal. [Download Golang](https://golang.org/dl/)
- **MongoDB**: Install dan jalankan MongoDB di sistem lokal atau gunakan layanan cloud seperti MongoDB Atlas.
- **Golang-Migrate**: Instal golang-migrate dengan menjalankan:
    ```sh
    go install -tags 'mongodb' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```

### Clone Repository
Clone project ke lokal:

```sh
git clone https://github.com/username/project-name.git
cd project-name
```

### Install Dependencies
Gunakan go mod untuk mengunduh dependencies:

```sh
go mod tidy
```

## Configuration
Konfigurasi aplikasi dikelola melalui variabel lingkungan (environment variables). Buat file .env di root project dan atur konfigurasi berikut:

```env
MONGO_DB_URL=localhost
MONGO_DB_PORT=27017
MONGO_DB_DATABASE=local
```

## Running the Application
Untuk menjalankan aplikasi, gunakan perintah berikut:

```sh
go run main.go
```