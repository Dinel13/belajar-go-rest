package controller

import (
	"github.com/julienschmidt/httprouter"
	"go-rest-api/helper"
	"go-rest-api/model/web"
	"go-rest-api/service"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService //interface not required to set pointer
}

func NewCategoryController(categoryService service.CategoryService) CategoryController  {
	return &CategoryControllerImpl {
		CategoryService: categoryService,
	}
}

func (controller CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)
	// before use helper
	//decoder := json.NewDecoder(r.Body)
	//err := decoder.Decode(categoryCreateRequest)
	//helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
	//before use helper
	//w.Header().Add("Content-Type", "application/json")
	//encoder := json.NewEncoder(w)
	//err = encoder.Encode(webResponse)
	//helper.PanicIfError(err)
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateUpdate := web.CategoryCreateUpdate{}
	helper.ReadFromRequestBody(r, &categoryCreateUpdate)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId) //change to int
	helper.PanicIfError(err)
	categoryCreateUpdate.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryCreateUpdate)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId) //change to int
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId) //change to int
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)

}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
