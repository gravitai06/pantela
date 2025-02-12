package handlers

import (
	"context"
	"errors"
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

		userID := strconv.FormatUint(uint64(user.ID), 10)

		response = append(response, users.UserResponse{
			Id:        &userID,
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

	userID := strconv.FormatUint(uint64(user.ID), 10)

	response := users.UserResponse{
		Id:        &userID,
		Email:     &user.Email,
		Password:  &user.Password,
		DeletedAt: deletedAt,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
	return users.PostUsers201JSONResponse(response), nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	userID, err := strconv.ParseUint(request.Id, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	if err := h.service.DeleteUser(uint(userID)); err != nil {
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

	userID, err := strconv.ParseUint(request.Id, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.service.UpdateUser(uint(userID), updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	userIDStr := strconv.FormatUint(uint64(user.ID), 10)

	response := users.UserResponse{
		Id:        &userIDStr,
		Email:     &user.Email,
		Password:  &user.Password,
		DeletedAt: deletedAt,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}

	return users.PatchUsersId200JSONResponse(response), nil
}

func strPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func (h *UserHandler) GetUsersIdTasks(ctx context.Context, request users.GetUsersIdTasksRequestObject) (users.GetUsersIdTasksResponseObject, error) {
	if h.service == nil {
		return nil, errors.New("service is nil")
	}

	userID, err := strconv.ParseUint(request.Id, 10, 64)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	tasks, err := h.service.GetTasksForUser(uint(userID))
	if err != nil {
		return nil, err
	}

	var taskResponses []users.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, users.TaskResponse{
			Id:     strPtr(strconv.FormatUint(uint64(task.ID), 10)),
			IsDone: boolPtr(task.IsDone),
			Task:   strPtr(task.Task),
		})
	}

	return users.GetUsersIdTasks200JSONResponse(taskResponses), nil
}
