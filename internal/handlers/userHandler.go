package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

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

	return users.GetUsers200JSONResponse(response), nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	user := userService.User{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}

	if err := h.service.CreateUser(&user); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
	if request.Body.Email != nil && *request.Body.Email != "" {
		updateData["email"] = *request.Body.Email
	}
	if request.Body.Password != nil && *request.Body.Password != "" {
		updateData["password"] = *request.Body.Password
	}

	user, err := h.service.UpdateUser(request.Id, updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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

	return users.PatchUsersId200JSONResponse(response), nil
}

func (h *UserHandler) GetTasksForUser(ctx context.Context, request users.GetTasksForUserRequestObject) (users.GetTasksForUserResponseObject, error) {
	userID, err := strconv.Atoi(request.Id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	tasks, err := h.service.GetTasksForUser(uint(userID))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response []users.TaskResponse
	for _, task := range tasks {
		id := int(task.ID)
		taskName := task.Task
		isDone := task.IsDone
		response = append(response, users.TaskResponse{
			Id:     &id,
			Task:   &taskName,
			IsDone: &isDone,
		})
	}

	return users.GetTasksForUser200JSONResponse(response), nil
}
