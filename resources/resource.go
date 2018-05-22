package resources

import (
	"github.com/NghiaTranUIT/artify-core/constant"
	"github.com/NghiaTranUIT/artify-core/model"
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Resource struct {
	Config     constant.Config
	PostgreSQL *gorm.DB
}

func (r *Resource) Close() {
	utils.LogInfo("Closing all connection ...")
	if r.Config.IsEnablePostgreSQL {
		r.PostgreSQL.Close()
	}
}

func Init(config constant.Config) (*Resource, error) {
	r := Resource{}
	r.Config = config

	// PostgreSQL
	if config.IsEnablePostgreSQL {
		utils.LogInfo("Initializing PostgreSQL connection")
		pg, err := initPostgreSQL(config.IsEnablePostgreSQLLogger)
		if err != nil {
			utils.LogError("Failed PostgreSQL connection ...", err)
			return nil, err
		}
		utils.LogInfo("Connected PostgreSQL successfully !!!")
		r.PostgreSQL = pg
	}

	return &r, nil
}

func (r *Resource) MigrateDatabase() {
	r.PostgreSQL.AutoMigrate(&model.Author{}, &model.Photo{})
}

func initPostgreSQL(enableLogger bool) (*gorm.DB, error) {
	dbSQL, err := gorm.Open("postgres",
		"postgres://"+constant.PostgresUser+
			":"+constant.PostgresPassword+"@"+constant.PostgresDockerName+
			":"+constant.PostgresPort+"/"+constant.PostgresDB+
			"?sslmode=disable")

	//"postgres://postgres:Password1@db/artify_dev?sslmode=disable"

	if err != nil {
		return dbSQL, err
	}
	dbSQL.DB()
	dbSQL.DB().SetMaxIdleConns(10)
	dbSQL.DB().SetMaxOpenConns(150)
	return dbSQL, err
}
