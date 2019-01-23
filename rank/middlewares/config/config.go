package config

const (
	// MongoDBHost is the URL host for MongoDB connections
	MongoDBHost = "mongodb://ferrei28:ferrei28@ds153763.mlab.com:53763/rank"

	// MongoDBDatabaseName is the name of MongoDB's database
	MongoDBDatabaseName = "rank"

	// MongoDBPool sets the connection pool for the database
	MongoDBPool = 50

	// GameCollection is the collection where games will be saved in the database
	GameCollection = "game"

	// Port is the port the server will run
	Port = ":8080"
)
