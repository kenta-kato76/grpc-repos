package database

import (
	"fmt"
	"grpc-repos/domain/entity"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type MySQLDB struct {
	DB *gorm.DB
}

func NewMySQLDB() (*MySQLDB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	// DSN の組み立て
	// ※ parseTime=True, loc=Local などオプションは用途に応じて変更
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// GORM の設定: ロガーやスキーマ名、naming strategy など必要に応じて設定
	// ここでは例として、ログレベルをWarnにし、テーブル名を単数形にする設定をしています
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	// DB接続
	db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// ここで db.DB() を呼び出すと、内部の *sql.DB を取得できます
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get *sql.DB from gorm DB: %w", err)
	}

	// 接続プールの設定
	sqlDB.SetConnMaxLifetime(3 * time.Minute)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)

	// Ping で確認（オプション）
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("MySQL connected successfully (GORM)")

	return &MySQLDB{DB: db}, nil
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{}, &entity.Author{}, &entity.Book{})
}
