package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LabTestMongo struct {
	ID              int64     `json:"id"  bson:"_id"`
	TestName        string    `json:"testName"`
	LoincCode       string    `json:"loincCode"`
	Slug 			primitive.Binary `json:"slug"`
	TypeOfTest      string    `json:"typeOfTest"`
	DateCreated     time.Time `json:"dateCreated"`
	DateUpdated     time.Time `json:"dateUpdated"`
	CreatedBy       int64     `json:"createdBy"`
	UpdatedBy       int64     `json:"updatedBy"`
	LtCode          string    `json:"ltCode"`
	AbbreviatedName string    `json:"abbreviatedName"`
}
