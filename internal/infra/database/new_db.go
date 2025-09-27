package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB é a instância global do GORM
var DB *gorm.DB

// NewDB inicializa a conexão com o banco SQLite usando GORM
func NewDB() (*gorm.DB, error) {
	// Configuração do GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log das queries
	}

	// Conecta ao banco SQLite
	db, err := gorm.Open(sqlite.Open("emailn.db"), config)
	if err != nil {
		return nil, err
	}

	// Configurações do pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configurações otimizadas para SQLite
	sqlDB.SetMaxOpenConns(1)  // SQLite funciona melhor com 1 conexão
	sqlDB.SetMaxIdleConns(1)  // Mantém 1 conexão idle

	log.Println("✅ Conectado ao SQLite com GORM com sucesso!")

	// Atribui à variável global
	DB = db

	return db, nil
}

// CloseDB fecha a conexão com o banco
func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		log.Println("🔒 Fechando conexão com o banco...")
		return sqlDB.Close()
	}
	return nil
}

// GetDB retorna a instância do GORM
func GetDB() *gorm.DB {
	return DB
}

// HealthCheck verifica se o banco está funcionando
func HealthCheck() error {
	if DB == nil {
		return gorm.ErrInvalidDB
	}
	
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	
	return sqlDB.Ping()
}
