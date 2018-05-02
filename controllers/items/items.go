package items

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ItemDetail model for items.details Collection
type ItemDetail struct {
	ID              bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty" valid:"-"`
	ItemID          bson.ObjectId `bson:"itemId,omitempty" json:"itemId,omitempty" valid:"-"`
	Outlet          string        `bson:"outlet" json:"outlet" valid:"-"`
	Category        string        `bson:"category,omitempty" json:"category" valid:"-"`
	DocumentNo      string        `bson:"documentNo" json:"documentNo" valid:"-"`
	BatchNo         string        `bson:"batchNo" json:"batchNo" valid:"-"`
	Date            *time.Time    `bson:"date,omitempty" json:"date,omitempty" valid:"-"`
	Cost            float64       `bson:"cost,omitempty" json:"cost,string,omitempty" valid:"-"`
	In              float32       `bson:"in,omitempty" json:"in,string,omitempty" valid:"-"`
	Out             float32       `bson:"out,omitempty" json:"out,string,omitempty" valid:"-"`
	BalanceQuantity float32       `bson:"balanceQuantity,omitempty" json:"balanceQuantity" valid:"-"`
	CreatedDate     *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy       string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate    *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy      string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}
