from flask import Flask
from flask_restful import request ,Resource, Api
import requests

POKE_API_URI_LIST = 'https://pokeapi.co/api/v2/pokemon/'

app = Flask(__name__)
api = Api(app)


class PokeList(Resource):
    def get(self):
        query_params = {}
        query_params['offset'] = request.args.get('offset')
        query_params['limit'] = request.args.get('limit')
        name_to_search = request.args.get('q')

        if len(query_params):
            target_uri = POKE_API_URI_LIST[:-1] + "?"
            first = True
            for param in query_params:
                if first:
                    first = False
                    target_uri = target_uri + param + "=" + query_params[param]
                else:
                    target_uri = target_uri + "&" + param + "=" +  query_params[param]                 
        else:
            target_uri = POKE_API_URI_LIST

        if name_to_search == None:
            poke_response = requests.get(target_uri).json()

            poke_list_result = []
            poke_list_received = poke_response['results']
            for pokemon in poke_list_received:
                pokemon_url = pokemon['url']
                pokemon_id = pokemon_url.split("pokemon/")[1][:-1]
                pokemon_name = pokemon['name']
                poke_data = requests.get(pokemon_url).json()
                poke_list_result.append(dict(
                    id = pokemon_id,
                    name = pokemon_name,
                    image = poke_data['sprites']['front_default']
                ))
            
            print(target_uri)
            final_response = dict(  total=poke_response['count'],
                                    limit=query_params['limit'],
                                    offset=query_params['offset'],
                                    data =  poke_list_result )
            return final_response



api.add_resource(PokeList, '/')

if __name__ == '__main__':
    app.run(debug=True)