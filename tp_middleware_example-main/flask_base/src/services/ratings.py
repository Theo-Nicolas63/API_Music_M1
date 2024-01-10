import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.rating import RatingSchema
from src.models.user import User as UserModel
from src.models.http_exceptions import *
import src.repositories.users as users_repository


ratings_url = "http://localhost:8080/ratings/"  # URL de l'API ratings (golang)


def get_rating(id):

    response = requests.request(method="GET", url=ratings_url+id)
    return response.json(), response.status_code


def get_ratings():
    response = requests.request(method="GET", url=ratings_url)
    return response.json(), response.status_code

def post_rating(new_rating):
    rating_schema = RatingSchema().loads(json.dumps(new_rating), unknown=EXCLUDE)
    response = requests.request(method="POST", url=ratings_url, json=rating_schema)
    return response.json(), response.status_code

def delete_rating(id):
    response = requests.request(method="DELETE", url=ratings_url+id)
    return response.status_code



def modify_rating(id, raiting_update):

    # s'il y a quelque chose à changer côté API (username, name)
    rating_schema = RatingSchema().loads(json.dumps(raiting_update), unknown=EXCLUDE)
    response = None
    if not RatingSchema.is_empty(rating_schema):
        # on lance la requête de modification
        response = requests.request(method="PUT", url=ratings_url+id, json=rating_schema)
        return response.json(), response.status_code
