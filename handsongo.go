package main

import (
	"encoding/base64"
	"fmt"
	"github.com/Sfeir/handsongo/dao"
	"github.com/Sfeir/handsongo/utils"
	"github.com/Sfeir/handsongo/web"
	logger "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

var (
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string

	port               = 8020
	logLevel           = "warning"
	db                 = "mongodb://mongo/spirits"
	logFormat          = utils.TextFormatter
	statisticsDuration = 20 * time.Second

	header, _ = base64.StdEncoding.DecodeString(
		"ICAgICAgICAgLF8tLS1+fn5+fi0tLS0uXyAgICAgICAgIAogIF8sLF8sKl5fX19fICAgICAgX19fX19gYCpnKlwiKi" +
			"wgCiAvIF9fLyAvJyAgICAgXi4gIC8gICAgICBcIF5AcSAgIGYgClsgIEBmIHwgQCkpICAgIHwgIHwgQCkpICAgbCA" +
			"gMCBfLyAgCiBcYC8gICBcfl9fX18gLyBfXyBcX19fX18vICAgIFwgICAKICB8ICAgICAgICAgICBfbF9fbF8gICAg" +
			"ICAgICAgIEkgICAKICB9ICAgICAgICAgIFtfX19fX19dICAgICAgICAgICBJICAKICBdICAgICAgICAgICAgfCB8I" +
			"HwgICAgICAgICAgICB8ICAKICBdICAgICAgICAgICAgIH4gfiAgICAgICAgICAgICB8ICAKICB8ICAgICAgICAgIC" +
			"AgICAgICAgICAgICAgICAgIHwgICAKICAgfCAgICAgICAgICAgICAgICAgICAgICAgICAgIHwg")
)

func main() {

	// new app
	app := cli.NewApp()
	app.Name = "handsongo"
	app.Usage = "handsongo service launcher"

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	app.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	app.Authors = []cli.Author{cli.Author{Name: "sfr"}}
	app.Copyright = "Sfeir " + strconv.Itoa(time.Now().Year())

	// command line flags
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Value:       port,
			Name:        "port",
			Usage:       "Set the listening port of the webserver",
			Destination: &port,
		},
		cli.StringFlag{
			Value:       db,
			Name:        "db",
			Usage:       "Set the mongo database connection string",
			Destination: &db,
		},
		cli.StringFlag{
			Value:       logLevel,
			Name:        "logl",
			Usage:       "Set the output log level (debug, info, warning, error)",
			Destination: &logLevel,
		},

		// TODO add string flag called "-logf" to choose the log format

		// TODO add a duration flag called "-statd" to define a statistic check period

	}

	// main action
	// sub action are also possible
	app.Action = func(c *cli.Context) error {
		// print header
		fmt.Println(string(header))

		fmt.Print("* --------------------------------------------------- *\n")
		fmt.Printf("|   port                    : %d\n", port)
		fmt.Printf("|   db                      : %s\n", db)
		fmt.Printf("|   logger level            : %s\n", logLevel)
		fmt.Printf("|   logger format           : %s\n", logFormat)
		fmt.Printf("|   statistic duration(s)   : %0.f\n", statisticsDuration.Seconds())
		fmt.Print("* --------------------------------------------------- *\n")

		// init log options from command line params
		err := utils.InitLog(logLevel, logFormat)
		if err != nil {
			logger.Warn("error setting log level, using debug as default")
		}

		// build the web server
		webServer, err := web.BuildWebServer(db, dao.DAOMongo, statisticsDuration)

		if err != nil {
			return err
		}

		// serve
		webServer.Run(":" + strconv.Itoa(port))

		return nil
	}

	// run the app
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Run error %q\n", err)
	}
}
