package resource

import (
	"context"
	"database/sql"
	"errors"
	"mc-stove/shared/types"
	"net/http"
	"os"

	"mc-stove/shared/connection/audit"
	"mc-stove/shared/connection/database"
	"mc-stove/shared/constant"
	"mc-stove/shared/herror"
	"mc-stove/shared/job"
	"mc-stove/shared/logger"
	"mc-stove/shared/middleware"
	"mc-stove/shared/port"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ServerResource struct {
	Db                *gorm.DB
	Env               *Environment
	Log               port.ILogger
	Router            *mux.Router
	Handlers          []port.IHandler
	HttpServer        *http.Server
	Restful           *Restful
	Email             port.IEmail
	Notification      port.INotification
	Middlewares       []mux.MiddlewareFunc
	fnCheckAccess     func(sr *ServerResource) port.ICheckAccessService
	scCheckAccess     port.ICheckAccessService
	fnMonitorTerminal func(sr *ServerResource) port.IMonitor
	MonitorTerminal   port.IMonitor
	// AwsS3             *aws.AwsS3
	// CouchBase    cbapi.CouchBaseApi
	Session port.ISession
	// Cache        port.ICache
	Document     port.IDocument
	Herror       *herror.HandlerError
	SysConf      *SystemConfig
	RecordsLimit int
}

func NewServerResource(configFilePath string) *ServerResource {

	server := &ServerResource{
		Db:          &gorm.DB{},
		Env:         NewEnvironment(configFilePath),
		Log:         logger.NewLoggerManager("server", 3, true),
		Router:      mux.NewRouter(),
		Restful:     NewRestful(logger.NewLoggerManager("request", 4, true), logger.NewLoggerManager("response", 4, true)),
		Handlers:    []port.IHandler{},
		HttpServer:  &http.Server{},
		Middlewares: make([]mux.MiddlewareFunc, 0),
		SysConf:     NewSystemConfig(),
	}
	return server
}

func (sr *ServerResource) Logger(level ...int) port.ILogger {

	callerLevel := 3
	if len(level) > 0 {
		callerLevel = level[0]
	}
	sr.Log.Level(callerLevel)
	return sr.Log
}

func (sr *ServerResource) CreateJobWorkDispatch(log port.ILogger, ctx context.Context) *job.Dispatcher {
	return job.NewDispatcher(log, ctx)
}

func (sr *ServerResource) ConnectDefaultDb(db *DataBase) (*sql.DB, error) {

	var err error
	sr.Db, err = database.Connection(db.Driver, db.Host, db.Port, db.Name, db.User, db.Password, db.Profile)

	if err != nil {
		if strings.Contains(err.Error(), "timeout:") {
			return nil, errors.New("connection timeout")
		} else {
			return nil, err
		}
	}
	sqlDB, err := sr.Db.DB()
	if err != nil {
		return nil, err
	}
	if err = sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, err
	}
	sqlDB.SetMaxIdleConns(db.MaxIdleConns)
	sqlDB.SetMaxOpenConns(db.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(db.LifeTimeConn) * time.Second)
	return sqlDB, nil
}

func (sr *ServerResource) AddHandler(handler port.IHandler) {
	sr.Handlers = append(sr.Handlers, handler)
}

func (sr *ServerResource) SetNotificationHandler(notification port.INotification) {
	sr.Notification = notification
}

func (sr *ServerResource) UseGlobalMiddleware(mw ...mux.MiddlewareFunc) {
	for i := 0; i < len(mw); i++ {
		sr.Middlewares = append(sr.Middlewares, mw[i])
	}
}

func (sr *ServerResource) defaultMiddlewares(router *mux.Router) {
	for _, mw := range sr.Middlewares {
		router.Use(mw)
	}
}

func (sr *ServerResource) DefaultRouter(pathPrefix string, loadDefaultMiddlewares bool) *mux.Router {
	router := sr.Router.PathPrefix(pathPrefix).Subrouter()
	if loadDefaultMiddlewares {
		sr.defaultMiddlewares(router)
	}
	return router
}

func (sr *ServerResource) SetServiceCheckAccess(fn func(sr *ServerResource) port.ICheckAccessService) {
	sr.fnCheckAccess = fn
}

func (sr *ServerResource) ConfigureServiceCheckAccess() {
	sr.scCheckAccess = sr.fnCheckAccess(sr)
}

func (sr *ServerResource) SetServiceMonitorTerminal(fn func(sr *ServerResource) port.IMonitor) {
	sr.fnMonitorTerminal = fn
}

func (sr *ServerResource) CheckAccess(user interface{}, endPoint, method string) (bool, error) {
	return sr.scCheckAccess.CheckAccess(user, endPoint, method)
}

func (sr *ServerResource) ConfigureServiceMonitorTerminal() {
	sr.MonitorTerminal = sr.fnMonitorTerminal(sr)
}

func (sr *ServerResource) IsConfiguredCheckAccess() bool {
	return sr.scCheckAccess != nil
}

// func (sr *ServerResource) GetTokenUser(r *http.Request) (*valueobject.UserPayloadVao, error) {

// 	user := valueobject.NewUserPayloadVao()
// 	jwtk := helper.NewJwtHelper(sr.Env.Server.SecretKey)
// 	jwtk.Load(r.Header.Get("Access-Token"))

// 	err := json.Unmarshal([]byte(jwtk.GetPayload("user").(string)), &user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (sr *ServerResource) GetTypeUser(r *http.Request) string {

// 	user, err := sr.GetTokenUser(r)
// 	if err != nil {
// 		return ""
// 	}

// 	session, _ := sr.Session.GetData(user.Id)
// 	return session["type"].(string)
// }

// func (sr *ServerResource) GetSessionUser(r *http.Request) (map[string]interface{}, error) {
// 	user, err := sr.GetTokenUser(r)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sr.Session.GetData(user.Id)
// }

// func (sr *ServerResource) GetUserVision(r *http.Request) map[string][]int64 {
// 	data := make(map[string][]int64)
// 	session, err := sr.GetSessionUser(r)
// 	if err != nil {
// 		return nil
// 	}
// 	var companies []int64
// 	if session["id_companies"] != nil {
// 		for _, id := range session["id_companies"].([]interface{}) {
// 			companies = append(companies, int64(id.(float64)))
// 		}
// 	}
// 	var regionals []int64
// 	if session["id_regionals"] != nil {
// 		for _, id := range session["id_regionals"].([]interface{}) {
// 			regionals = append(regionals, int64(id.(float64)))
// 		}
// 	}
// 	data["id_companies"] = companies
// 	data["id_regionals"] = regionals
// 	return data
// }

// func (sr *ServerResource) GetUserVisionCondition(r *http.Request, fieldNameCompany, fieldNameRegional string, condition ...interface{}) []interface{} {

// 	vision := sr.GetUserVision(r)
// 	var cond []interface{}
// 	strcond := make([]string, 0)
// 	companies := make([]string, 0)
// 	regionals := make([]string, 0)
// 	valcond := make(map[string]string, 0)

// 	for _, IdCompany := range vision["id_companies"] {
// 		companies = append(companies, strconv.FormatInt(IdCompany, 10))
// 	}

// 	if len(companies) > 0 {
// 		strcond = append(strcond, fieldNameCompany+" IN (?)")
// 		valcond["companies"] = strings.Join(companies, ",")
// 	}

// 	for _, IdRegional := range vision["id_regionals"] {
// 		regionals = append(regionals, strconv.FormatInt(IdRegional, 10))
// 	}

// 	if len(regionals) > 0 {
// 		strcond = append(strcond, fieldNameRegional+" IN (?)")
// 		valcond["regionals"] = strings.Join(regionals, ",")
// 	}

// 	var strInCondition string
// 	var valConditionIn []interface{}

// 	if len(condition) > 0 {
// 		strInCondition = "AND " + condition[0].(string)
// 		valConditionIn = condition[1:]
// 	}

// 	if len(companies) > 0 || len(regionals) > 0 {
// 		if len(companies) > 0 && len(regionals) > 0 {
// 			if len(valConditionIn) > 0 {
// 				cond = append(cond, strings.Join(strcond, " AND ")+" "+strInCondition, valcond["companies"], valcond["regionals"])
// 				for _, v := range valConditionIn {
// 					cond = append(cond, v)
// 				}
// 			} else {
// 				cond = append(cond, strings.Join(strcond, " AND "), valcond["companies"], valcond["regionals"])
// 			}
// 		} else if len(companies) > 0 {
// 			if len(valConditionIn) > 0 {
// 				cond = append(cond, strings.Join(strcond, " AND ")+" "+strInCondition, valcond["companies"])
// 				for _, v := range valConditionIn {
// 					cond = append(cond, v)
// 				}
// 			} else {
// 				cond = append(cond, strings.Join(strcond, " AND "), valcond["companies"])
// 			}
// 		} else {
// 			if len(valConditionIn) > 0 {
// 				cond = append(cond, strings.Join(strcond, " AND ")+" "+strInCondition, valcond["regionals"])
// 				for _, v := range valConditionIn {
// 					cond = append(cond, v)
// 				}
// 			} else {
// 				cond = append(cond, strings.Join(strcond, " AND "), valcond["regionals"])
// 			}
// 		}
// 	} else {
// 		if len(condition) > 0 {
// 			return condition
// 		} else {
// 			return cond
// 		}
// 	}

// 	return cond
// }

// func (sr *ServerResource) GetUserVisionConditionCompany(r *http.Request, fieldNameCompany string, condition ...interface{}) []interface{} {

// 	vision := sr.GetUserVision(r)
// 	var cond []interface{}
// 	strcond := ""
// 	valcond := ""
// 	companies := make([]string, 0)

// 	for _, IdCompany := range vision["id_companies"] {
// 		companies = append(companies, strconv.FormatInt(IdCompany, 10))
// 	}

// 	var strInCondition string
// 	var valConditionIn []interface{}

// 	if len(companies) > 0 {

// 		strcond = fieldNameCompany + " IN (?)"
// 		valcond = strings.Join(companies, ",")

// 		if len(condition) > 0 {
// 			strInCondition = "AND " + condition[0].(string)
// 			valConditionIn = condition[1:]
// 		}

// 		if len(valConditionIn) > 0 {
// 			cond = append(cond, strcond+" "+strInCondition, valcond)
// 			for _, v := range valConditionIn {
// 				cond = append(cond, v)
// 			}
// 		} else {
// 			cond = append(cond, strcond, valcond)
// 		}

// 	} else {

// 		if len(condition) > 0 {
// 			return condition
// 		} else {
// 			return cond
// 		}
// 	}

// 	return cond
// }

// func (sr *ServerResource) GetUserVisionConditionRegional(r *http.Request, fieldNameRegional string, condition ...interface{}) []interface{} {

// 	vision := sr.GetUserVision(r)
// 	var cond []interface{}
// 	strcond := ""
// 	valcond := ""
// 	regionals := make([]string, 0)

// 	for _, IdCompany := range vision["id_regionals"] {
// 		regionals = append(regionals, strconv.FormatInt(IdCompany, 10))
// 	}

// 	var strInCondition string
// 	var valConditionIn []interface{}

// 	if len(regionals) > 0 {

// 		strcond = fieldNameRegional + " IN (?)"
// 		valcond = strings.Join(regionals, ",")

// 		if len(condition) > 0 {
// 			strInCondition = "AND " + condition[0].(string)
// 			valConditionIn = condition[1:]
// 		}

// 		if len(valConditionIn) > 0 {
// 			cond = append(cond, strcond+" "+strInCondition, valcond)
// 			for _, v := range valConditionIn {
// 				cond = append(cond, v)
// 			}
// 		} else {
// 			cond = append(cond, strcond, valcond)
// 		}

// 	} else {

// 		if len(condition) > 0 {
// 			return condition
// 		} else {
// 			return cond
// 		}
// 	}

// 	return cond
// }

// func (sr *ServerResource) GetUserCompanies(r *http.Request) []int64 {
// 	vision := sr.GetUserVision(r)
// 	return vision["id_companies"]
// }

// func (sr *ServerResource) GetUserRegionals(r *http.Request) []int64 {
// 	vision := sr.GetUserVision(r)
// 	return vision["id_regionals"]
// }

func (sr *ServerResource) TypeUser(IdUser int64) string {
	id := strconv.FormatInt(IdUser, 10)
	session, _ := sr.Session.GetData(id)

	if _, ok := session["type"].(string); !ok {
		return ""
	}

	return session["type"].(string)
}

func (sr *ServerResource) IsUser(IdUser int64) bool {
	return sr.TypeUser(IdUser) == constant.USER_TYPE_USER
}

func (sr *ServerResource) IsAdm(IdUser int64) bool {
	return sr.TypeUser(IdUser) == constant.USER_TYPE_ADM
}

func (sr *ServerResource) IsSuper(IdUser int64) bool {
	return sr.TypeUser(IdUser) == constant.USER_TYPE_SUPER
}

// func (sr *ServerResource) LoadSysConf() error {

// 	if sr.Db != nil {
// 		tx := sr.Db.Session(&gorm.Session{Logger: glogger.Default.LogMode(glogger.Silent)})
// 		rows, err := tx.Raw("SELECT * FROM sca.config").Rows()
// 		if err != nil {
// 			return err
// 		}
// 		sr.SysConf.Load(rows)
// 	}
// 	return nil
// }

func (sr *ServerResource) GetRecordsLimit() (int, error) {

	var defaultValue int

	v := os.Getenv("RECORD_LIMIT")
	if len(v) == 0 {
		defaultValue := 1000
		return defaultValue, nil
	}

	defaultValue, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	return defaultValue, nil
}

func (sr *ServerResource) Run(ctx context.Context) {

	sr.Logger(2).Info("Initialize Dependencies...")

	var err error
	db := sr.Env.GetDefaultDb()
	if db != nil {
		sr.Logger(2).Info("DB: %v:%v", db.Host, db.Port)
		sqlDB, err := sr.ConnectDefaultDb(db)
		if err != nil {
			sr.Logger(2).Warn("DB: %s", err.Error())
			os.Exit(0)
		} else {
			defer sqlDB.Close()
		}
	}

	// if sr.Env.AwsS3 != nil {
	// 	sr.Logger(2).Info("Aws-S3: %v", sr.Env.AwsS3.Bucket)
	// 	sr.AwsS3, err = aws.NewAwsS3(sr.Env.AwsS3.Key, sr.Env.AwsS3.Secret, sr.Env.AwsS3.Bucket, sr.Env.AwsS3.Region)
	// 	if err != nil {
	// 		sr.Logger(2).Error("Aws-S3: %s", err)
	// 	}
	// }

	// if sr.Env.CouchBase != nil {
	// 	sr.Logger(2).Info("CouchBase: %v", sr.Env.CouchBase.Dsn)
	// 	sr.CouchBase, err = cbapi.New(sr.Env.CouchBase.Dsn, sr.Env.CouchBase.User, sr.Env.CouchBase.Password)
	// 	if err != nil {
	// 		sr.Logger(2).Error("CouchBase: %s", err)
	// 	}
	// }

	// if sr.Env.Session != nil && sr.CouchBase != nil {
	// 	sr.Session = adapter.NewSession(sr.Env.Session.Bucket, sr.Env.Session.Prefix, sr.Env.Server.TokenExpires, sr.CouchBase)
	// 	if err != nil {
	// 		sr.Logger(2).Error("Session: %s", err)
	// 	}
	// }

	// if sr.Env.GetDefaultCache() != nil {
	// 	sr.Cache = managercache.NewManagerCacheClient(sr.Env.Cache.Host, sr.Env.Cache.MaxPoolConnection)
	// 	if err != nil {
	// 		sr.Logger(2).Error("Cache: %s", err)
	// 	} else {
	// 		sr.Logger(2).Info("Cache: %v", "Active")
	// 	}
	// }

	// if sr.CouchBase != nil {
	// 	sr.Document = adapter.NewDocument(sr.CouchBase)
	// }

	sr.Herror = herror.NewHandlerError()

	// err = sr.LoadSysConf()
	// if err != nil {
	// 	sr.Logger(2).Error("System Config: %s", err)
	// } else {
	// 	sr.Logger(2).Info("System Config: %s", "Loaded")
	// }

	//GetRecordsLimit
	sr.RecordsLimit, err = sr.GetRecordsLimit()
	if err != nil {
		sr.Logger(2).Error("Load error records limit: %s", err)
	}

	sr.Logger(2).Info("Record limit: %d", sr.RecordsLimit)

	// sr.Logger(2).Info("Adding Brokers...")

	// brokerConfig := sr.Env.GetDefaultBroker()
	// if brokerConfig != nil {
	// 	RabbitMQ, err := adapter.NewRabbitMQ(brokerConfig.Login, brokerConfig.Password, brokerConfig.Host, strconv.Itoa(brokerConfig.Port))
	// 	if err != nil {
	// 		sr.Logger(2).Error("%s", err)
	// 	} else {
	// 		sr.Email = adapter.NewEmailBroker(RabbitMQ)
	// 	}
	// }

	if auditBroker := sr.Env.GetBrokerId("audit"); auditBroker != nil {
		if _, err = audit.NewDefaultPostman(
			&audit.Config{
				Sender:    sr.Env.Server.Name,
				Exchange:  auditBroker.Exchange,
				QueueName: auditBroker.QueueName,
				Publisher: &audit.PublisherConfig{
					User:     auditBroker.Login,
					Password: auditBroker.Password,
					Host:     auditBroker.Host,
					Port:     strconv.Itoa(auditBroker.Port),
				},
			},
		); err != nil {
			sr.Logger(2).Error("%s", err)
		}
	}

	sr.Logger(2).Info("Adding Handlers Routes...")

	for _, handler := range sr.Handlers {
		handler.MakeRoutes()
	}

	if sr.Notification != nil {
		sr.Logger(2).Info("Configure Notification Service...")
		sr.Notification.MakeServices()
	}

	if sr.fnCheckAccess != nil {
		sr.Logger(2).Info("Configure Check Access Service...")
		sr.ConfigureServiceCheckAccess()
	}

	if sr.fnMonitorTerminal != nil {
		sr.Logger(2).Info("Configure Monitor Terminal Service...")
		sr.ConfigureServiceMonitorTerminal()
		sr.MonitorTerminal.StartOnce()
	}

	sr.Logger(2).Info("Server Starting (%s %s): PID (%d)...", sr.Env.Server.Name, sr.Env.Server.Tag, sr.Env.Server.PID)

	var address string
	listen := sr.Env.GetDefaultListner()
	address = listen.Host + ":" + strconv.Itoa(listen.Port)

	if sr.Env.Server.Cors {
		sr.Logger(2).Info("Cors: %v", sr.Env.Server.Cors)
		sr.Router.Use(middleware.Cors)
	}
	sr.Router.Use(middleware.Reqid)
	sr.Router.Use(middleware.ManagerContext(sr.Env.Server.SecretKey, sr.Log))
	sr.HttpServer = &http.Server{Addr: address, Handler: sr.Router}

	apigateway := sr.Env.GetDefaultApiGateway()
	if apigateway != nil {
		sr.Logger(2).Info("Initialize Gateway Comunication...")
		Gateway := NewGateway(apigateway.Scheme, apigateway.Host, strconv.Itoa(apigateway.Port), apigateway.Interval, apigateway.IntervalInc, logger.NewLoggerManager("gateway", 3, true))
		Gateway.register(listen.Host, strconv.Itoa(listen.Port), sr.Env.Server.Name)
		defer Gateway.deregister(listen.Host, strconv.Itoa(listen.Port), sr.Env.Server.Name)
	}

	go func() {
		sr.Logger(2).Info("Listening on: %s", address)
		if !listen.Ssl.Enabled {
			err = sr.HttpServer.ListenAndServe()
		} else {
			err = sr.HttpServer.ListenAndServeTLS(listen.Ssl.CertFile, listen.Ssl.KeyFile)
		}
		if err != http.ErrServerClosed {
			sr.Logger(2).Error("%s", err)
			os.Exit(0)
		}
	}()
	<-ctx.Done()

	sr.Logger(2).Info("Server Stopped...")

	if listen.GracefulShutdown {
		ctxShutdown, cancel := context.WithTimeout(context.Background(), time.Duration(listen.ShutdownTimeout)*time.Second)
		defer func() {
			cancel()
		}()
		if err = sr.HttpServer.Shutdown(ctxShutdown); err != nil {
			sr.Logger(2).Error("Shutdown Failed: %s", err)
		} else {
			sr.Logger(2).Info("Shutdown Graceful")
		}
	}
}

func (sr *ServerResource) ManagerContext(r *http.Request) *types.ManagerContext {
	if i := r.Context().Value(constant.CONTEXT_ID); i != nil {
		return i.(*types.ManagerContext)
	}
	return nil
}

func (sr *ServerResource) GetDB() *gorm.DB {
	return sr.Db
}

// func (sr *ServerResource) GetCache() port.ICache {
// 	return sr.Cache
// }

func (sr *ServerResource) RootRouteStr(r *http.Request) string {
	s, _ := mux.CurrentRoute(r).GetPathTemplate()
	return strings.Split(s, "/")[1]
}

//com cache
// func (sr *ServerResource) Provider(r *http.Request) port.IResourceProvider {
// 	rp := NewResourceProvider(sr.Db, sr.ManagerContext(r), sr.Cache)
// 	rp.CurrentRoute, _ = mux.CurrentRoute(r).GetPathTemplate()
// 	rp.RootRoute = strings.Split(rp.CurrentRoute, "/")[1]

//		return rp
//	}

func (sr *ServerResource) Provider(r *http.Request) port.IResourceProvider {
	rp := NewResourceProvider(sr.Db, sr.ManagerContext(r))
	rp.CurrentRoute, _ = mux.CurrentRoute(r).GetPathTemplate()
	rp.RootRoute = strings.Split(rp.CurrentRoute, "/")[1]

	return rp
}
