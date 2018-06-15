
<p align="center">
  <img src="https://github.com/NghiaTranUIT/artify-macos/blob/master/images/logo.png" alt="Artify App Logo" width="600" height="auto"/>
</p>

<h2>
  Artify
  <a href="https://circleci.com/gh/NghiaTranUIT/artify-core">
    <img src="https://circleci.com/gh/NghiaTranUIT/artify-core/tree/master.svg?style=svg" alt="Build Status">
  </a>
  <a href="https://github.com/NghiaTranUIT/artify-macos/releases/tag/0.5.1">
    <img src="https://img.shields.io/badge/version-0.5.1-green.svg" alt="0.5.1">
  </a>
  <a href="./LICENSE">
    <img src="https://img.shields.io/badge/license-GPL--3.0-blue.svg">
  </a>
</h2>

A Golang Backend core of [Artify](https://github.com/NghiaTranUIT/artify-macos), which is an application for bringing dedicatedly 18th century Arts to everyone.

<a href="#features">Features</a> •
<a href="#downloads">Downloads</a> •
<a href="#technologies">Technologies</a> •
<a href="#faq">FAQ</a> •

### Features
* 😍 Hand-picked 18th Arts.
* 👨‍💻 Generate beautiful wallpaper depend on your screen size.
* 👑 Automatically fetch feature art for everydays.
* 🌎 On-demand art, You can pick your favorites art-style, artist (Coming soon 🙇🏻‍♂️)
* 🍉 Open-source project.
* 💯 Totally Free.

### Technologies & 3rd libraries
* [Golang](https://golang.org)
* [PostgreSQL](https://www.postgresql.org)
* [Docker](https://www.docker.com)
* [Docker-Compose](https://docs.docker.com/compose/)
* [Echo](https://github.com/labstack/echo)
* [Dep](https://github.com/golang/dep)
* [Pop](https://github.com/gobuffalo/pop)
* [Soda Database Migration](https://github.com/gobuffalo/pop/tree/master/fizz)
* [Go Releaser](https://github.com/goreleaser/goreleaser)
* [CircleCI](https://circleci.com)

### Development
Fortunately, Artify Core is composed by Docker and Docker-Composed, so we don't need to install PostgreSQL manually.
In order to start the server, just following this short instruction.

* Clone this project to your GO_PATH (/go/src/github.com/NghiaTranUIT/artify-core).
* Open `docker-compose.yml` and update your postgre ENV
* Open `ArtifyWorkspace.xcworkspace`
```
POSTGRES_DB: your_db_name
POSTGRES_USER: your_username
POSTGRES_PASSWORD: your_password
```

* Happy coding 😍

## Installation
1. Clone project to your GO_PATH (/go/src/github.com/NghiaTranUIT/artify-core)
2. Open `docker-compose.yml` and update your postgre ENV
```
POSTGRES_DB: your_db_name
POSTGRES_USER: your_username
POSTGRES_PASSWORD: your_password
```
2. Open `database.yml` and fill same configuration.
```
development:
  dialect: postgres
  database: your_db_name
  user: your_username
  password: your_password
  host: db
  pool: 5
```
3. Execute those commands 
```
# docker-compose build
# docker-compose up -d
```
4. Enjoy your trip 😍

### Downloads Artify Core
<a href='https://github.com/NghiaTranUIT/artify-core/releases/download/0.4.1/artify-core_0.4.1_macOS_64-bit.tar.gz'>Download Artify Core 0.4.1</a>

[More Download Options](https://github.com/NghiaTranUIT/artify-core/releases)

### Downloads Artify macOS app
<a href='https://github.com/NghiaTranUIT/artify-macos/releases/download/0.5.1/Artify.zip'>Download Artify 0.5.1</a>

[More Download Options](https://github.com/NghiaTranUIT/artify-macos/releases)
