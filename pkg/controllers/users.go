package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/nihadtz/simple_user/models"
	usr "github.com/nihadtz/simple_user/pkg/user/repository"
	uss "github.com/nihadtz/simple_user/pkg/user/service"
	"github.com/nihadtz/simple_user/services"
)

type UserController struct {
	service uss.IUserService
}

func NewUserController() *UserController {
	repo := usr.NewUserRepository()
	service := uss.NewUserService(repo)
	resolver := &UserController{service}

	return resolver
}

func (ur UserController) UserCreate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		services.LogError("Error Unmarshal User", err)
		services.Renderer.Error(res, http.StatusInternalServerError, err.Error())
	}

	if len(user.Name) == 0 {
		err = errors.New("name is not set")
		services.LogError("Error", err)
		services.Renderer.Error(res, http.StatusBadRequest, err.Error())
		return
	}

	if user.YearOfBirth == 0 {
		err = errors.New("year of birth is not set")
		services.LogError("Error", err)
		services.Renderer.Error(res, http.StatusBadRequest, err.Error())
		return
	}

	userID, err := ur.service.RegisterUser(user)

	if err != nil {
		services.LogError("Error Creating User", err)
		services.Renderer.Error(res, http.StatusInternalServerError, "Couldn't create User")
		return
	}

	apiResponse := fmt.Sprintf("New user with id %d has been created", *userID)

	services.Renderer.Render(res, http.StatusOK, apiResponse)
}

func (ur UserController) RetrieveUserByID(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	userID, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		services.LogError("User ID not provided", err)
		services.Renderer.Error(res, http.StatusBadRequest, err.Error())
	}

	user, err := ur.service.RetrieveUserByID(int64(userID))

	if err != nil {
		services.LogError("Error reading User", err)
		services.Renderer.Error(res, http.StatusInternalServerError, "Couldn't read User")
	}

	apiResponse := fmt.Sprintf("This is %s born in %d", user.Name, user.YearOfBirth)

	services.Renderer.Render(res, http.StatusOK, apiResponse)
}

func (ur UserController) UserUpdate(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var user models.User

	err := json.NewDecoder(req.Body).Decode(&user)

	if err != nil {
		services.LogError("Error Unmarshal User", err)
		services.Renderer.Error(res, http.StatusInternalServerError, err.Error())
	}

	if len(user.Name) == 0 {
		err = errors.New("name is not set")
		services.LogError("Error", err)
		services.Renderer.Error(res, http.StatusBadRequest, err.Error())
		return
	}

	if user.YearOfBirth == 0 {
		err = errors.New("year of birth is not set")
		services.LogError("Error", err)
		services.Renderer.Error(res, http.StatusBadRequest, err.Error())
		return
	}

	if user.ID == 0 {
		err = errors.New("user ID is not provided")
		services.LogError("Error", err)
		services.Renderer.Error(res, http.StatusBadRequest, err.Error())
		return
	}

	_, err = ur.service.UpdateUser(user)

	if err != nil {
		services.LogError("Error Updating User", err)
		services.Renderer.Error(res, http.StatusInternalServerError, "Couldn't update User")
		return
	}

	apiResponse := "User has been updated"

	services.Renderer.Render(res, http.StatusOK, apiResponse)
}
