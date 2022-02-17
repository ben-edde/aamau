package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var CFG_path string

func Get_connection() *gorm.DB {
	CFG_map := Read_config(&CFG_path)
	server := CFG_map["SERVER"]
	port := CFG_map["PORT"]
	user := CFG_map["USERNAME"]
	password := CFG_map["PASSWORD"]
	database := CFG_map["DATABASE"]
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, server, port, database)

	conn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return nil
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("DB Connected!\n")
	return conn
}

// read config from file
func Read_config(cfg_path *string) map[string]string {
	CFG := make(map[string]string)
	cfg_dir, _ := path.Split(*cfg_path)
	cfg_file := strings.ReplaceAll(path.Base(*cfg_path), path.Ext(*cfg_path), "")
	viper.SetConfigName(cfg_file)
	viper.AddConfigPath(cfg_dir)
	// dir must be in /
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	CFG["USERNAME"] = viper.GetString("USERNAME")
	CFG["PASSWORD"] = viper.GetString("PASSWORD")
	CFG["SERVER"] = viper.GetString("SERVER")
	CFG["PORT"] = viper.GetString("PORT")
	CFG["DATABASE"] = viper.GetString("DATABASE")
	return CFG
}

func Check_path_exists(path_str string) (bool, error) {
	dir, _ := path.Split(path_str)
	if dir == "" {
		dir, _ = os.Getwd()
	}
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
