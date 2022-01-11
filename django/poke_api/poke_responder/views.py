from django.shortcuts import render
from rest_framework.views import APIView
from poke_responder.services import ConsumePokeData
from rest_framework.response import Response

class GetPokeData(APIView):

    def get(self, request):
        if request.method == 'GET':
            poke_data=ConsumePokeData()
            return Response(poke_data)