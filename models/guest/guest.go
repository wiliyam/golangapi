package guest



collection := "guest"







func InsertOne(params struct) struct  {
	mongodb.GetDb().Collection(collection)
}
