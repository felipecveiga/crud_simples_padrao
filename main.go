package main

import (
	"github.com/felipecveiga/crud_simples_padrao/handler"
	"github.com/felipecveiga/crud_simples_padrao/model"
	"github.com/felipecveiga/crud_simples_padrao/repository"
	"github.com/felipecveiga/crud_simples_padrao/service"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//CONFIGURAR A CONEXAO COM O BANCO DE DADOS MYSQL.
	dsn := "root:admin@tcp(localhost:3306)/crud_simples?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha com a conexao com o database")
	}
	
	db.AutoMigrate(&model.User{})


	//INICIALIZAR OS REPOSITORIOS, SERVICOS E HANDLERS
	userReposi := repository.NewUserRepository(db)
	userService := service.NewUserService(userReposi)
	UserHandler := handler.NewUserHandler(userService)


	// INICIAR O ECHO
	e := echo.New()


	//ROTAS
	e.GET("/users", UserHandler.GetUsers)

	//INICIAR O SERVIDOR
	e.Logger.Fatal(e.Start(":8080"))

}
