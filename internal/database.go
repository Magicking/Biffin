package internal

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"biffin/models"
	"fmt"
)

type MapFileDb struct {
	gorm.Model
	Obj *models.MapFile
}

func InsertMapFile(ctx *Context, map_file *models.MapFile) error {
	_map_file := MapFileDb{Obj: map_file}

	if err := ctx.Db.Create(&_map_file).Error; err != nil {
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
		ret[i] = element.Obj
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
