package dao

import (
	"MigrationLabTest/db"
	"MigrationLabTest/model"
	"context"
	log "github.com/sirupsen/logrus"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type LabTestDao struct {
}

func (d LabTestDao) Paginate(pagenumber int64, nperpage int64) ([]model.LabTestMongo, error) {
	log.Println(pagenumber,nperpage)
	findOptions := options.Find()
	findOptions.SetLimit(nperpage)
	findOptions.SetSort(bson.M{})
	findOptions.SetSkip(pagenumber)

	db := db.GetMongoDB()
	cur, err := db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	var jobs []model.LabTestMongo
	for cur.Next(context.TODO()) {
		var job model.LabTestMongo
		err := cur.Decode(&job)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, job)
	}
	log.Println(len(jobs))
	return jobs, nil
}
func (d LabTestDao) GetCount() (int64, error) {
	db := db.GetMongoDB()
	return db.Collection(os.Getenv("DATA_MONGODB_COLLECTION")).CountDocuments(context.TODO(),bson.D{})
}
func (d LabTestDao) BulkInsert(Entity []model.LabTest, nperpage int64) error {
	sqldb := db.GetMysqlDB()
	b := make([]interface{}, len(Entity))
	for i := range Entity {
		b[i] = Entity[i]
	}
	err := gormbulk.BulkInsert(sqldb, b, int(nperpage))
	if err != nil {
		log.Printf("error in saving labtest")
		return err
	}
	return nil

}
