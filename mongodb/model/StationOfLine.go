package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type StationOfLine struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	LineNo        string             `bson:"LineNo" json:"LineNo"`
	LineID        string             `bson:"LineID" json:"LineID"`
	Stations      []Station          `bson:"Stations" json:"Stations"`
	SrcUpdateTime string             `bson:"SrcUpdateTime" json:"SrcUpdateTime"`
	UpdateTime    string             `bson:"UpdateTime" json:"UpdateTime"`
	VersionID     int32              `bson:"VersionID" json:"VersionID"`
}

type UnwindStationOfLine struct {
	ID            primitive.ObjectID `bson:"_id" json:"_id"`
	LineNo        string             `bson:"LineNo" json:"LineNo"`
	LineID        string             `bson:"LineID" json:"LineID"`
	Stations      Station            `bson:"Stations" json:"Stations"`
	SrcUpdateTime string             `bson:"SrcUpdateTime" json:"SrcUpdateTime"`
	UpdateTime    string             `bson:"UpdateTime" json:"UpdateTime"`
	VersionID     int32              `bson:"VersionID" json:"VersionID"`
}

type Station struct {
	Sequence    int32       `bson:"Sequence" json:"Sequence"`
	StationID   string      `bson:"StationID" json:"StationID"`
	StationName StationName `bson:"StationName" json:"StationName"`
}
