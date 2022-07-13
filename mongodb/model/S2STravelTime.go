package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type S2STravelTime struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`

	LineNo        string       `bson:"LineNo" json:"LineNo"`
	LineID        string       `bson:"LineID" json:"LineID"`
	RouteID       string       `bson:"RouteID" json:"RouteID"`
	TrainType     int32        `bson:"TrainType" json:"TrainType"`
	TravelTimes   []TravelTime `bson:"TravelTimes" json:"TravelTimes"`
	SrcUpdateTime string       `bson:"SrcUpdateTime" json:"SrcUpdateTime"`
	UpdateTime    string       `bson:"UpdateTime" json:"UpdateTime"`
	VersionID     int32        `bson:"VersionID" json:"VersionID"`
}

type UnwindS2STravelTime struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`

	LineNo        string     `bson:"LineNo" json:"LineNo"`
	LineID        string     `bson:"LineID" json:"LineID"`
	RouteID       string     `bson:"RouteID" json:"RouteID"`
	TrainType     int32      `bson:"TrainType" json:"TrainType"`
	TravelTimes   TravelTime `bson:"TravelTimes" json:"TravelTimes"`
	SrcUpdateTime string     `bson:"SrcUpdateTime" json:"SrcUpdateTime"`
	UpdateTime    string     `bson:"UpdateTime" json:"UpdateTime"`
	VersionID     int32      `bson:"VersionID" json:"VersionID"`
}

type TravelTime struct {
	Sequence        int32       `bson:"Sequence" json:"Sequence"`
	FromStationID   string      `bson:"FromStationID" json:"FromStationID"`
	FromStationName StationName `bson:"FromStationName" json:"FromStationName"`
	ToStationID     string      `bson:"ToStationID" json:"ToStationID"`
	ToStationName   StationName `bson:"ToStationName" json:"ToStationName"`
	RunTime         int32       `bson:"RunTime" json:"RunTime"`
	StopTime        int32       `bson:"StopTime" json:"StopTime"`
}
