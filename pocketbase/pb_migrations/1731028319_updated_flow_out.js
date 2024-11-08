/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("zno6j7e8y2w8u5f")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ybwl9d45",
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
  const collection = dao.findCollectionByNameOrId("zno6j7e8y2w8u5f")

  // update
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "ybwl9d45",
    "name": "status",
    "type": "select",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "maxSelect": 1,
      "values": [
        "draft",
        "on progress",
        "done"
      ]
    }
  }))

  return dao.saveCollection(collection)
})
