package core

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type SignDao struct {
	mongo      *Mongo
	collection *mgo.Collection
}

func NewSignDao(m *Mongo) *SignDao {
	return &SignDao{m, m.GetCollection(C_SIGN)}
}

func (sDao *SignDao) FindSignByDate(signName string, date string) (*Sign, error) {
	result := sDao.collection.Find(bson.M{"date": date, "name": signName})
	var s Sign
	err := result.One(&s)

	return &s, err
}

func (sDao *SignDao) Save(sign *Sign) error {
	return sDao.collection.Insert(sign)
}
