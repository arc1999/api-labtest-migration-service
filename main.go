package main


import "MigrationLabTest/service"

var s service.LabTestService

func main() {
	s.Migrate()
}
