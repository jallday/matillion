package config

import (
	"gitlab.com/joshuaAllday/matillion/pkg/utils"
	"gitlab.com/joshuaAllday/matillion/pkg/utils/models"
)

type Config struct {
	Service     *models.CoreService
	SqlSettings *models.SqlSettings
}

func LoadConfig() *Config {
	c := &Config{}
	c.loadService()
	c.loadSqlService()
	return c
}

func (c *Config) loadService() {
	c.Service = &models.CoreService{
		Port: utils.GetEnvInt("PORT", models.PORT),
	}
}

func (c *Config) loadSqlService() {
	c.SqlSettings = &models.SqlSettings{
		MasterURL:   utils.GetEnvString("MASTER_URL", models.MasterURL),
		ReplicaURLS: utils.GetEnvStringArray("REPLICA_URLS", nil),
		Seed:        utils.GetEnvBool("SEED_DATABASE", true),
	}
}
