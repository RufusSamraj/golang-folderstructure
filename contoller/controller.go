package contoller

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"main.go/constant"
	"main.go/dto"
	"main.go/service"
	"net/http"
	"strconv"
)

type CtrlStc struct {
	Service service.SvcInfc
}

type CtrlInfc interface {
}

func NewController(svc service.SvcInfc) CtrlStc {
	return CtrlStc{Service: svc}
}

func (c CtrlStc) GetStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rollNoStr := vars["rollNo"]
	rollNo, _ := strconv.Atoi(rollNoStr)

	response, svcErr := c.Service.GetStudent(rollNo)
	if svcErr != nil {
		constant.WriteResponse(w, 400, svcErr)
		return
	}

	constant.WriteResponse(w, 200, response)
}

func (c CtrlStc) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var request dto.Student

	jsonErr := json.NewDecoder(r.Body).Decode(&request)
	if jsonErr != nil {
		constant.WriteResponse(w, 400, jsonErr)
		return
	}

	validate := validator.New()
	valErr := validate.Struct(request)
	if valErr != nil {
		constant.WriteResponse(w, 400, "Missing mandatory fields in request")
		return
	}

	response, svcErr := c.Service.CreateStudent(request)
	if svcErr != nil {
		constant.WriteResponse(w, 400, svcErr)
		return
	}

	constant.WriteResponse(w, 200, response)
}
