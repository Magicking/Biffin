package internal

import (
	"context"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"encoding/json"
	"fmt"

	"github.com/Magicking/Biffin/models"
)

type MapFileDb struct {
	gorm.Model
	Height  int64
	Width   int64
	JsonObj []byte
}

func InsertMapFile(ctx context.Context, map_file *models.MapFile2) error {
	DB := DBFromContext(ctx)
	jsonObj, err := json.Marshal(map_file)
	if err != nil {
		return err
	}
	map_file_db := MapFileDb{
		Height:  map_file.Height,
		Width:   map_file.Width,
		JsonObj: jsonObj}

	if err := DB.Create(&map_file_db).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMapFile(ctx context.Context) ([]*models.MapFile2, error) {
	DB := DBFromContext(ctx)
	var _map_files []MapFileDb

	if DB.Find(&_map_files).RecordNotFound() {
		return nil, fmt.Errorf("RecordNotFound: No map file in database")
	}

	ret := make([]*models.MapFile2, len(_map_files))
	for i, element := range _map_files {
		tmp := models.MapFile2{}
		err := json.Unmarshal(element.JsonObj, &tmp)
		if err != nil {
			return nil, err
		}
		ret[i] = &tmp
	}
	return ret, nil
}

func InitDatabase(dbDsn string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB

	for i := 1; i < 10; i++ {
		db, err = gorm.Open("postgres", dbDsn)
		if err == nil || i == 10 {
			break
		}
		sleep := (2 << uint(i)) * time.Second
		log.Printf("Could not connect to DB: %v", err)
		log.Printf("Waiting %v before retry", sleep)
		time.Sleep(sleep)
	}
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&MapFileDb{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
