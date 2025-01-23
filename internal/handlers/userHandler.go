package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"pantela/internal/userServise"
	"pantela/internal/web/users"
)

type UserHandler struct {
	service *userService.Service
}

func NewUserHandler(service *userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx echo.Context) error {
	userList, err := h.service.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response []users.UserResponse
	for _, user := range userList {
		var deletedAt *time.Time
		if user.DeletedAt.Valid {
			deletedAt = &user.DeletedAt.Time
		}

		response = append(response, users.UserResponse{
			Id:        &user.ID,
			Email:     &user.Email,
			Password:  &user.Password,
			DeletedAt: deletedAt,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *UserHandler) PostUsers(ctx echo.Context) error {
	var request users.CreateUserRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	user := userService.User{
		Email:    request.Email,
		Password: request.Password,
	}

	if err := h.service.CreateUser(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	response := users.UserResponse{
		Id:        &user.ID,
		Email:     &user.Email,
		Password:  &user.Password,
		DeletedAt: deletedAt,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (h *UserHandler) DeleteUsersId(ctx echo.Context, id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.service.DeleteUser(uint(idUint)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *UserHandler) PatchUsersId(ctx echo.Context, id string) error {
	var request users.UpdateUserRequest
	if err := ctx.Bind(&request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	updateData := make(map[string]interface{})
	if request.Email != nil {
		updateData["email"] = *request.Email
	}
	if request.Password != nil {
		updateData["password"] = *request.Password
	}

	user, err := h.service.UpdateUser(uint(idUint), updateData)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	response := users.UserResponse{
		Id:        &user.ID,
		Email:     &user.Email,
		Password:  &user.Password,
		DeletedAt: deletedAt,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response)
}
