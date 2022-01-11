from json.decoder import JSONDecodeError
from django.shortcuts import render
from rest_framework.views import APIView
from poke_responder.services import ConsumePokeData
from rest_framework.response import Response

POKE_API_URI_LIST = 'https://pokeapi.co/api/v2/pokemon/'
SEARCH_LIMIT = 1000

class GetPokeList(APIView):

    def get(self, request):
        if request.method == 'GET':
            query_params = {}
            try:
                query_params['offset'] = int(request.GET.get('offset'))
            except TypeError:
                query_params['offset'] = 0
            except ValueError:
                return Response({"error": "parameter offset must be a positive integer" })     
            try:
                query_params['limit']= int(request.GET.get('limit'))
            except TypeError:
                query_params['limit'] = 10
            except ValueError:
                return Response({"error": "parameter limit must be a positive integer" } )  
            query_params['name_to_search'] = request.GET.get('q')


            if query_params['name_to_search'] == None:
                target_uri = POKE_API_URI_LIST[:-1] + "?offset=" + str(query_params['offset']) + "&limit=" + str(query_params['limit'])
                poke_response = ConsumePokeData(target_uri)
                poke_list_result = []
                poke_list_received = poke_response['results']
                for pokemon in poke_list_received:
                    pokemon_url = pokemon['url']
                    pokemon_id = pokemon_url.split("pokemon/")[1][:-1]
                    pokemon_name = pokemon['name']
                    poke_data = ConsumePokeData(pokemon_url)
                    poke_list_result.append(dict(
                        id = pokemon_id,
                        name = pokemon_name,
                        image = poke_data['sprites']['front_default']
                    ))
                
                final_response = dict(  total=poke_response['count'],
                                        limit=query_params['limit'],
                                        offset=query_params['offset'],
                                        data =  poke_list_result )
                return Response(final_response)
            
            else:
                poke_response = ConsumePokeData(POKE_API_URI_LIST[:-1] + "?limit=" + str(SEARCH_LIMIT) )
                pokes_found =0
                poke_list_result = []
                while pokes_found<query_params['limit']:
                    poke_list_received = poke_response['results']
                    for pokemon in poke_list_received:
                        if query_params['name_to_search'] in  pokemon['name']:
                            pokes_found += 1
                            pokemon_url = pokemon['url']
                            pokemon_id = pokemon_url.split("pokemon/")[1][:-1]
                            pokemon_name = pokemon['name']
                            poke_data = ConsumePokeData(pokemon_url)
                            poke_list_result.append(dict(
                                id = pokemon_id,
                                name = pokemon_name,
                                image = poke_data['sprites']['front_default']
                            ))
                    if poke_response['next'] == None:
                        break
                    poke_response = ConsumePokeData(poke_response['next'])
                
                first_req_element = query_params['offset']
                last_req_element = query_params['offset'] + query_params['limit'] 
                final_response = dict(  total=poke_response['count'],
                                        limit=query_params['limit'],
                                        offset=query_params['offset'],
                                        data =  poke_list_result[first_req_element:last_req_element] )
                return Response(final_response)


class GetPokeDetail(APIView):

    def get(self, request, poke_id):
        if request.method == 'GET':
            try:
                poke_response = ConsumePokeData(POKE_API_URI_LIST + poke_id )
            except JSONDecodeError:
                return Response({"error": "Pokemon not found in database"})        
            abilities = [ability['ability']['name'] for ability in poke_response['abilities']]
            base_experience = poke_response['base_experience']
            forms = [form['name'] for form in  poke_response['forms']]
            height = poke_response['height']
            id = poke_response['id']
            location_area_encounters = poke_response['location_area_encounters']
            moves = [move['move']['name'] for move in  poke_response['moves']] 
            name = poke_response['name']
            order = poke_response['order']    
            species = poke_response['species']['name']       
            sprites = [poke_response['sprites'][sprite] for sprite in  poke_response['sprites'] if poke_response['sprites'][sprite] != None and sprite != "other" and sprite != "versions"] 
            stats = { stat['stat']['name']  : stat['base_stat'] for stat in poke_response['stats'] }
            types = [type_of_poke['type']['name']  for type_of_poke in  poke_response['types']]
            weight = poke_response['weight']

            final_response = dict(  
                abilities = abilities,
                base_experience = base_experience,
                forms = forms,
                height = height,
                id = id,
                location_area_encounters = location_area_encounters,
                moves = moves,
                name = name, 
                order = order,
                species = species,
                sprites = sprites,
                stats = stats,
                types = types,
                weight = weight
            )
            return Response(final_response)