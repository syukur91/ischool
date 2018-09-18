package main

import (
	"os"

	"github.com/Gurpartap/logrus-stack"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
)

var nfLog *log.Entry

func init() {
	// set logger Format
	log.SetFormatter(&log.JSONFormatter{})

	callerLevels := []log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel}
	stackLevels := []log.Level{log.PanicLevel, log.FatalLevel, log.ErrorLevel}

	log.AddHook(logrus_stack.NewHook(callerLevels, stackLevels))

	// Log default fields
	nfLog = log.WithFields(log.Fields{
		"app":  "container-mgmt",
		"type": "backend",
	})
}

func main() {

	// Its possible to specify env file
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}

	// Load env data
	err := godotenv.Load(envFile)
	if err != nil {
		nfLog.Fatal(err.Error())
		panic(err)
	}

	e := echo.New()
	// e.HideBanner = true

	//e.Use(nfmiddleware.Logger(os.Getenv("APP_NAME")))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Postgres
	db, err := gorm.Open(os.Getenv("DATABASE_DIALECT"), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to Database: " + err.Error())
	}
	logMode := os.Getenv("DATABASE_LOGMODE")
	db.LogMode(logMode == "true")
	defer db.Close()

	setupHandlers(e, db, nfLog)

	servicePort := os.Getenv("PORT")
	nfLog.Info("Container mgmt API Started at Port " + servicePort)
	e.Start(servicePort)
}
