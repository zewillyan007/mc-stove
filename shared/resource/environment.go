package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Environment struct {
	Server            Server             `json:"server"`
	Brokers           []*Broker          `json:"brokers"`
	DataBases         []*DataBase        `json:"databases"`
	ApiGateways       []*ApiGateway      `json:"apigateways"`
	DeployFiles       *DeployFiles       `json:"deployfiles"`
	GetEdition        *GetEdition        `json:"getedition"`
	GetEditionLaChica *GetEditionLaChica `json:"geteditionlachica"`
	TicketImage       *TicketImage       `json:"ticketimage"`
	EnvEmail          *EnvEmail          `json:"envemail"`
	PathIndex         *PathIndex         `json:"pathindex"`
	Doorman           *Doorman           `json:"doorman"`
	AwsS3             *AwsS3             `json:"awss3"`
	CouchBase         *CouchBase         `json:"couchbase"`
	Session           *Session           `json:"session"`
	Cache             *Cache             `json:"cache"`
	Iccid             *Iccid             `json:"iccid"`
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

type Broker struct {
	Id            string `json:"id"`
	Port          int    `json:"port"`
	Host          string `json:"host"`
	Login         string `json:"login"`
	Password      string `json:"password"`
	DefaultBroker bool   `json:"defaultbroker"`
	Exchange      string `json:"exchange"`
	QueueName     string `json:"queueName"`
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

type DeployFiles struct {
	PathApplication string `json:"pathapplication"`
	PathUpload      string `json:"pathupload"`
}

type Iccid struct {
	Size int `json:"size"`
}
type GetEdition struct {
	Url      string `json:"url"`
	Port     int    `json:"port"`
	EndPoint string `json:"endpoint"`
}

type GetEditionLaChica struct {
	Url      string `json:"url"`
	Port     int    `json:"port"`
	EndPoint string `json:"endpoint"`
	Result   string `json:"result"`
}

type TicketImage struct {
	Url      string `json:"url"`
	Port     int    `json:"port"`
	EndPoint string `json:"endpoint"`
}

type EnvEmail struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type PathIndex struct {
	PathHtml     string `json:"pathhtml"`
	Link         string `json:"link"`
	Reactivate   string `json:"reactivate"`
	ResumeTicket string `json:"resumeTicket"`
}

type Doorman struct {
	Create string `json:"create"`
	Get    string `json:"get"`
}

type AwsS3 struct {
	Key              string `json:"key"`
	Region           string `json:"region"`
	Bucket           string `json:"bucket"`
	Secret           string `json:"secret"`
	InvoicesFilePath string `json:"invoicesfilepath"`
}

type CouchBase struct {
	Dsn      string `json:"dsn"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Session struct {
	Bucket string `json:"bucket"`
	Prefix string `json:"prefix"`
}

type Cache struct {
	Host              string `json:"host"`
	MaxPoolConnection int    `json:"maxPoolConnection"`
	DefaultCache      bool   `json:"defaultcache"`
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

func (env *Environment) GetDefaultBroker() *Broker {

	for _, broker := range env.Brokers {
		if broker.DefaultBroker {
			return broker
		}
	}
	return nil
}

func (env *Environment) GetBrokerId(id string) *Broker {

	for _, broker := range env.Brokers {
		if broker.Id == id {
			return broker
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

func (env *Environment) GetDeployFiles() *DeployFiles {
	return env.DeployFiles
}

func (env *Environment) GetEnvEmail() *EnvEmail {
	return env.EnvEmail
}

func (env *Environment) GetPathIndex() *PathIndex {
	return env.PathIndex
}
