package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	// "gorm.io/driver/postgres"
	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbEnv struct {
	DbUser    string `env:"DB_POSTGRES_USER" envDefault:"postgres"`
	DbPass    string `env:"DB_POSTGRES_PASS" envDefault:"Jul3225501"`
	DbName    string `env:"DB_POSTGRES_NAME" envDefault:"golang-micro"`
	DbAddres  string `env:"DB_POSTGRES_ADDRESS" envDefault:"localhost"`
	DbPort    string `env:"DB_POSTGRES_PORT" envDefault:"5432"`
	DbDebug   bool   `env:"DB_POSTGRES_DEBUG" envDefault:"true"`
	DbType    string `env:"DB_POSTGRES_TYPE" envDefault:"postgres"`
	SslMode   string `env:"DB_POSTGRES_SSL_MODE" envDefault:"disable"`
	DbTimeout string `env:"DB_POSTGRES_TIMEOUT" envDefault:"30"`
}

var (
	// Dbcon ..
	Dbcon *gorm.DB

	// Errdb ..
	Errdb error
	dbEnv DbEnv
)

func init() {
	fmt.Println("DB POSTGRES")
	if err := env.Parse(&dbEnv); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if DbOpen() != nil {
		fmt.Println("Can't Open ", dbEnv.DbName, " DB", DbOpen())
	}
	Dbcon = GetDbCon()
}

// DbOpen ..
func DbOpen() error {
	args := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%s", dbEnv.DbAddres, dbEnv.DbPort, dbEnv.DbUser, dbEnv.DbPass, dbEnv.DbName, dbEnv.SslMode, dbEnv.DbTimeout)
	Dbcon, Errdb = gorm.Open(postgres.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if Errdb != nil {
		fmt.Println(fmt.Sprintf("open db Err :%s ", Errdb.Error()))
		return Errdb
	}

	db, err := Dbcon.DB()
	if err != nil {
		fmt.Println(fmt.Sprintf("Db Not Connect test Ping : %s", err.Error()))
		fmt.Println("Can't Open db Postgres")
	}

	if errping := db.Ping(); errping != nil {
		fmt.Println(fmt.Sprintf("Db Not Connect test Ping : %s", errping.Error()))
		fmt.Println("Can't Open db Postgres")
		return errping
	}
	return nil
}

// GetDbCon ..
func GetDbCon() *gorm.DB {
	//TODO looping try connection until timeout
	// using channel timeout
	if Dbcon == nil {
		if errping := DbOpen(); errping != nil {
			fmt.Errorf(fmt.Sprintf("try to connect again but error : %s", errping.Error()))
		}
	}
	db, err := Dbcon.DB()
	if err != nil {
		fmt.Errorf(fmt.Sprintf("Db Not Connect test Ping : %s", err.Error()))
		fmt.Println("Can't Open db Postgres")
	}
	if errping := db.Ping(); errping != nil {
		fmt.Errorf(fmt.Sprintf("Db Not Connect test Ping : %s", errping.Error()))
		//errping = nil
		if errping = DbOpen(); errping != nil {
			fmt.Errorf(fmt.Sprintf("try to connect again but error : %s", errping.Error()))
		}
	}

	return Dbcon
}

// TotalRow ..
type TotalRow struct {
	Total int64 `gorm:"column(total)"`
}

// AsyncRawQuery ..
func AsyncRawQuery(query string, order string, res interface{}, gormchan chan *gorm.DB) {

	sql := Dbcon.Raw(query + order).Scan(res)

	gormchan <- sql
}
