package handlers

import (
	"context"
	"net/http"

	"pantela/internal/taskServise"
	"pantela/internal/web/tasks"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service *taskServise.Service
}

func NewTaskHandler(service *taskServise.Service) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	taskList, err := h.service.GetAllTasks()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response []tasks.TaskResponse
	for _, task := range taskList {
		id := int(task.ID)
		taskName := task.Task
		isDone := task.IsDone
		response = append(response, tasks.TaskResponse{
			Id:     &id,
			Task:   &taskName,
			IsDone: &isDone,
		})
	}

	return tasks.GetTasks200JSONResponse(response), nil
}

func (h *TaskHandler) GetTasksByUserID(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	taskList, err := h.service.GetTasksForUser(uint(request.UserID))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var response []tasks.TaskResponse
	for _, task := range taskList {
		id := int(task.ID)
		taskName := task.Task
		isDone := task.IsDone
		response = append(response, tasks.TaskResponse{
			Id:     &id,
			Task:   &taskName,
			IsDone: &isDone,
		})
	}

	return tasks.GetTasks200JSONResponse(response), nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Request body is required")
	}

	task := taskServise.Task{
		Task:   request.Body.Task,
		IsDone: request.Body.IsDone,
	}

	if err := h.service.CreateTask(&task); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id := int(task.ID)
	response := tasks.TaskResponse{
		Id:     &id,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}

	return tasks.PostTasks201JSONResponse(response), nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	if err := h.service.DeleteTask(uint(request.Id)); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return tasks.DeleteTasksId204Response{}, nil
}

func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	updateData := make(map[string]interface{})
	if request.Body.Task != nil {
		updateData["task"] = *request.Body.Task
	}
	if request.Body.IsDone != nil {
		updateData["is_done"] = *request.Body.IsDone
	}

	task, err := h.service.UpdateTask(uint(request.Id), updateData)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	id := int(task.ID)
	response := tasks.TaskResponse{
		Id:     &id,
		Task:   &task.Task,
		IsDone: &task.IsDone,
	}

	return tasks.PatchTasksId200JSONResponse(response), nil
}
