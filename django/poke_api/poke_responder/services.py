import requests
import json

def ConsumePokeData(uri):
    poke_response = requests.get(uri)
    return json.loads(poke_response.text)