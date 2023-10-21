# Canteen API

This is the backend API for the Canteen App. It is built using Golang and Gin.

## Table of Contents

-   [Database Schema](#database-schema)
-   [External API](#external-api)
-   [Documentation](#documentation)
-   [Installation](#installation)
    -   [Environment Variables](#environment-variables)
    -   [Required Packages](#required-packages)
    -   [Running Program](#running-program)

## Database Schema

![image](https://github.com/backendlearning2023/backend1/assets/76490419/d62a6b5a-38ce-4e76-ae79-931e068f1817)

## External API

This API uses the following 3<sup>rd</sup> party libraries:

-   [Yelp](https://docs.developer.yelp.com/docs)

## Documentation

[API Reference](https://canteen-api.up.railway.app/api/v1/docs/index.html)

## Installation

-   ### Environment Variables

    To run this project, you will need to add the following environment variables to your .env file

    -   #### Database Connection

        `DB_HOST`

        `DB_PORT`

        `DB_USERNAME`

        `DB_PASSWORD`

        `DB_DATABASENAME`

    -   #### YELP (3<sup>rd</sup> Party) API Connection

        `API_KEY`

-   ### Required packages

```bash
  github.com/joho/godotenv@v1.5.1
  github.com/onsi/ginkgo/v2@v2.1.4
  github.com/onsi/gomega@v1.19.0
  github.com/swaggo/files@v1.0.1
  github.com/swaggo/gin-swagger@v1.6.0
  github.com/swaggo/swag@v1.16.1
  gorm.io/driver/postgres@v1.4.5
  gorm.io/gorm@v1.24.1-0.20221019064659-5dd2bb482755
  github.com/KyleBanks/depth@v1.2.1
  github.com/bytedance/sonic@v1.8.8
  github.com/chenzhuoyu/base64x@v0.0.0-20221115062448-fe3a3abad311
  github.com/gin-contrib/sse@v0.1.0
  github.com/go-openapi/jsonpointer@v0.19.6
  github.com/go-openapi/jsonreference@v0.20.2
  github.com/go-openapi/spec@v0.20.9
  github.com/go-openapi/swag@v0.22.3
  github.com/go-playground/locales@v0.14.1
  github.com/go-playground/universal-translator@v0.18.1
  github.com/go-playground/validator/v10@v10.13.0
  github.com/goccy/go-json@v0.10.2
  github.com/google/go-cmp@v0.5.8
  github.com/josharian/intern@v1.0.0
  github.com/json-iterator/go@v1.1.12
  github.com/klauspost/cpuid/v2@v2.2.4
  github.com/leodido/go-urn@v1.2.4
  github.com/lib/pq@v1.10.7
  github.com/mailru/easyjson@v0.7.7
  github.com/mattn/go-isatty@v0.0.18
  github.com/modern-go/concurrent@v0.0.0-20180306012644-bacd9c7ef1dd
  github.com/modern-go/reflect2@v1.0.2
  github.com/pelletier/go-toml/v2@v2.0.7
  github.com/twitchyliquid64/golang-asm@v0.15.1
  github.com/ugorji/go/codec@v1.2.11
  golang.org/x/arch@v0.3.0
  golang.org/x/tools@v0.9.1
  google.golang.org/protobuf@v1.30.0
  gopkg.in/yaml.v3@v3.0.1
  github.com/gin-gonic/gin@v1.9.0
  github.com/golang-jwt/jwt/v4@v4.4.3
  github.com/jackc/chunk
```

-   ### Running Program

```bash
  go run .
```
