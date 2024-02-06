package env

import (
	"github.com/JOKR-Services/goutil/util"
	"github.com/JOKR-Services/logr-go"
	"github.com/joho/godotenv"
)

var Env *Environment

type Environment struct {
	Env string
	Ifood
	Storage
}

type Ifood struct {
	URL          string
	ClientID     string
	ClientSecret string
}

type Storage struct {
	MongoUri string
	DbName   string
}

func Get() *Environment {
	if Env != nil {
		return Env
	}

	var env Environment
	env.Env = util.GetenvStr("ENV", "development")
	if env.Env == "development" {
		err := godotenv.Load()
		if err != nil {
			logr.LogPanic("App", err, logr.KindInfra, nil)
		}
	}

	env.Ifood = Ifood{
		URL:          util.GetenvStr("IFOOD_URL", ""),
		ClientID:     util.GetenvStr("IFOOD_CLIENT_ID", ""),
		ClientSecret: util.GetenvStr("IFOOD_CLIENT_SECRET", ""),
	}

	env.Storage = Storage{
		MongoUri: util.GetenvStr("MONGO_URI", "mongodb://localhost:27017"),
		DbName:   util.GetenvStr("MONGO_DB_NAME", "ifood_nfs_rerun"),
	}

	Env = &env
	return Env
}
