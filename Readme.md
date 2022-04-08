# OZON INTERNSHIP REST API

## Clone 
* Clone from GitHub:

``` bash
$ git clone https://github.com/Arcady1/OZON_Internship_API.git
```

- Prepare `.env` file:

```
PORT = 8080

DB_HOST=127.0.0.1
DB_USER = <YOUR DB USER>
DB_PASSWORD = <YOUR PASSWORD>
DB_NAME = <YOUR DB NAME>
DB_SSLMODE = disable
DB_PORT=5432
```

## Build and run

``` bash
$ cd OZON_Internship_API/
$ go build -o app.bin
usage: ./app.bin l [local storage - default] 
usage: ./app.bin p [postgresql storage]
```

Server is listening on localhost: 8080

## Test

```bash
$ go test -v
```

## Docker
- Clone the repository
- Start docker:

``` bash
$ docker-compose up --build
```

Server is listening on localhost: 8080

- Stop docker:

``` bash
$ docker-compose down
```

## API

### Get an original URL by short URL

_Endpoint:_

```GET: http://127.0.0.1:8080/api/v1.0/url```

_Example request:_

```
GET: http://127.0.0.1:8080/api/v1.0/url
?short=PBA4n_2n4s
```

_Request parameters:_

| Parameter | Description                                                            | Required |
|-----------|------------------------------------------------------------------------|----------|
| short     | The short URL.                                                         | True     |

_The example of the response for the short **PBA4n_2n4s**_

```json
{
    "status": 200,
    "message": "Getting the original URL",
    "data": {
        "originalUtl": "https://google.com"
    }
}
```

_Response object:_

| Property            | Description                                          | 
|---------------------|------------------------------------------------------|
| status              | The status of the request.                           |
| message             | Status description.                                  |
| data -> originalUtl | The original URL.                                    |

---

### Get a short URL by original URL

_Endpoint:_

```POST: http://127.0.0.1:8080/api/v1.0/url```

_Example request:_

```
GET: http://127.0.0.1:8080/api/v1.0/url
?original=https://google.com
```

_Request parameters:_

| Parameter | Description                                                            | Required |
|-----------|------------------------------------------------------------------------|----------|
| original  | The original URL.                                                      | True     |

_The example of the response for the short **https://google.com**_

```json
{
    "status": 201,
    "message": "Original URL is saved",
    "data": {
        "shortUtl": "PBA4n_2n4s"
    }
}
```

_Response object:_

| Property            | Description                                          | 
|---------------------|------------------------------------------------------|
| status              | The status of the request.                           |
| message             | Status description.                                  |
| data -> shortUtl    | The shortUtl URL.                                    |
