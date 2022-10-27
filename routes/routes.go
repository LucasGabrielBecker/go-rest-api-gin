package routes

import (
	"github.com/LucasGabrielBecker/go-rest-api-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaPorCpf)

	r.POST("/alunos", controllers.CriaNovoAluno)

	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.Run()
}
