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
    ...
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
    ...
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
  "game": {
    "id": "<game id>",
    "status": 0,
    "startTime": "0001-01-01T00:00:00Z"
  },
  "player": {
    "name": "<name>",
    "admin": true,
    "money": 40,
    "country": {
      "id": 1,
      "name": "China",
      "initialMoney": 40,
      "coal": {
        "capacity": 2,
        "scalar": 1
      },
      "oil": {
        "capacity": 2,
        "scalar": 1
      },
      "gas": {
        "capacity": 2,
        "scalar": 1
      },
      "nuclear": {
        "capacity": 2,
        "scalar": 1
      },
      "geothermal": {
        "capacity": 2,
        "scalar": 1
      },
      "solar": {
        "capacity": -1,
        "scalar": 1
      },
      "wind": {
        "capacity": -1,
        "scalar": 1
      },
      "hydroelectric": {
        "capacity": 2,
        "scalar": 1
      }
    },
    "coal": {
      "owned": 0,
      "operational": 0
    },
    "oil": {
      "owned": 0,
      "operational": 0
    },
    "gas": {
      "owned": 0,
      "operational": 0
    },
    "nuclear": {
      "owned": 0,
      "operational": 0
    },
    "geothermal": {
      "owned": 0,
      "operational": 0
    },
    "solar": {
      "owned": 0,
      "operational": 0
    },
    "wind": {
      "owned": 0,
      "operational": 0
    },
    "hydroelectric": {
      "owned": 0,
      "operational": 0
    }
  }
}
```

### Purchase One Installation

- `POST /api/purchase/<resource id>`

#### Request

- Headers: `Authorization: <token>`

#### Response

- Code: `200 OK`

### Set number of installations operational

- `POST /api/operational/<resource id>`

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

### Start Game (Admin only)

- `POST /api/operational/start`

#### Request

- Headers: `Authorization: <token>`

#### Response

- Code: `200 OK`
