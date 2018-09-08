package settings

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Application is model for settings.application Collection
type Application struct {
	ID           bson.ObjectId `bson:"_id,omitempty" json:"_id" valid:"-"`
	LockedDate   *time.Time    `bson:"code" json:"code" valid:"required"`
	CreatedDate  *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy    string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}
