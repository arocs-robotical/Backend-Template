/// <reference path="../pb_data/types.d.ts" />
migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tcmejk6xbug3vyi")

  // remove
  collection.schema.removeField("sqifmkrq")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "vgby5vkc",
    "name": "product_in",
    "type": "relation",
    "required": false,
    "presentable": false,
    "unique": false,
    "options": {
      "collectionId": "easzhp7xsm349tu",
      "cascadeDelete": false,
      "minSelect": null,
      "maxSelect": null,
      "displayFields": null
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("tcmejk6xbug3vyi")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "sqifmkrq",
    "name": "product_in",
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
  collection.schema.removeField("vgby5vkc")

  return dao.saveCollection(collection)
})
