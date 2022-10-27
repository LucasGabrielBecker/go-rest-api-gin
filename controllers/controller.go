package controllers

import (
	"net/http"

	"github.com/LucasGabrielBecker/go-rest-api-gin/database"
	"github.com/LucasGabrielBecker/go-rest-api-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(http.StatusOK, gin.H{
		"API diz: ": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var input models.CriarAlunoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	aluno := models.Aluno{Nome: input.Nome, Cpf: input.Cpf, RG: input.RG}

	database.DB.Create(&aluno)
	c.JSON(http.StatusCreated, &aluno)

}

func ExibeAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	alunoId := c.Params.ByName("id")
	database.DB.First(&aluno, alunoId)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaPorCpf(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Params.ByName("cpf")

	database.DB.Where(&models.Aluno{Cpf: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, &aluno)
}
