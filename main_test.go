package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/LucasGabrielBecker/go-rest-api-gin/controllers"
	"github.com/LucasGabrielBecker/go-rest-api-gin/database"
	"github.com/LucasGabrielBecker/go-rest-api-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	return r
}

func CriaAlunoMock() {
	aluno := models.Aluno{
		Nome: "aluno teste",
		Cpf:  "18273847512",
		RG:   "182938178"}

	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestCheckStatusCodeSaudacao(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/someuser", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
	mockRes := `{"API diz: ":"E ai someuser, tudo beleza?"}`
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, mockRes, string(resBody))
}

func TestRequestFailIfNotParam(t *testing.T) {
	r := SetupTestRoutes()

	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, res.Code, http.StatusNotFound)
}

func TestListAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupTestRoutes()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscaAlunoPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupTestRoutes()
	r.GET("alunos/cpf/:cpf", controllers.BuscaPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/18273847512", nil)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestBuscaAlunoPorId(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupTestRoutes()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)

	searchPath := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("GET", searchPath, nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var alunoMock models.Aluno

	json.Unmarshal(res.Body.Bytes(), &alunoMock)

	assert.Equal(t, "aluno teste", alunoMock.Nome)
	assert.Equal(t, "18273847512", alunoMock.Cpf)
	assert.Equal(t, "182938178", alunoMock.RG)

}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()

	r := SetupTestRoutes()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	searchPath := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", searchPath, nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

}

func TestUpdateAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupTestRoutes()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	searchPath := "/alunos/" + strconv.Itoa(ID)
	aluno := models.Aluno{
		Nome: "aluno atualizado",
		Cpf:  "18273847500",
		RG:   "182938100"}

	alunoJSON, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", searchPath, bytes.NewBuffer(alunoJSON))

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var updatedUser models.Aluno
	json.Unmarshal(res.Body.Bytes(), &updatedUser)

	assert.Equal(t, "18273847500", updatedUser.Cpf)
	assert.Equal(t, "182938100", updatedUser.RG)
	assert.Equal(t, "aluno atualizado", updatedUser.Nome)

}
