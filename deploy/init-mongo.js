b = db.getSiblingDB('FunDatabase');

// Indexes creation
db.getCollection('events').createIndex({ starts_at: 1 });
db.getCollection('events').createIndex({ ends_at: 1 });
db.getCollection('events').createIndex({ prov_id: 1 }, { unique: true });
	  