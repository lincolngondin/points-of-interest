package poi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type serv interface {
	RegisterNewPOI(poi *POI) error
	GetAllPOI() ([]POI, error)
	GetAllPOIByDistance(refPoint *point, distanceMax uint64) ([]POI, error)
}

type handler struct {
	service serv
}

func NewHandler(svc serv) *handler {
	return &handler{service: svc}
}

type ResponseErr struct {
	Error string `json:"erro"`
}

func (hnd *handler) RegisterHandler(response http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(response)
	decoder := json.NewDecoder(request.Body)
	poi := NewDefaultPOI()
	decoderErr := decoder.Decode(&poi)
	if decoderErr != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ResponseErr{"Corpo da requisição inválido!"})
		return
	}
	// checa se o poi é válido, todos os campos são obrigatorios
	if !poi.IsValid() {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusBadRequest)
		encoder.Encode(ResponseErr{"Todos os campos do corpo da requisição são obrigatórios!"})
		return
	}

	registerErr := hnd.service.RegisterNewPOI(poi)
	if registerErr != nil {
		response.Header().Set("Content-Type", "application/json")
		response.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(ResponseErr{registerErr.Error()})
		return
	}

	response.WriteHeader(http.StatusCreated)
}

type queryValues struct {
	x     uint64
	y     uint64
	d_max uint64
}

func parseQueryParams(values url.Values) (*queryValues, error) {
	if !values.Has("x") && !values.Has("y") && !values.Has("d_max") {
		return nil, nil
	}
	xValue, convErr := strconv.ParseUint(values.Get("x"), 10, 64)
	yValue, convErr := strconv.ParseUint(values.Get("y"), 10, 64)
	dValue, convErr := strconv.ParseUint(values.Get("d_max"), 10, 64)
	if convErr != nil {
		return nil, convErr
	}
	return &queryValues{xValue, yValue, dValue}, nil
}

func (hnd *handler) GetHandler(response http.ResponseWriter, request *http.Request) {
	encoder := json.NewEncoder(response)
	queryValue, errQueryValues := parseQueryParams(request.URL.Query())
	if queryValue == nil && errQueryValues == nil {
		pois, err := hnd.service.GetAllPOI()
		if err != nil {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(ResponseErr{err.Error()})
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			encoder.Encode(pois)
		}
	} else {
		if errQueryValues != nil {
			response.WriteHeader(http.StatusNotFound)
			return
		}
		pois, err := hnd.service.GetAllPOIByDistance(newPoint(int64(queryValue.x), int64(queryValue.y)), queryValue.d_max)
		if err != nil {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(ResponseErr{err.Error()})
		} else {
			response.Header().Set("Content-Type", "application/json")
			response.WriteHeader(http.StatusOK)
			encoder.Encode(pois)
		}
	}
}
