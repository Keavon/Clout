### Create a new game

- `POST /api/create`

#### Request

- Body:
```json
{
  "username": "<username>"
}
```

#### Response

- Code: `201 CREATED`
- Body:
```json
{
  "gameid": "<gameid>",
  "token": "<token>",
  "country": {
    "name": "<name>"
  }
}
```

### Join a existing game

- `POST /api/create`

#### Request

- Body:
```json
{
  "username": "<username>",
  "gameid": "<game id>"
}
```

#### Response

- Code: `201 CREATED`
- Body:
```json
{
  "token": "<token>",
  "country": {
    "name": "<name>"
  }
}
```

### Get player

- `GET /api/player`

#### Request

- Headers: `Authorization: <token>`

#### Response

- Code: `200 OK`
- Body:
```json
{
  "name": "<username>",
  "admin": true,
  "country": {
    "name": "<name>"
  }
}
```

### Purchase One Installation

- `GET /api/purchase/<resource id>`

#### Request

- Headers: `Authorization: <token>`

#### Response

- Code: `200 OK`

### Set number of installations operational

- `GET /api/operational/<resource id>`

#### Request

- Headers: `Authorization: <token>`

- Body:
```json
{
  "number": 10
}
```


#### Response

- Code: `200 OK`
