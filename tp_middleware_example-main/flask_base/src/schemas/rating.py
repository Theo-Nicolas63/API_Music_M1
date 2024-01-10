from marshmallow import Schema, fields, validates_schema, ValidationError


# Schéma utilisateur de sortie (renvoyé au front)
class RatingSchema(Schema):
    id = fields.String(description="UUID")
    song_id = fields.String(description="UUID")
    user_id = fields.String(description="UUID")
    content = fields.String(description="content")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("song_id") or obj.get("song_id") == "") and \
               (not obj.get("user_id") or obj.get("user_id") == "") #and \
               #(not obj.get("inscription_date") or obj.get("inscription_date") == "")


class BaseRatingSchema(Schema):
    song_id = fields.String(description="UUID")
    user_id = fields.String(description="UUID")
    content = fields.String(description="content")


# Schéma rating de modification (content)
class RatingUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou non
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("content" in data and data["content"] != "")):
            raise ValidationError("content must be specified")
