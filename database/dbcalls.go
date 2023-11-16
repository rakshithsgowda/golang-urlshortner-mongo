package database

import (
	"context"
	"url_shortner/constant"
	"url_shortner/types"

	"gopkg.in/mgo.v2/bson"
)

func (mgr *manager) Insert(data interface{}, collectionName string) (interface{}, error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	result, err := inst.InsertOne(context.TODO(), data)
	return result.InsertedID, err
}

func (mgr *manager) GetUrlFromCode(code string, collectionName string) (resp types.URLDB, err error) {
	inst := mgr.connection.Database(constant.Database).Collection(collectionName)
	err = inst.FindOne(context.TODO(), bson.M{"url_code": code}).Decode(&resp)
	return resp, err
}
