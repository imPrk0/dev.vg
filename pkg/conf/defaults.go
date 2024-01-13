package conf

var (
	DebugMode = false

	WebServer = &webserver{
		Host: "0.0.0.0",
		Port: 8080,
	}

	Limit = &limit{
		MaxWords:       500,
		MaxSentences:   200,
		MaxParagraphs:  100,
		ImageMaxWidth:  3000,
		ImageMaxHeight: 3000,
	}

	Functions = &functions{
		EnableShortLink:        false,
		EnableRandomText:       true,
		EnableImagePlaceholder: true,
		EnableRandomPick:       true,
	}

	ImagePlaceholderMode = "redirect"

	Database = &database{
		Type:   "sqlite3",
		Prefix: "dev_vg_",
	}

	MysqlDB = &mysql_db{
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "root",
		Database: "dev_vg",
	}

	SQLite3DB = &sqlite3_db{
		Database: "./prk_database.db",
	}

	DefaultEnvFile = `
# Web server
WEB_SERVER_HOST="0.0.0.0"
WEB_SERVER_PORT=8080

# Limit
MAX_WORDS=500
MAX_SENTENCES=200
MAX_PARAGRAPHS=100
IMAGE_MAX_WIDTH=3000
IMAGE_MAX_HEIGHT=3000

# Functions
ENABLE_SHORT_LINK=false
ENABLE_RANDOM_TEXT=true
ENABLE_IMAGE_PLACEHOLDER=true
ENABLE_RANDOM_PICK=true

# Image placeholder work mode (local / redirect)
IMAGE_PLACEHOLDER_MODE="redirect"

########################################
# If you need the short link function, #
#   you must configure the database!   #
########################################

# DataBase (mysql / sqlite3)
DB_TYPE="sqlite3"
DB_PREFIX="dev_vg_"

# Only MySQL DataBase
MYSQL_HOST="127.0.0.1"
MYSQL_PORT=3306
MYSQL_USERNAME="root"
MYSQL_PASSWORD="root"
MYSQL_DATABASE="dev_vg"

# Only SQLite3 DataBase
SQLITE3_DB_PATH="./prk_database.db"
`
)
