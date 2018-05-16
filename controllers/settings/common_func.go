package settings

import "gopkg.in/mgo.v2/bson"

func (c *Controller) isKeyFieldsExist(id string, collectionName string, requestCode string, requestDescription string) (bool, error) {
	session := c.store.DB.Copy()
	defer session.Close()

	collection := session.DB(c.databaseName).C(collectionName)

	var selector bson.M
	if len(id) > 0 {
		selector = bson.M{
			"_id": bson.M{"$ne": bson.ObjectIdHex(id)},
			"$or": []bson.M{
				bson.M{"code": bson.RegEx{Pattern: requestCode, Options: "i"}},
				bson.M{"description": bson.RegEx{Pattern: requestDescription, Options: "i"}},
			}}
	} else {
		selector = bson.M{
			"$or": []bson.M{
				bson.M{"code": bson.RegEx{Pattern: requestCode, Options: "i"}},
				bson.M{"description": bson.RegEx{Pattern: requestDescription, Options: "i"}},
			}}
	}

	count, err := collection.Find(selector).Count()
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
