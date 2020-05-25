package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GusetUserModel struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	DeviceId   string             `json:"deviceId"`
	IpAddress  string             `json:"ipAddress"`
	DeviceType string             `json:deviceType`
	LoginAt    time.Time          `json:loginAt`
}
