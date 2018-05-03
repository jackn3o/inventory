package items

import (
	"net/http"
	"time"

	utility "../../base/utilities"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// ItemDetail model for items.details Collection
type ItemDetail struct {
	ID           bson.ObjectId `bson:"id,omitempty" json:"id,omitempty" valid:"-"`
	OutletID     bson.ObjectId `bson:"outletId,omitempty" json:"outletId,omitempty" valid:"-"`
	DocumentNo   string        `bson:"documentNo,omitempty" json:"documentNo,omitempty" valid:"-"`
	BatchNo      string        `bson:"batchNo,omitempty" json:"batchNo,omitempty" valid:"-"`
	Date         *time.Time    `bson:"date,omitempty" json:"date,omitempty" valid:"-"`
	In           int           `bson:"in" json:"in,string" valid:"-"`
	Out          int           `bson:"out" json:"out,string" valid:"-"`
	UnitPrice    float64       `bson:"unitPrice" json:"unitPrice,string" valid:"-"`
	TotalCost    float64       `bson:"totalCost,omitempty" json:"totalCost,string,omitempty" valid:"-"`
	TotalSales   float64       `bson:"totalSales,omitempty" json:"totalSales,string,omitempty" valid:"-"` // only for out
	CreatedDate  *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy    string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy   string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
}

// DetailResponseDto model
type DetailResponseDto struct {
	Code        string        `bson:"code" json:"code" valid:"required"`
	Description string        `bson:"description" json:"description" valid:"required"`
	Details     []*ItemDetail `bson:"details,omitempty" json:"details,omitempty" valid:"-"`
}

// CostAndBalance model
type CostAndBalance struct {
	UnitCost        float64 `bson:"unitCost" json:"unitCost" valid:"-"`
	BalanceQuantity int     `bson:"balanceQuantity" json:"balanceQuantity" valid:"-"`
	BalanceCost     float64 `bson:"balanceCost" json:"balanceCost" valid:"-"`
}

// AddDetail to item
func (c *Controller) AddDetail() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		itemID := mux.Vars(req)["id"]
		if len(itemID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		detail := &ItemDetail{}
		err := u.UnmarshalWithValidation(detail)
		if err != nil {
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ItemsCollection)

		target := bson.M{"_id": bson.ObjectIdHex(itemID)}
		data := bson.M{"$push": bson.M{"details": detail}}

		err = collection.Update(target, data)
		if err != nil {
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

		var costAndBalance CostAndBalance
		err = collection.
			FindId(bson.ObjectIdHex(itemID)).
			Select(bson.M{"unitCost": 1, "balance": 1}).
			One(&costAndBalance)

		lastestCostAndBalance := calculateCostAndBalance(costAndBalance, detail)

		u.WriteJSON("success", http.StatusCreated)
	})
}

// GetItemDetailsByID ...
func (c *Controller) GetItemDetailsByID() http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		u := utility.New(writer, req)

		itemID := mux.Vars(req)["id"]

		if len(itemID) <= 0 {
			u.WriteJSONError("ID is require", http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		var dto DetailResponseDto
		collection := session.DB(c.databaseName).C(ItemsCollection)
		err := collection.
			FindId(bson.ObjectIdHex(itemID)).
			Select(bson.M{"code": 1, "description": 1, "details": 1}).
			One(&dto)

		if err != nil {
			u.WriteJSONError("Something wrong, please try again later", http.StatusInternalServerError)
			return
		}

		u.WriteJSON(dto)
	})
}

func calculateCostAndBalance(costAndBalance CostAndBalance, detail *ItemDetail) CostAndBalance {

	lastUnitCost := costAndBalance.UnitCost
	lastBalanceQuantity := costAndBalance.BalanceQuantity
	lastBalanceCost := costAndBalance.BalanceCost

	currentQuantity := 0
	isIn := detail.In > 0

	if isIn {
		currentQuantity = detail.In
		detail.TotalCost = detail.UnitPrice * float64(currentQuantity)

	} else {
		currentQuantity = detail.Out
		detail.TotalCost = lastUnitCost * float64(currentQuantity)
		detail.TotalSales = detail.UnitPrice * float64(currentQuantity)

	}

	lastestCostAndBalance := CostAndBalance{
		UnitCost:        lastBalanceCost / float64(lastBalanceQuantity),
		BalanceCost:     lastBalanceCost + detail.TotalCost,
		BalanceQuantity: lastBalanceQuantity + currentQuantity,
	}

	return lastestCostAndBalance
}
