package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/junior/products-api/internal/products"
	"github.com/junior/products-api/pkg/web"
)

type CreateRequestDto struct {
	ID          int       `json:"id"`
	Nome        string    `json:"nome"`
	Cor         string    `json:"cor"`
	Preco       float64   `json:"preco"`
	Estoque     int       `json:"estoque"`
	Codigo      string    `json:"codigo"`
	Tamanho     string    `json:"tamanho"`
	Publicacao  bool      `json:"publicacao"`
	DataCriacao time.Time `json:"data_criacao"`
}

type UpdateRequestDto struct {
	Nome        string    `json:"nome" binding:"required"`
	Cor         string    `json:"cor" binding:"required"`
	Preco       float64   `json:"preco" binding:"required"`
	Estoque     int       `json:"estoque" binding:"required"`
	Codigo      string    `json:"codigo" binding:"required"`
	Tamanho     string    `json:"tamanho" binding:"required"`
	Publicacao  bool      `json:"publicacao" binding:"required"`
	DataCriacao time.Time `json:"data_criacao" binding:"required"`
}

type UpdateFieldsRequestDto struct {
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

type ProductHandler struct {
	service products.Service
}

func NewProduct(p products.Service) *ProductHandler {
	return &ProductHandler{
		service: p,
	}
}

func (c *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro interno"))
			return
		}

		if len(p) == 0 {
			ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, "Nenhum produto encontrado"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))
	}
}

func (c *ProductHandler) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id_prod, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
		}

		p, err := c.service.GetByID(id_prod)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro interno"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, p, ""))

	}
}

func (c *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
		}

		_, err = c.service.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Produto não encontrado"))
			return
		}

		var req UpdateRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "Não foi possível processar as instruções presentes"))
			return
		}

		// como estou usando o binding para verificar se algum dos campos estão vazios não precisa-se verificar.

		p, err := c.service.Update(id, req.Nome, req.Cor, req.Codigo, req.Tamanho, req.Estoque, req.Preco, req.Publicacao)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro interno"))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))

	}
}

func (c *ProductHandler) UpdateFields() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
		}

		_, err = c.service.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Produto não encontrado"))
			return
		}

		var req UpdateFieldsRequestDto
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "Não foi possível processar as instruções presentes"))
			return
		}

		// como estou usando o binding para verificar se algum dos campos estão vazios não precisa-se verificar.

		p, err := c.service.UpdateFields(id, req.Nome, req.Preco)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro interno"))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

func (c *ProductHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequestDto
		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, web.NewResponse(http.StatusUnprocessableEntity, nil, "Não foi possível processar as instruções presentes"))
			return
		}

		p, err := c.service.Create(req.Nome, req.Cor, req.Codigo, req.Tamanho, req.Estoque, req.Preco, req.Publicacao)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Erro interno"))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, p, ""))
	}
}

func (c *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "ID inválido"))
		}

		_, err = c.service.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Produto não encontrado"))
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "Produto não encontrado"))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, nil, "Produto removido com sucesso"))
	}
}
