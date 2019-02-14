package model

import (
	"time"
)

type NameTag struct {
	Name string `bson:"name" json:"name"`
	//Desc string `bson:"desc,omitempty" json:"desc,omitempty"`
}

type Props map[string]interface{}

type SetupMeta struct {
	NameTag  `bson:",inline"`
	Type     string    `bson:"type" json:"type"`
	Notes    string    `bson:"notes" json:"notes"`
	Status   string    `bson:"status" json:"status"`
	Since    string    `bson:"since" json:"since"`
	Props    Props     `bson:"props" json:"props"`
	Created  time.Time `bson:"created" json:"created"`
	Modified time.Time `bson:"modified" json:"modified"`
}

type CSetup struct {
	SetupMeta  `bson:"meta" json:"meta"`
	Client     Client       `bson:"client" json:"client"`
	Contacts   []*Contact   `bson:"contacts" json:"contacts"`
	Components []*Component `bson:"components" json:"components"`
}

type Client struct {
	NameTag `bson:",inline"`
}

type Contact struct {
	NameTag `bson:",inline"`
	Email   string `bson:"email" json:"email"`
	Uname   string `bson:"uname" json:"uname"`
}

type Component struct {
	NameTag `bson:",inline"`
	Type    string `bson:"type" json:"type"`
	Version string `bson:"version" json:"version"`
	Props   Props  `bson:"props" json:"props"`
	Host    struct {
		Type  string `bson:"type" json:"type"`
		Props Props  `bson:"props" json:"props"`
	} `bson:"host" json:"host"`
}
