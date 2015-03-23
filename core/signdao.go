package core

import (
	"labix.org/v2/mgo"
)

type SignDao struct {
	mongo *Mongo
	collection *mgo.Collection
}

func NewSignDao(m *Mongo) *SignDao{
	return &SignDao{m, m.GetCollection(C_SIGN)}
}

func findSignByDate(date string) (*Sign, )
