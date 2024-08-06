package migration

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	model "app-todo/models/entity"

	"app-todo/pkg/constant"
	driver "app-todo/pkg/database"
	"app-todo/pkg/util/env"
)

type Migration interface {
	AutoMigrate()
	SetDb(*gorm.DB)
}

type migration struct {
	Db            *gorm.DB
	DbModels      *[]interface{}
	IsAutoMigrate bool
}

func Init() {
	if !env.NewEnv().GetBool(constant.MIGRATION_ENABLED) {
		return
	}

	mgConfigurations := map[string]Migration{
		constant.DB_NAME: &migration{
			DbModels: &[]interface{}{
				model.MasterTaskModel{},
				model.MasterCategoryModel{},
				model.TTaskCategoryModel{},
			},
			IsAutoMigrate: true,
		},
	}

	for k, v := range mgConfigurations {
		dbConnection, err := driver.GetConnection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("Failed to run migration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.AutoMigrate()
			logrus.Info(fmt.Sprintf("Successfully run migration for database %s", k))
		}
	}
}

func (m *migration) AutoMigrate() {
	if m.IsAutoMigrate {
		m.Db.AutoMigrate(*m.DbModels...)
	}
}

func (m *migration) SetDb(db *gorm.DB) {
	m.Db = db
}
