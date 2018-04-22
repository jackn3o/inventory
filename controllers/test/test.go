package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-mgo/mgo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2-unstable/bson"
)

/* We will first create a new type called Product
   This type will contain information about VR experiences */
type Product struct {
	Id          int
	Name        string
	Slug        string
	Description string
}

/* We will create our catalog of VR experiences and store them in a slice. */
var products = []Product{
	Product{Id: 1, Name: "Hover Shooters", Slug: "hover-shooters", Description: "Shoot your way to the top on 14 different hoverboards"},
	Product{Id: 2, Name: "Ocean Explorer", Slug: "ocean-explorer", Description: "Explore the depths of the sea in this one of a kind underwater experience"},
	Product{Id: 3, Name: "Dinosaur Park", Slug: "dinosaur-park", Description: "Go back 65 million years in the past and ride a T-Rex"},
	Product{Id: 4, Name: "Cars VR", Slug: "cars-vr", Description: "Get behind the wheel of the fastest cars in the world."},
	Product{Id: 5, Name: "Robin Hood", Slug: "robin-hood", Description: "Pick up the bow and arrow and master the art of archery"},
	Product{Id: 6, Name: "Real World VR", Slug: "real-world-vr", Description: "Explore the seven wonders of the world in VR"},
}

/* The status handler will be invoked when the user calls the /status route
   It will simply return a string with the message "API is up and running" */
var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

/* The products handler will be called when the user makes a GET request to the /products endpoint.
   This handler will return a list of products available for users to review */
var ProductsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Here we are converting the slice of products to json
	payload, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

/* The feedback handler will add either positive or negative feedback to the product
   We would normally save this data to the database - but for this demo we'll fake it
   so that as long as the request is successful and we can match a product to our catalog of products
   we'll return an OK status. */
var AddFeedbackHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, p := range products {
		if p.Slug == slug {
			product = p
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if product.Slug != "" {
		payload, _ := json.Marshal(product)
		w.Write([]byte(payload))
	} else {
		w.Write([]byte("Product Not Found"))
	}
})

//mongo

type Game struct {
	Winner       string    `bson:"winner"`
	OfficialGame bool      `bson:"official_game"`
	Location     string    `bson:"location"`
	StartTime    time.Time `bson:"start"`
	EndTime      time.Time `bson:"end"`
	Players      []Player  `bson:"players"`
}

type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
	return Player{
		Name:   name,
		Decks:  [2]string{firstDeck, secondDeck},
		Points: points,
		Place:  place,
	}
}

var test = func() {
	Host := []string{
		"localhost:27017",
		// replica set addrs...
	}
	const (
		Username   = "sa"
		Password   = "p@ssw0rd"
		Database   = "inventory"
		Collection = "test"
	)
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: Username,
		Password: Password,
		Database: Database,
	})
	if err != nil {
		panic(err)
	}
	defer session.Close()

	game := Game{
		Winner:       "Dave",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Now(),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}

	if isDropMe {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection
	c := session.DB(Database).C(Collection)

	// Insert
	if err := c.Insert(game); err != nil {
		panic(err)
	}

	// Find and Count
	player := "Dave"
	gamesWon, err := c.Find(bson.M{"winner": player}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", player, gamesWon)

	// Find One (with Projection)
	var result Game
	err = c.Find(bson.M{"winner": player, "location": "Austin"}).Select(bson.M{"official_game": 1}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Is game in Austin Official?", result.OfficialGame)

	// Find All
	var games []Game
	err = c.Find(nil).Sort("-start").All(&games)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of Games", len(games))

	// Update
	newPlayer := "John"
	selector := bson.M{"winner": player}
	updator := bson.M{"$set": bson.M{"winner": newPlayer}}
	if err := c.Update(selector, updator); err != nil {
		panic(err)
	}

	// Update All
	info, err := c.UpdateAll(selector, updator)
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated", info.Updated)
}

// // Remove
// info, err = c.RemoveAll(bson.M{"winner": newPlayer})
// if err != nil {
// 	panic(err)
// }
// fmt.Println("Removed", info.Removed)
