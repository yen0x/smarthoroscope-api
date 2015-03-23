package core

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type SignDao struct {
	mongo *Mongo
	collection *mgo.Collection
}

func NewSignDao(m *Mongo) *SignDao{
	return &SignDao{m, m.GetCollection(C_SIGN)}
}

func (sDao *SignDao)findSignByDate(signName string, date string) (*Sign, error) {
	//TODO check date & name
	s := sDao.collection.Find(bson.M{"date": date, "name" : signName})
	return &s, nil
}
