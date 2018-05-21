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
	ID              bson.ObjectId `bson:"id,omitempty" json:"id,omitempty" valid:"-"`
	OutletID        bson.ObjectId `bson:"outletId,omitempty" json:"outletId,omitempty" valid:"-"`
	DocumentNo      string        `bson:"documentNo,omitempty" json:"documentNo,omitempty" valid:"-"`
	BatchNo         string        `bson:"batchNo,omitempty" json:"batchNo,omitempty" valid:"-"`
	Date            *time.Time    `bson:"date,omitempty" json:"date,omitempty" valid:"-"`
	In              int           `bson:"in,omitempty" json:"in,string" valid:"-"`
	UnitCost        float64       `bson:"unitCost,omitempty" json:"unitCost,string" valid:"-"`
	TotalCost       float64       `bson:"totalCost,omitempty" json:"totalCost,string,omitempty" valid:"-"`
	Out             int           `bson:"out,omitempty" json:"out" valid:"-"`
	SellingPrice    float64       `bson:"sellingPrice,omitempty" json:"sellingPrice" valid:"-"`              // only for out
	TotalSales      float64       `bson:"totalSales,omitempty" json:"totalSales,string,omitempty" valid:"-"` // only for out
	UpToDateBalance int           `bson:"utdBalance,omitempty" json:"utdBalance,string,omitempty" valid:"-"`
	UpToDateCost    float64       `bson:"utdCost,omitempty" json:"utdCost,string,omitempty" valid:"-"`
	Remark          string        `bson:"remark,omitempty" json:"remark,omitempty" valid:"-"`
	CreatedDate     *time.Time    `bson:"createdDate,omitempty" json:"createdDate,omitempty" valid:"-"`
	CreatedBy       string        `bson:"createBy,omitempty" json:"createBy,omitempty" valid:"-"`
	ModifiedDate    *time.Time    `bson:"modifiedDate,omitempty" json:"modifiedDate,omitempty" valid:"-"`
	ModifiedBy      string        `bson:"modifiedBy,omitempty" json:"modifiedBy,omitempty" valid:"-"`
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
			c.logger.Warn(err)
			u.WriteJSONError(err, http.StatusBadRequest)
			return
		}

		session := c.store.DB.Copy()
		defer session.Close()

		collection := session.DB(c.databaseName).C(ItemsCollection)
		var costAndBalance CostAndBalance
		err = collection.
			FindId(bson.ObjectIdHex(itemID)).
			Select(bson.M{"unitCost": 1, "balanceQuantity": 1, "balanceCost": 1}).
			One(&costAndBalance)

		lastestCostAndBalance := calculateCostAndBalance(costAndBalance, detail)

		timeNow := time.Now()
		selector := bson.M{"_id": bson.ObjectIdHex(itemID)}
		updator := bson.M{
			"$set": bson.M{
				"modifiedBy":      "todo",
				"modifiedDate":    &timeNow,
				"unitCost":        lastestCostAndBalance.UnitCost,
				"balanceQuantity": lastestCostAndBalance.BalanceQuantity,
				"balanceCost":     lastestCostAndBalance.BalanceCost},
			"$push": bson.M{"details": detail},
		}

		err = collection.Update(selector, updator)
		if err != nil {
			c.logger.Error(err)
			u.WriteJSONError("Something Wrong, Please try again later", http.StatusInternalServerError)
			return
		}

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
			Select(bson.M{"code": 1, "description": 1, "balanceQuantity": 1, "balanceCost": 1, "details": 1}).
			One(&dto)

		if err != nil {
			c.logger.Error(err)
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
		detail.TotalCost = detail.UnitCost * float64(currentQuantity)

	} else {
		currentQuantity = detail.Out
		detail.TotalCost = lastUnitCost * float64(currentQuantity)
		detail.TotalSales = detail.SellingPrice * float64(currentQuantity)
	}

	detail.UpToDateCost = lastBalanceCost + detail.TotalCost
	detail.UpToDateBalance = lastBalanceQuantity + currentQuantity
	uptoDateCostAndBalance := CostAndBalance{
		UnitCost:        detail.UpToDateCost / float64(detail.UpToDateBalance),
		BalanceCost:     detail.UpToDateCost,
		BalanceQuantity: detail.UpToDateBalance,
	}

	return uptoDateCostAndBalance
}
