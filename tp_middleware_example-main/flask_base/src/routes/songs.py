import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.song import *
from src.schemas.song import BaseSongSchema
from src.schemas.errors import *
import src.services.songs as songs_service

# from routes import songs
songs = Blueprint(name="Songs", import_name=__name__)

@songs.route('/<id>', methods=['GET'])
def get_song(id):
    """
    ---
    get:
        description: Getting a song
        parameters:
            - in: path
              name: id
              schema:
                type: uuidv4
              required: true
              description: UUID of song id
        responses:
            '200':
                description: Ok
                content:
                    application/json:
                        schema: Song
                    application/yaml:
                        schema: Song
            '401':
                description: Unauthorized
                content:
                    application/json:
                        schema: Unauthorized
                    application/yaml:
                        schema: Unauthorized
            '404':
                description: Not found
                content:
                    application/json:
                        schema: NotFound
                    application/yaml:
                        schema: NotFound
        tags:
            - songs
    """
    return songs_service.get_song(id)

@songs.route('/', methods=['GET'])
def get_songs():
    """
    ---
    get:
        description: Getting all songs
        responses:
            '200':
                description: Ok
                content:
                    application/json:
                        schema: List[Song]
                    application/yaml:
                        schema: List[Song]
            '401':
                description: Unauthorized
                content:
                    application/json:
                        schema: Unauthorized
                    application/yaml:
                        schema: Unauthorized
        tags:
            - songs
    """
    return songs_service.get_songs()

@songs.route('/<id>', methods=['DELETE'])
def delete_song(id):
    """
    ---
    delete:
        description: Deleting a song
        parameters:
            - in: path
              name: id
              schema:
                type: uuidv4
              required: true
              description: UUID of song id
        responses:
            '204':
                description: No content
            '401':
                description: Unauthorized
                content:
                    application/json:
                        schema: Unauthorized
                    application/yaml:
                        schema: Unauthorized
            '404':
                description: Not found
                content:
                    application/json:
                        schema: NotFound
                    application/yaml:
                        schema: NotFound
        tags:
            - songs
    """
    return songs_service.delete_song(id)

@songs.route('', methods=['POST'])
def post_song():
    """
    ---
    post:
        description: Creating a new song
        requestBody:
            required: true
            content:
                application/json:
                    schema: BaseSongSchema
                application/yaml:
                    schema: BaseSongSchema
        responses:
            '201':
                description: Created
                content:
                    application/json:
                        schema: SongSchema
                    application/yaml:
                        schema: SongSchema
            '400':
                description: Bad request
                content:
                    application/json:
                        schema: BadRequest
                    application/yaml:
                        schema: BadRequest
            '401':
                description: Unauthorized
                content:
                    application/json:
                        schema: Unauthorized
                    application/yaml:
                        schema: Unauthorized
        tags:
            - songs
    """
    try:
        new_song = BaseSongSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    return songs_service.post_song(new_song)



@songs.route('/<id>', methods=['PUT'])
def put_song(id):
    """
    ---
    put:
      description: Updating a song
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of song id
      requestBody:
        required: true
        content:
            application/json:
                schema: SongUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Song
            application/yaml:
              schema: Song
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - songs
    """ 
 
    try:
        song_update = SongUpdateSchema().loads(json_data=request.data.decode('utf-8'))
    except ValidationError as e:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": e.messages.__str__()}))
        return error, error.get("code")

    try:
        return songs_service.modify_song(id, song_update)
    except Conflict:
        error = ConflictSchema().loads(json.dumps({"message": "song already exists"}))
        return error, error.get("code")
    except UnprocessableEntity:
        error = UnprocessableEntitySchema().loads(json.dumps({"message": "One required field was empty"}))
        return error, error.get("code")
    except Forbidden:
        error = ForbiddenSchema().loads(json.dumps({"message": "Can't manage other songs"}))
        return error, error.get("code")
    except Exception:
        error = SomethingWentWrongSchema().loads("{}")
        return error, error.get("code")
