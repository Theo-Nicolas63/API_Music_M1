from marshmallow import Schema, fields, validates_schema, ValidationError


class SongSchema(Schema):
    id = fields.String(description="UUID")
    name = fields.String(description="name")
    singer = fields.String(description="singer")
    
    @staticmethod
    def is_empty(obj):
        return (not obj.get("id") or obj.get("id") == "") and \
            (not obj.get("name") or obj.get("name") == "") and \
            (not obj.get("singer") or obj.get("singer") == "") # and \
            #(not obj.get("inscription_date") or obj.get("inscription_date") == "")


class BaseSongSchema(Schema):
    name = fields.String(description="name")
    singer = fields.String(description="singer")


# Schéma songs  de modification (name, singer)
class SongUpdateSchema(BaseSongSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not(("name" in data and data["name"] != "") or 
              ("singer" in data and data["singer"] != "")):
            raise ValidationError("at least one of ['name','singer'] must be specified")
