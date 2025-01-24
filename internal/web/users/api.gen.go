// Package users provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package users

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateUserRequest defines model for UpdateUserRequest.
type UpdateUserRequest struct {
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// UserResponse defines model for UserResponse.
type UserResponse struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at"`
	Email     *string    `json:"email,omitempty"`
	Id        *string    `json:"id,omitempty"`
	Password  *string    `json:"password,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody = CreateUserRequest

// PatchUsersIdJSONRequestBody defines body for PatchUsersId for application/json ContentType.
type PatchUsersIdJSONRequestBody = UpdateUserRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx echo.Context) error
	// Create a new user
	// (POST /users)
	PostUsers(ctx echo.Context) error
	// Delete a user
	// (DELETE /users/{id})
	DeleteUsersId(ctx echo.Context, id string) error
	// Update a user
	// (PATCH /users/{id})
	PatchUsersId(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUsers(ctx)
	return err
}

// PostUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostUsers(ctx)
	return err
}

// DeleteUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.DeleteUsersId(ctx, id)
	return err
}

// PatchUsersId converts echo context to params.
func (w *ServerInterfaceWrapper) PatchUsersId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PatchUsersId(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/users", wrapper.GetUsers)
	router.POST(baseURL+"/users", wrapper.PostUsers)
	router.DELETE(baseURL+"/users/:id", wrapper.DeleteUsersId)
	router.PATCH(baseURL+"/users/:id", wrapper.PatchUsersId)

}

type GetUsersRequestObject struct {
}

type GetUsersResponseObject interface {
	VisitGetUsersResponse(w http.ResponseWriter) error
}

type GetUsers200JSONResponse []UserResponse

func (response GetUsers200JSONResponse) VisitGetUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUsersRequestObject struct {
	Body *PostUsersJSONRequestBody
}

type PostUsersResponseObject interface {
	VisitPostUsersResponse(w http.ResponseWriter) error
}

type PostUsers201JSONResponse UserResponse

func (response PostUsers201JSONResponse) VisitPostUsersResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteUsersIdRequestObject struct {
	Id string `json:"id"`
}

type DeleteUsersIdResponseObject interface {
	VisitDeleteUsersIdResponse(w http.ResponseWriter) error
}

type DeleteUsersId204Response struct {
}

func (response DeleteUsersId204Response) VisitDeleteUsersIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type PatchUsersIdRequestObject struct {
	Id   string `json:"id"`
	Body *PatchUsersIdJSONRequestBody
}

type PatchUsersIdResponseObject interface {
	VisitPatchUsersIdResponse(w http.ResponseWriter) error
}

type PatchUsersId200JSONResponse UserResponse

func (response PatchUsersId200JSONResponse) VisitPatchUsersIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(ctx context.Context, request GetUsersRequestObject) (GetUsersResponseObject, error)
	// Create a new user
	// (POST /users)
	PostUsers(ctx context.Context, request PostUsersRequestObject) (PostUsersResponseObject, error)
	// Delete a user
	// (DELETE /users/{id})
	DeleteUsersId(ctx context.Context, request DeleteUsersIdRequestObject) (DeleteUsersIdResponseObject, error)
	// Update a user
	// (PATCH /users/{id})
	PatchUsersId(ctx context.Context, request PatchUsersIdRequestObject) (PatchUsersIdResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// GetUsers operation middleware
func (sh *strictHandler) GetUsers(ctx echo.Context) error {
	var request GetUsersRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.GetUsers(ctx.Request().Context(), request.(GetUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(GetUsersResponseObject); ok {
		return validResponse.VisitGetUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PostUsers operation middleware
func (sh *strictHandler) PostUsers(ctx echo.Context) error {
	var request PostUsersRequestObject

	var body PostUsersJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PostUsers(ctx.Request().Context(), request.(PostUsersRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUsers")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PostUsersResponseObject); ok {
		return validResponse.VisitPostUsersResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// DeleteUsersId operation middleware
func (sh *strictHandler) DeleteUsersId(ctx echo.Context, id string) error {
	var request DeleteUsersIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteUsersId(ctx.Request().Context(), request.(DeleteUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(DeleteUsersIdResponseObject); ok {
		return validResponse.VisitDeleteUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// PatchUsersId operation middleware
func (sh *strictHandler) PatchUsersId(ctx echo.Context, id string) error {
	var request PatchUsersIdRequestObject

	request.Id = id

	var body PatchUsersIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.PatchUsersId(ctx.Request().Context(), request.(PatchUsersIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PatchUsersId")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(PatchUsersIdResponseObject); ok {
		return validResponse.VisitPatchUsersIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
