package internal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"biffin/models"
	"fmt"
	"encoding/json"
)

type MapFileDb struct {
	gorm.Model
	Height int64
	Width int64
	JsonObj []byte
}

func InsertMapFile(ctx *Context, map_file *models.MapFile) error {
	jsonObj, err := json.Marshal(map_file)
	if err != nil {
		return err
	}
	map_file_db := MapFileDb{
		Height: map_file.Height,
		Width: map_file.Width,
		JsonObj: jsonObj}

	if err := ctx.Db.Create(&map_file_db).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMapFile(ctx *Context) ([]*models.MapFile, error) {
	var _map_files []MapFileDb

	if ctx.Db.Find(&_map_files).RecordNotFound() {
		return nil, fmt.Errorf("RecordNotFound: No map file in database")
	}

	ret := make([]*models.MapFile, len(_map_files))
	for i, element := range _map_files {
		tmp := models.MapFile{}
		err := json.Unmarshal(element.JsonObj, &tmp)
		if err != nil {
			return nil, err
		}
		ret[i] = &tmp
	}
	return ret, nil
}

func InitDatabase(dbDsn string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbDsn)
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&MapFileDb{}).Error; err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
