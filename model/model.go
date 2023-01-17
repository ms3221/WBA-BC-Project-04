package model

import (
	"context"
	"fmt"
	conf "lecture/WBA-BC-Project-04/conf"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	client  *mongo.Client
	gameCol *mongo.Collection
	game    Game
	daemon  Daemon
}
type Game struct {
	privateKey      string
	netUrl          string
	ownerAddress    string
	contractAddress string
}

type Daemon struct {
	url string
}

func NewModel(cfg *conf.Config) (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := cfg.Database.Host
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(cfg.Database.DB)
		r.gameCol = db.Collection(cfg.Database.GameCollection)
		fmt.Println(db)
		// r.colPersons = db.Collection("tPerson")
		// r.colMenus = db.Collection("tMenu")
		// r.colOrders = db.Collection("tOrder")
	}

	r.game.privateKey = cfg.Contract.PrivateKey
	r.game.netUrl = cfg.Contract.NetUrl
	r.game.ownerAddress = cfg.Contract.OwnerAddress
	r.game.contractAddress = cfg.Contract.ContractAddress
	r.daemon.url = cfg.Daemon.Url

	return r, nil
}
