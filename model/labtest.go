package model

import "time"

type LabTest struct {
	ID              int64     `json:"id" gorm:"column:id;type:bigint;"`
	TestName        string    `json:"testName"`
	LoincCode       string    `json:"loincCode"`
	Description     string    `json:"description"`
	TypeOfTest      string    `json:"typeOfTest"`
	DateCreated     time.Time `json:"dateCreated"`
	DateUpdated     time.Time `json:"dateUpdated"`
	CreatedBy       int64     `json:"createdBy"`
	UpdatedBy       int64     `json:"updatedBy"`
	LtCode          string    `json:"ltCode"`
	AbbreviatedName string    `json:"abbreviatedName"`
	Instruction     string    `json:"instruction"`
}
