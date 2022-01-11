import requests
import json

def ConsumePokeData():
    poke_response = requests.get('https://pokeapi.co/api/v2/pokemon/20/')
    return json.loads(poke_response.text)