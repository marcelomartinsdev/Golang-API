package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o router com as configurações padrões do gin
	router := gin.Default()
	// Define a rota e a função que será executada quando a rota for acessada, essa funcao e chamada de handler e recebe um parametro do tipo gin.Context que contem todas as informações da requisição
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Inicia o servidor na porta 8080, listen and serve on
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
