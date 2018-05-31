package resources

import (
	"github.com/NghiaTranUIT/artify-core/constant"
	"github.com/NghiaTranUIT/artify-core/utils"
	"github.com/gobuffalo/pop"
)

type Resource struct {
	Config     constant.Config
	PostgreSQL *pop.Connection
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
		pg, err := initPostgreSQL()
		if err != nil {
			utils.LogError("Failed PostgreSQL connection ...", err)
			return nil, err
		}
		utils.LogInfo("Connected PostgreSQL successfully !!!")
		r.PostgreSQL = pg
	}

	return &r, nil
}

func initPostgreSQL() (*pop.Connection, error) {
	dbSQL, err := pop.Connect("development")
	if err != nil {
		return dbSQL, err
	}
	return dbSQL, err
}
