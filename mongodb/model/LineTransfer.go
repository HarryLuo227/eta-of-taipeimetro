package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type LineTransfer struct {
	ID primitive.ObjectID `bson:"_id" json:"_id"`

	FromLineNo      string      `bson:"FromLineNo" json:"FromLineNo"`
	FromLineID      string      `bson:"FromLineID" json:"FromLineID"`
	FromLineName    LineName    `bson:"FromLineName" json:"FromLineName"`
	FromStationID   string      `bson:"FromStationID" json:"FromStationID"`
	FromStationName StationName `bson:"FromStationName" json:"FromStationName"`

	ToLineNo      string      `bson:"ToLineNo" json:"ToLineNo"`
	ToLineID      string      `bson:"ToLineID" json:"ToLineID"`
	ToLineName    LineName    `bson:"ToLineName" json:"ToLineName"`
	ToStationID   string      `bson:"ToStationID" json:"ToStationID"`
	ToStationName StationName `bson:"ToStationName" json:"ToStationName"`

	IsOnSiteTransfer    int32  `bson:"IsOnSiteTransfer" json:"IsOnSiteTransfer"`
	TransferTime        int32  `bson:"TransferTime" json:"TransferTime"`
	TransferDescription string `bson:"TransferDescription" json:"TransferDescription"`
	SrcUpdateTime       string `bson:"SrcUpdateTime" json:"SrcUpdateTime"`
	UpdateTime          string `bson:"UpdateTime" json:"UpdateTime"`
	VersionID           int32  `bson:"VersionID" json:"VersionID"`
}
