package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	Cpf  string `json:"cpf"`
	RG   string `json:"rg"`
}

type CriarAlunoInput struct {
	Nome string `json:"nome" binding:"required"`
	Cpf  string `json:"cpf" binding:"required"`
	RG   string `json:"rg" binding:"required"`
}
