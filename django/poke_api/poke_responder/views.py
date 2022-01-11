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