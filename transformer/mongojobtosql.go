package transformer

import (
	"MigrationLabTest/model"
)

func Transform(mlabtests []model.LabTestMongo) []model.LabTest{
	var labtests []model.LabTest
	for _, mlabtest := range mlabtests {
		var labtest model.LabTest
		labtest.DateUpdated = mlabtest.DateUpdated
		labtest.ID = mlabtest.ID
		labtest.CreatedBy = mlabtest.CreatedBy
		labtest.DateCreated = mlabtest.DateCreated
		labtest.UpdatedBy = mlabtest.UpdatedBy
		labtest.LtCode = mlabtest.LtCode
		labtest.TestName=mlabtest.TestName
		labtest.AbbreviatedName=mlabtest.TestName
		labtest.LoincCode=mlabtest.LoincCode
		labtest.TypeOfTest=mlabtest.TypeOfTest
		labtests = append(labtests, labtest)
	}
	return labtests
}
