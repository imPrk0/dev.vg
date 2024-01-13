package conf

import (
	"github.com/imPrk0/dev.vg/pkg/util"
	"github.com/joho/godotenv"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	envFilePath = ".env"
)

type webserver struct {
	Host string
	Port int
}

type limit struct {
	MaxWords       int
	MaxSentences   int
	MaxParagraphs  int
	ImageMaxWidth  int
	ImageMaxHeight int
}

type functions struct {
	EnableShortLink        bool
	EnableRandomText       bool
	EnableImagePlaceholder bool
	EnableRandomPick       bool
}

type database struct {
	Type   string
	Prefix string
}

type mysql_db struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type sqlite3_db struct {
	Database string
}

var (
	// Debug mode
	Debug bool

	// Web server
	WebHost string
	WebPort int

	// Limit
	MaxWords       int
	MaxSentences   int
	MaxParagraphs  int
	ImageMaxWidth  int
	ImageMaxHeight int

	// Functions
	EnableShortLink        bool
	EnableRandomText       bool
	EnableImagePlaceholder bool
	EnableRandomPick       bool

	// Image placeholder work mode
	ImagePlaceholderWorkMode string

	// Database
	DatabaseType   string
	DatabasePrefix string

	// MySQL
	MySQLHost     string
	MySQLPort     int
	MySQLUsername string
	MySQLPassword string
	MySQLDatabase string

	// SQLite 3
	SQLite3Database string
)

func InitConfig() {
	if _, err := os.Stat(envFilePath); os.IsNotExist(err) {
		util.Log().Warning("There is no environment variable file, the program will create one, please modify it as needed and run it again.")
		err := ioutil.WriteFile(envFilePath, []byte(DefaultEnvFile), 0644)
		if nil != err {
			util.Log().Error("Error creating .env file: ", err)
			return
		}
		os.Exit(1)
	}

	err := godotenv.Load()
	if err != nil {
		util.Log().Error("Error loading .env file: ", err)
		os.Exit(1)
		return
	}
}

func changeConfigEnv() {
	// Web server host
	host := os.Getenv("WEB_SERVER_HOST")
	if "" != host {
		WebHost = host
	} else {
		WebHost = WebServer.Host
	}

	// Web server port
	port := os.Getenv("WEB_SERVER_PORT")
	if "" != port {
		WebPort, _ = strconv.Atoi(port)
	} else {
		WebPort = WebServer.Port
	}
}
