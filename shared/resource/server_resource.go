package resource

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"os"

	"mp-projeto/shared/connection/database"
	"mp-projeto/shared/logger"
	"mp-projeto/shared/middleware"
	"mp-projeto/shared/port"
	"strconv"
	"strings"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ServerResource struct {
	Db         *gorm.DB
	Env        *Environment
	Log        port.ILogger
	Router     *mux.Router
	Handlers   []port.IHandler
	HttpServer *http.Server
	Restful    *Restful
	//Email        port.IEmail
	Notification port.INotification
	Middlewares  []mux.MiddlewareFunc
	//Session      port.ISession
	Cache port.ICache
	//Document     port.IDocument
	SysConf *SystemConfig
}

func NewServerResource(configFilePath string) *ServerResource {

	server := &ServerResource{
		Db:          &gorm.DB{},
		Env:         NewEnvironment(configFilePath),
		Log:         logger.NewLoggerPhoenix("server", 3, true),
		Router:      mux.NewRouter(),
		Restful:     NewRestful(logger.NewLoggerPhoenix("request", 4, true), logger.NewLoggerPhoenix("response", 4, true)),
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

// func (sr *ServerResource) SetServiceCheckAccess(fn func(sr *ServerResource) port.ICheckAccessService) {
// 	sr.fnCheckAccess = fn
// }

// func (sr *ServerResource) ConfigureServiceCheckAccess() {
// 	sr.scCheckAccess = sr.fnCheckAccess(sr)
// }

// func (sr *ServerResource) CheckAccess(user interface{}, endPoint, method string) (bool, error) {
// 	return sr.scCheckAccess.CheckAccess(user, endPoint, method)
// }

// func (sr *ServerResource) IsConfiguredCheckAccess() bool {
// 	return sr.scCheckAccess != nil
// }

// func (sr *ServerResource) TypeUser(IdUser int64) string {
// 	id := strconv.FormatInt(IdUser, 10)
// 	session, _ := sr.Session.GetData(id)

// 	if _, ok := session["type"].(string); !ok {
// 		return ""
// 	}

// 	return session["type"].(string)
// }

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

	if sr.Env.GetDefaultCache() != nil {
		config := bigcache.Config{
			Shards:             sr.Env.Cache.Shards,
			LifeWindow:         sr.Env.Cache.LifeWindow * time.Second,
			CleanWindow:        sr.Env.Cache.CleanWindow * time.Second,
			MaxEntrySize:       sr.Env.Cache.MaxEntrySize,
			HardMaxCacheSize:   sr.Env.Cache.HardMaxCacheSize,
			MaxEntriesInWindow: sr.Env.Cache.MaxEntriesInWindow,
			Logger:             bigcache.DefaultLogger(),
		}
		sr.Cache, _ = bigcache.New(context.Background(), config)
		if err != nil {
			sr.Logger(2).Error("Cache: %s", err)
		} else {
			sr.Logger(2).Info("Cache: %v", "Active")
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

	sr.Logger(2).Info("Server Starting (%s %s): PID (%d)...", sr.Env.Server.Name, sr.Env.Server.Tag, sr.Env.Server.PID)

	var address string
	listen := sr.Env.GetDefaultListner()
	address = listen.Host + ":" + strconv.Itoa(listen.Port)

	if sr.Env.Server.Cors {
		sr.Logger(2).Info("Cors: %v", sr.Env.Server.Cors)
		sr.Router.Use(middleware.Cors)
	}
	sr.Router.Use(middleware.Reqid)
	sr.HttpServer = &http.Server{Addr: address, Handler: sr.Router}

	apigateway := sr.Env.GetDefaultApiGateway()
	if apigateway != nil {
		sr.Logger(2).Info("Initialize Gateway Comunication...")
		Gateway := NewGateway(apigateway.Scheme, apigateway.Host, strconv.Itoa(apigateway.Port), apigateway.Interval, apigateway.IntervalInc, logger.NewLoggerPhoenix("gateway", 3, true))
		go Gateway.register(listen.Host, strconv.Itoa(listen.Port), sr.Env.Server.Name)
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
