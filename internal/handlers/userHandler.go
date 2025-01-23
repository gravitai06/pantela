package handlers

import (
	"context"
	"net/http"

	"pantela/internal/userServise"
	"pantela/internal/web/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	service *userService.Service
}

func NewUserHandler(service *userService.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	userList, err := h.service.GetAllUsers()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response []users.UserResponse
	for _, user := range userList {
		response = append(response, users.UserResponse{
			Id:        &user.ID,
			Email:     &user.Email,
			Password:  &user.Password,
			DeletedAt: &user.DeletedAt,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		})
	}

	return users.GetUsers200JSONResponse(response), nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	user := userService.User{
		Email:    *request.Body.Email,
		Password: *request.Body.Password,
	}

	if err := h.service.CreateUser(&user); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := users.UserResponse{
		Id:        &user.ID,
		Email:     &user.Email,
		Password:  &user.Password,
		DeletedAt: &user.DeletedAt,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return users.PostUsers201JSONResponse(response), nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	if err := h.service.DeleteUser(request.Id); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	updateData := make(map[string]interface{})
	if request.Body.Email != nil {
		updateData["email"] = *request.Body.Email
	}
	if request.Body.Password != nil {
		updateData["password"] = *request.Body.Password
	}

	user, err := h.service.UpdateUser(request.Id, updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := users.UserResponse{
		Id:        &user.ID,
		Email:     &user.Email,
		Password:  &user.Password,
		DeletedAt: &user.DeletedAt,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return users.PatchUsersId200JSONResponse(response), nil
}
