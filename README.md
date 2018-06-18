
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
<a href="#technologies--3rd-libraries">Technologies & 3rd libraries</a> •
<a href="#development">Development</a> •
<a href="#faq">FAQ</a> •

<div align="center">
    <img src="https://github.com/NghiaTranUIT/artify-core/blob/master/images/code_block.jpg" width="100%" />
</div>

## Features
* 😍 Hand-picked 18th Arts.
* 👨‍💻 Generate beautiful wallpaper depend on your screen size.
* 👑 Automatically fetch feature art for everydays.
* 🌎 On-demand art, You can pick your favorites art-style, artist (Coming soon 🙇🏻‍♂️)
* 🍉 Open-source project.
* 💯 Totally Free.

## Technologies & 3rd libraries
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

## Development
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

## Downloads
### Artify Core
<a href='https://github.com/NghiaTranUIT/artify-core/releases/download/0.4.1/artify-core_0.4.1_macOS_64-bit.tar.gz'>Download Artify Core 0.4.1</a>

### Artify macOS app
<a href='https://github.com/NghiaTranUIT/artify-macos/releases/download/0.5.1/Artify.zip'>Download Artify 0.5.1</a>

## FAQ

* **Why Golang? Why isn't it Ruby or NodeJS?**

> As the simplicity of Golang, I finally fall in love, whereas comparing how the complexity of Swift.

> Inspired by [Talk](https://www.youtube.com/watch?v=rFejpH_tAHM) by **Rob Pike** at dotGo 2015.

* **You're mobile software Engineer. Why do you learn Golang, which intends for Backend?**
> Just aspire to become better version myself. I don't want to live in a comfortable zone permanently. Learn new stuff, new programming perspective is never fruitless 👨‍💻.

* **What is the primary purpose of the existence of Artify Core?**

> The existence of Artify Core leads to the existence of [Artify macOS app](https://github.com/NghiaTranUIT/artify-macos).

> It's kind of a heart 🍣 in your chest.

* **What is the current progress?**

> Here is [Open Ticket](https://github.com/NghiaTranUIT/artify-core/issues?q=is%3Aopen+is%3Aissue) and [Close Ticket](https://github.com/NghiaTranUIT/artify-core/issues?q=is%3Aissue+is%3Aclosed)

* **Where can I get the Artify macOS app?**
> Here [Landing Page](https://artify.launchaco.com) and [Github](https://github.com/NghiaTranUIT/artify-macos)

* **What is the technologies and 3rd libraries behind the scene?**
> Pls take a look at [Technologies](#technologies--3rd-libraries)

* **Why the code is so clumsy?**
> It's the first time I've written Golang for a production product. I aspire to learn Golang, and it's the best opportunity to apply what I've to learn 😅

> I appreciate your PR to refactor my code 🙇🏻‍♂️.

* **Where does the Artify's resource come from?**

> Every art pictures are hand-picked from [WikiArt](https://www.wikiart.org).

> If you wonder how I collect the data. Here is my partner, [Spider Man](https://github.com/NghiaTranUIT/artify-core/blob/master/scripts/spider.ruby), which is a Ruby script.

> The conjunction of [Nokogiri](http://www.nokogiri.org) and [Watir](http://watir.com) are perfect for this scenario. Indeed, I'm a lazy man, I don't want to collect data like a manual labor 😅.

* **Why do you use Docker? Is it over-engineering?***
> Docker and Docker-Compose minimize the redundant and boring steps when you setups the project runs smoothly at the first time.

> It often gets me nut when trying to install PostgreSQL driver manually and do all necessary stuff.

> Then, **Docker** is a rescue 👑.

* **How to you think about Pop and Soda?**
> [Pop](https://github.com/gobuffalo/pop) and [Soda](https://github.com/gobuffalo/pop#migrations) are best combination ever I've seen.

> I tried [Gorm](https://github.com/jinzhu/gorm), but hit the wall when I'd do a db migration as simple as possible.

> Then, I come to [Soda](https://github.com/gobuffalo/pop). It mimics the Active Record from Rail. By using Soda, I could do all relation dabase, and Migration without potential riskes.

> The [Pull Request](https://github.com/NghiaTranUIT/artify-core/issues/23) if you're interested 👨‍💻.

* **Can I become a contributor?**

> Defintely, I appreciate your effort to become a contributor. Clone the project and setup your workspace. Happy coding guys 🚢

* **Do you have personal blog?**
> Yes, I often write blog at [My lab](www.nghiatran.me) 👨‍🍳

* **How do I contact you?**

> Don't hesitate to open Issue on Github if you encounter any problems. Or give a welcome hug to me at vinhnghiatran@gmail.com.
