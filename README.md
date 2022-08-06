# 📦 Go Analytics API

It's small analytics microservice for storing data to db.

**Not ready for use!**

## Plans

- [x] Basic structure
- [x] Requests validation
- [x] Unit tests (partial coverage)
- [ ] Request for get filtered data
- [ ] Analytics visualisation (graphs)

## Contents

- [Installation](#installation)
- [How to use](#how-to-use)
    - [Compact](#compact)
    - [Detailed](#detailed)

## Installation

- Clone or download project;
    ```
    # git clone git@github.com:AlekseyPleshkov/go-analytics-api.git
    ```
- Create `config.yml` file in the project directory. Set correctly data for postgres and generate unique token for requests;
    ```
    token: 'PLACE YOUR SECUTIRY TOKEN FOR REQUESTS'
    host: '0.0.0.0:8080'
    db:
        host: 'localhost'
        port: '5432'
        username: 'postgres'
        password: 'postgres'
        database: 'analytics'
    ```
- Create docker image and start container. You can add this Dockerfile to exists docker-compose file in your ecosystem;
    ```
    # docker build --no-cache -t go-analytics-api .
    # docker run -d -p 8080:8080 --restart=always --name go-analytics-api go-analytics-api
    ```
- Or start microservice without docker;
    ```
    # go build cmd/main.go
    # ./main
    ```

## How to use

Requests should be sent with parameter `token` in header with value from `config.yml`

### Compact

Save simple analytics data with dynamic structure.

```
{POST} /analytics/v1/compact

{
    "user_id": string,
    "platform": string,
    "data": object
}
```

- `user_id` is unique user identifier
- `platform` is application who send analytics data (ios, android, device-service-api and etc.)
- `data` is object with custom key/value data for save to db

### Detailed

Save detailed analytics data with more important parameters.

```
{POST} /analytics/v1/detailed

{
    "user_id": string,
    "platform": string,
    "event": string,
    "category": string?,
    "value": string?,
    "additionals": object?
}
```

- `user_id` is unique user identifier
- `platform` is application who send analytics data (ios, android, device-service-api and etc.)
- `event` is name of analytics event (click, visit, tap, open and etc.)
- `category` is type of event (button, link, screen and etc.)
- `value` is detailed information about event ("Primary Button", "Sum field" and etc.)
- `additionals` is object with custom key/value data for save additionals information