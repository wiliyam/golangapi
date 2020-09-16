package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GusetUserModelPost struct {
	// ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`\
	DeviceTypeMsg string    `json:"deviceTypeMsg"`
	DeviceId      string    `json:"deviceId"`
	IpAddress     string    `json:"ipAddress"`
	DeviceType    int       `json:deviceType`
	LoginAt       time.Time `json:loginAt`
}
type GusetUserModelGet struct {
	ID            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	DeviceId      string             `json:"deviceId"`
	IpAddress     string             `json:"ipAddress"`
	DeviceTypeMsg string             `json:"deviceTypeMsg"`
	DeviceType    int                `json:deviceType`
	LoginAt       time.Time          `json:loginAt`
}
