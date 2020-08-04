package restapi

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/muharihar/d3ta-go/system/config"
	"github.com/muharihar/d3ta-go/system/handler"
	"github.com/muharihar/d3ta-go/system/initialize"
)

func initConfig(h *handler.Handler) (*config.Config, error) {
	//init config
	cfg, viper, err := config.NewConfig("./")
	if err != nil {
		panic(err)
	}
	h.SetConfig(cfg)

	viper.OnConfigChange(func(e fsnotify.Event) {
		// fmt.Println("config file changed:", e.Name)
		c := new(config.Config)
		if err := viper.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}

		h.SetConfig(c)
		initialize.LoadAllDatabase(h)
	})

	return cfg, nil
}

func initializeSystems(h *handler.Handler) error {

	// initialize database
	if err := initialize.LoadAllDatabase(h); err != nil {
		panic(err)
	}

	return nil
}

// StartRestAPIServer is a function to StartRestAPIServer
func StartRestAPIServer() error {

	// init super handler
	superHandler := new(handler.Handler)

	// init configuration
	cfg, err := initConfig(superHandler)
	if err != nil {
		// fmt.Errorf("StartRestAPIServer.initConfig: %s", err.Error())
		return err
	}

	// initialize Systems
	err = initializeSystems(superHandler)
	if err != nil {
		// fmt.Errorf("StartRestAPIServer.initializeSystems: %s", err.Error())
		return err
	}

	// init echo server
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	// e.Debug = true

	// set header - banner
	e.HideBanner = cfg.Applications.Options.ShowEngineHeader
	if e.HideBanner {
		printSvrHeader(e, cfg)
	}

	// Set routers
	SetRouters(e, superHandler)

	// Start server with Gracefull shutdown
	httpPort := fmt.Sprintf(":%s", cfg.Applications.Options.Listener.Port)
	go func() {
		if err := e.Start(httpPort); err != nil {
			e.Logger.Infof("Shutting down the server [%s]", err.Error())
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	return nil
}
