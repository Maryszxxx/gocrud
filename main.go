package main

import (
	"log"

	"github.com/Maryszxxx/gocrud.git/src/controller/model/view/test/config/logger"
	"github.com/Maryszxxx/gocrud.git/src/controller/routes"
	"github.com/gin-gonic/gin"
)

// main é o ponto de entrada da aplicação.
// É aqui que o servidor HTTP é configurado e colocado no ar.
func main() {
	// Loga uma mensagem informando que a aplicação está subindo.
	// Usa o logger customizado (baseado no Zap) em vez do log padrão do Go,
	// para ter logs estruturados em JSON.
	logger.Info("Starting the application...")

	// Cria uma instância do Gin já com os middlewares padrão
	// (Logger e Recovery, que evita que a aplicação quebre em caso de panic).
	router := gin.Default()

	// Registra todas as rotas da aplicação (GET, POST, PUT, DELETE)
	// passando o RouterGroup raiz do Gin.
	// InitRoutes fica em routes.go e conecta cada URL a um controller.
	routes.InitRoutes(&router.RouterGroup)

	// Sobe o servidor HTTP na porta 8080.
	// router.Run() bloqueia a execução aqui — o programa fica escutando requisições.
	// Se der erro ao subir o servidor (ex: porta já em uso), o programa é encerrado com log.Fatal.
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
