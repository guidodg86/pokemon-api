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
            poke_response = requests.get(target_uri)
            return poke_response.json()



api.add_resource(PokeList, '/')

if __name__ == '__main__':
    app.run(debug=True)