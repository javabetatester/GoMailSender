package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB é a instância global do GORM
var DB *gorm.DB

// NewDB inicializa a conexão com o banco PostgreSQL usando GORM
func NewDB() (*gorm.DB, error) {
	// Configuração do GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log das queries
	}

	// String de conexão PostgreSQL
	dsn := getDSN()

	// Conecta ao banco PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com PostgreSQL: %w", err)
	}

	// Configurações do pool de conexões
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configurações otimizadas para PostgreSQL
	sqlDB.SetMaxOpenConns(25)    // Máximo de conexões abertas
	sqlDB.SetMaxIdleConns(10)    // Máximo de conexões idle
	sqlDB.SetConnMaxLifetime(0)  // Conexões não expiram

	log.Println("✅ Conectado ao PostgreSQL com GORM com sucesso!")

	// Atribui à variável global
	DB = db

	return db, nil
}

// getDSN retorna a string de conexão do PostgreSQL
func getDSN() string {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "emailn")
	sslmode := getEnv("DB_SSLMODE", "disable")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
}

// getEnv retorna o valor da variável de ambiente ou um valor padrão
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
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
