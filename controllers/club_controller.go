package controllers

import (
	"net/http"
	"pg-go-clean-architecture/web"
	"strconv"

	"github.com/labstack/echo/v4"
	"pg-go-clean-architecture/domain/club"
)

type ClubController struct {
	svc club.FootballClubService
}

func NewClubController(service club.FootballClubService) *ClubController {
	return &ClubController{
		svc: service,
	}
}

// GET /clubs/:id
func (c *ClubController) GetClubByID(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err.Error()))
	}

	club, err := c.svc.GetById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, web.SuccessResponse(club))
}

// GET /clubs
func (c *ClubController) GetClubs(ctx echo.Context) error {
	var request club.FootballClubPaginatedRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err.Error()))
	}

	clubs, err := c.svc.Get(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, web.SuccessResponse(clubs))
}

// POST /clubs
func (c *ClubController) CreateClub(ctx echo.Context) error {
	var fc club.FootballClub
	if err := ctx.Bind(&fc); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err.Error()))
	}

	if err := c.svc.Create(fc); err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, web.SuccessResponse("Football club created successfully"))
}

// PUT /clubs/:id
func (c *ClubController) UpdateClub(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err.Error()))
	}

	var fc club.FootballClub
	if err := ctx.Bind(&fc); err != nil {
		return ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err.Error()))
	}
	fc.Id = id

	if err := c.svc.Update(fc); err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, web.SuccessResponse("Football club updated successfully"))
}

// DELETE /clubs/:id
func (c *ClubController) DeleteClub(ctx echo.Context) error {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, web.ErrorResponse(err.Error()))
	}

	if err := c.svc.Delete(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, web.ErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, web.SuccessResponse("Football club deleted successfully"))
}
