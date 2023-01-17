package model
import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	cont "lecture/WBA-BC-Project-04/contracts"
	"math/big"

)

type Room struct {
	RoomNo     int    `json:"roomNo" bson:"roomNo"`
	Title      string `json:"title" bson:"title"`
	Maker	   string `json:"address" bson:"address"`
	MatchPrice int    `json:"matchPrice" bson:"matchPrice"`
}
