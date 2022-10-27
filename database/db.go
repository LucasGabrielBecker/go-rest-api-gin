package database

import (
	"fmt"
	"log"

	"github.com/LucasGabrielBecker/go-rest-api-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  (*gorm.DB)
	err error
)

func ConectaComBancoDeDados() {
	db_host := "localhost"
	db_user := "root"
	db_password := "root"
	db_name := "root"
	db_port := "5435"
	stringDeConexao := fmt.Sprintf(`host=%s
	user=%s
	password=%s
	dbname=%s
	port=%s
	sslmode=disable
	TimeZone=America/Sao_Paulo`, db_host, db_user, db_password, db_name, db_port)
	database, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	database.AutoMigrate(&models.Aluno{})
	DB = database
}
