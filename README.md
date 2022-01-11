# pokemon-api
Basic pokemon api

## Goal
Using PokeApi (https://pokeapi.co/) develop the following endpoints:
#### GET / -> Pokemon list
**Parameters**  
*q:* pokemon look up (must allow partial name)  
*offset:* offset from start  
*limit:* amount of pokes to deliver on the response  

**Response**
```json
{
    "total": 1118,
    "limit": 2,
    "offset": 0,
    "data": [
        {
            "id": "1",
            "name": "bulbasaur",
            "image": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"
        },
        {
            "id": "2",
            "name": "ivysaur",
            "image": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/2.png"
        }
    ]
}
```

#### GET /{pokemonName}
**Response**
```json
{
    "abilities": [
        "limber",
        "imposter"
    ],
    "base_experience": 101,
    "forms": [
        "ditto"
    ],
    "height": 3,
    "id": 132,
    "location_area_encounters": "https://pokeapi.co/api/v2/pokemon/132/encounters",
    "moves": [
        "transform"
    ],
    "name": "ditto",
    "order": 203,
    "species": "ditto",
    "sprites": [
        "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/132.png",
        "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/shiny/132.png",
        "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png",
        "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/shiny/132.png"
    ],
    "stats": {
        "hp": 48,
        "attack": 48,
        "defense": 48,
        "special-attack": 48,
        "special-defense": 48,
        "speed": 48
    },
    "types": [
        "normal"
    ],
    "weight": 40
}
```

## Requeriments - Flask
aniso8601==9.0.1  
certifi==2021.5.30  
charset-normalizer==2.0.4  
click==8.0.1  
Flask==2.0.1  
Flask-RESTful==0.3.9  
idna==3.2  
itsdangerous==2.0.1  
Jinja2==3.0.1  
MarkupSafe==2.0.1  
pytz==2021.1  
requests==2.26.0  
six==1.16.0  
urllib3==1.26.6  
Werkzeug==2.0.1  

## How to use? - Flask
From linux bash. Also can be run from REPL
```
python poke-api.py
```
After that you can start making api get request to 127.0.0.1:5000

## Requeriments - Django
asgiref==3.4.1
certifi==2021.10.8
charset-normalizer==2.0.10
Django==3.2.7
djangorestframework==3.12.4
idna==3.3
python-decouple==3.5
pytz==2021.1
requests==2.27.1
sqlparse==0.4.1
urllib3==1.26.8

## How to use? - Django
From linux bash
```
python ./manage.py runserver
```
After that you can start making api get request to 127.0.0.1:8000