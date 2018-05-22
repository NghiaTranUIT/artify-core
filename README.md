
<div align="center">
    <img src="https://github.com/NghiaTranUIT/artify-core/blob/master/images/logo.png" width="70%" />
</div>

## Artify 
Bring 18th century Art to everyone.    
  
[![CircleCI](https://circleci.com/gh/NghiaTranUIT/artify-core/tree/master.svg?style=svg)](https://circleci.com/gh/NghiaTranUIT/artify-core/tree/master)

## MacOS app
[Artify macOS App](https://github.com/NghiaTranUIT/artify-macos) is in progress.

## Requirement
1. Latest Docker
2. Golang 1.10.1

## Installation
1. Clone project to your GO_PATH (/go/src/github.com/NghiaTranUIT/artify-core)
2. Open `docker-compose.yml` and update your postgre ENV
```
POSTGRES_DB: your_db_name
POSTGRES_USER: your_username
POSTGRES_PASSWORD: your_password
```
2. Open `constant/constant.go`
```
PostgresDockerName      = "db" // Dont' change
PostgresPort            = "5432"
PostgresUser            = "your_username"
PostgresDB              = "your_db_name"
PostgresPassword        = "your_password"
```
3. Execute those commands 
```
# docker-compose build
# docker-compose up -d
```
4. Enjoy your trip 😍
