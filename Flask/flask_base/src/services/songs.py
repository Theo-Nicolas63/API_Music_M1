import json
import requests
from flask import jsonify
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user

from src.schemas.user import UserSchema
from src.schemas.song import SongSchema
from src.schemas.song import SongUpdateSchema
from src.models.http_exceptions import *

songs_url = "http://localhost:8081/Songs/"

def get_song(id):
    response = requests.request(method="GET", url=songs_url+id)
    return response.json(), response.status_code

def get_songs():
    response = requests.request(method="GET", url=songs_url)
    return response.json(), response.status_code

def delete_song(id):
    response = requests.request(method="DELETE", url=songs_url+id)
    if response.status_code != 204:
        return jsonify({'message': 'Song deleted successfully'}), response.status_code
    else:
        return jsonify({'error': 'Failed to delete song'}), 204 

def post_song(new_song):
    song_schema = SongSchema().loads(json.dumps(new_song), unknown=EXCLUDE)
    response = requests.request(method="POST", url=songs_url, json=song_schema)
    return response.json(), response.status_code

def modify_song(id, updated_song):
    song_schema = SongSchema().loads(json.dumps(updated_song, default=str), unknown=EXCLUDE)
    print(song_schema)
    response = requests.put(songs_url+id, json=updated_song)
    print(response.status_code)
    if response.status_code != 200:
        return jsonify({'error': 'Failed to update song'}), response.status_code
    else:
        return jsonify({'message': 'Song updated successfully'}), 200



