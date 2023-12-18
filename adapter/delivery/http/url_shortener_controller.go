package http

import (
	"errors"
	"github.com/Arvin619/url-shortener/application/usecase"
	errors2 "github.com/Arvin619/url-shortener/domain/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlShortenerController struct {
	takeOutOriginalUrlUseCase *usecase.TakeOutOriginalUrlUseCase
	addShortUrlUseCase        *usecase.AddShortUrlUseCase
}

func NewUrlShortenerController(takeOutOriginalUrlUseCase *usecase.TakeOutOriginalUrlUseCase, addShortUrlUseCase *usecase.AddShortUrlUseCase) *UrlShortenerController {
	return &UrlShortenerController{
		takeOutOriginalUrlUseCase: takeOutOriginalUrlUseCase,
		addShortUrlUseCase:        addShortUrlUseCase,
	}
}

type TakeOutOriginalUrlResponse struct {
	ShortId     string `json:"short_id"`
	OriginalUrl string `json:"original_url"`
}

func (u *UrlShortenerController) TakeOutOriginalUrl(ctx *gin.Context) {
	shortId := ctx.Param("shortId")
	cmd := &usecase.TakeOutOriginalUrlCommand{
		ShortId: shortId,
	}
	r, err := u.takeOutOriginalUrlUseCase.Execute(ctx, cmd)
	if err != nil {
		if errors.Is(err, errors2.ErrUrlShortenerNotFound) {
			ctx.Error(NewError(http.StatusNotFound, err.Error()))
		} else {
			ctx.Error(NewError(http.StatusBadRequest, err.Error()))
		}
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, TakeOutOriginalUrlResponse{
		ShortId:     r.ShortId,
		OriginalUrl: r.OriginalUrl,
	})
}

type AddShortUrlRequest struct {
	OriginalUrl string `json:"original_url"`
}

type AddShortUrlResponse struct {
	ShortId     string `json:"short_id"`
	OriginalUrl string `json:"original_url"`
}

func (u *UrlShortenerController) AddShortUrl(ctx *gin.Context) {
	var req AddShortUrlRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(ErrBadRequestBodyMustBeJson)
		ctx.Abort()
		return
	}
	cmd := &usecase.AddShortUrlCommand{
		OriginalUrl: req.OriginalUrl,
	}

	r, err := u.addShortUrlUseCase.Execute(ctx, cmd)
	if err != nil {
		ctx.Error(NewError(http.StatusBadRequest, err.Error()))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, AddShortUrlResponse{
		ShortId:     r.ShortId,
		OriginalUrl: r.OriginalUrl,
	})
}
