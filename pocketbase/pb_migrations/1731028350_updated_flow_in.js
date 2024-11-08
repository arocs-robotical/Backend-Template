/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tcmejk6xbug3vyi")

  // remove
  collection.schema.removeField("fhxtmd6q")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ayxv9mbf",
    "name": "status",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "draft",
        "on_progress",
        "done"
      ]
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tcmejk6xbug3vyi")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "fhxtmd6q",
    "name": "status",
    "type": "text",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // remove
  collection.schema.removeField("ayxv9mbf")

  return dao.saveCollection(collection)
})
