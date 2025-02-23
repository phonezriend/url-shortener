# URL Shortener

A simple **URL Shortener** built with **Go (Golang)**, **Redis**, and **Docker**.  
The application lets you shorten long URLs into shorter links and automatically redirects users when they access the short URL.

## Features
- Shorten long URLs into compact links
- Auto-redirect to the original URL
- Redis-powered URL storage
- Dockerized for easy deployment
- Web interface built with HTML, CSS, and JavaScript
- Basic URL expiration (24 hours)

---

## Prerequisites
Before running the project, ensure you have the following installed:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)

---

## Installation

1. **Clone this repository:**

```bash
git clone https://github.com/phonezriend/url-shortener.git
cd url-shortener
```

2. **Build and run the project using Docker Compose:**

```bash
docker-compose up --build -d
```

3. **Check running containers:**

```bash
docker ps
```
You should see two containers running:

    url-shortener-app (The main Go application)
    redis (Redis database)

### 📤 Shorten a URL (via API using `curl`)

You can also use a `curl` command to shorten a URL from the terminal:

```bash
curl -X POST http://localhost/shorten \
-H "Content-Type: application/json" \
-d "{\"url\": \"https://example.com/very/long/url\"}"
