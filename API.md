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

- `POST /api/join`

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

### Get Rankings (from highest to lowest)

- `POST /api/rankings/<gameid>`

#### Response

- Code: `200 OK`

- Body:
```json
[
  {
    ... // player object
  }
]
```
## Resource IDs

- Coal: 0
- Oil: 1
- Gas: 2
- Nuclear: 3
- Geothermal: 4
- Solar: 5
- Wind: 6
- Hydroelectric: 7

## API errors

Errors are returned in the format `{"errorID": <id>}`

### Error IDs and meanings

- 0: Invalid username chose
- 1: Invalid authorization token
- 2: Invalid game ID
- 3: No room left in game
- 4: Invalid Resource ID
- 5: Attempted purchase was greater than players money
- 6: Player is already at installation capacity for the Resource
- 7: Invalid number of operational units specified
- 8: Admin permissions required
- 9: The game has already started
