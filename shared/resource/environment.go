package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

type Environment struct {
	Server      Server        `json:"server"`
	DataBases   []*DataBase   `json:"databases"`
	ApiGateways []*ApiGateway `json:"apigateways"`
	Doorman     *Doorman      `json:"doorman"`
	Cache       *Cache        `json:"cache"`
}

type Server struct {
	PID          int     `json:"-"`
	Tag          string  `json:"tag"`
	Name         string  `json:"name"`
	Cors         bool    `json:"cors"`
	IdAdm        []int64 `json:"idadm"`
	CipherKey    string  `json:"cipher_key"`
	SecretKey    string  `json:"secret_key"`
	Application  string  `json:"application"`
	TokenExpires int     `json:"token_expires"`
	Listens      []*Listen
}

type Listen struct {
	Id               string `json:"id"`
	Port             int    `json:"port"`
	Host             string `json:"host"`
	Network          string `json:"network"`
	DefaultListen    bool   `json:"defaultlisten"`
	ShutdownTimeout  int    `json:"shutdownsimeout"`
	GracefulShutdown bool   `json:"gracefulshutdown"`
	Ssl              Ssl    `json:"ssl"`
}

type Ssl struct {
	Enabled  bool   `json:"enabled"`
	KeyFile  string `json:"keyfile"`
	CertFile string `json:"certfile"`
}

type DataBase struct {
	Id           string `json:"id"`
	Driver       string `json:"driver"`
	Port         int    `json:"port"`
	Host         string `json:"host"`
	Name         string `json:"name"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Profile      bool   `json:"profile"`
	DefaultDb    bool   `json:"defaultdb"`
	MaxOpenConns int    `json:"maxopenconns"`
	MaxIdleConns int    `json:"maxidleconns"`
	LifeTimeConn int    `json:"lifetime"`
}

type ApiGateway struct {
	Id                string `json:"id"`
	Port              int    `json:"port"`
	Host              string `json:"host"`
	Scheme            string `json:"scheme"`
	Interval          int    `json:"interval"`
	IntervalInc       int    `json:"intervalinc"`
	RegisterName      string `json:"registername"`
	DefaultApiGateway bool   `json:"defaultapigateway"`
}

type Doorman struct {
	Create string `json:"create"`
	Get    string `json:"get"`
}

type Cache struct {
	Shards             int           `json:"shards"`
	LifeWindow         time.Duration `json:"lifewindow"`
	CleanWindow        time.Duration `json:"cleanwindow"`
	MaxEntrySize       int           `json:"maxentrysize"`
	HardMaxCacheSize   int           `json:"hardmaxcachesize"`
	MaxEntriesInWindow int           `json:"maxentriesinwindow"`
	DefaultCache       bool          `json:"defaultcache"`
}

func NewEnvironment(fileFullPath string) *Environment {
	return loadConfigFile(fileFullPath)
}

func loadConfigFile(fileFullPath string) *Environment {

	env := new(Environment)
	extension := filepath.Ext(fileFullPath)
	file, err := ioutil.ReadFile(fileFullPath)

	if err != nil {
		fmt.Printf("File erro: %v\n", err)
		os.Exit(1)
	}

	if extension == ".json" {

		if err := json.Unmarshal(file, env); err != nil {
			panic(err)
		}

	} else if extension == ".toml" {

		if err := toml.Unmarshal(file, env); err != nil {
			panic(err)
		}
	}

	env.Server.PID = os.Getpid()

	return env
}

func (env *Environment) GetDefaultListner() *Listen {

	for _, listen := range env.Server.Listens {
		if listen.DefaultListen {
			return listen
		}
	}
	return nil
}

func (env *Environment) GetListnerById(id string) *Listen {

	for _, listen := range env.Server.Listens {
		if listen.Id == id {
			return listen
		}
	}
	return nil
}

func (env *Environment) GetDefaultDb() *DataBase {

	for _, db := range env.DataBases {
		if db.DefaultDb {
			return db
		}
	}
	return nil
}

func (env *Environment) GetDbById(id string) *DataBase {

	for _, db := range env.DataBases {
		if db.Id == id {
			return db
		}
	}
	return nil
}

func (env *Environment) GetDefaultApiGateway() *ApiGateway {

	for _, apigateway := range env.ApiGateways {
		if apigateway.DefaultApiGateway {
			return apigateway
		}
	}
	return nil
}

func (env *Environment) GetApiGatewayId(id string) *ApiGateway {

	for _, apigateway := range env.ApiGateways {
		if apigateway.Id == id {
			return apigateway
		}
	}
	return nil
}

func (env *Environment) GetDefaultCache() *Cache {

	if env.Cache != nil && env.Cache.DefaultCache {
		return env.Cache
	}
	return nil
}
