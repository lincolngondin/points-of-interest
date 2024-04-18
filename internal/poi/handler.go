package poi

import (
	"encoding/json"
	"log"
	"net/http"
)

type handler struct {

}

func NewHandler() *handler {
    return &handler{}
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
    
    response.WriteHeader(http.StatusOK)
    response.Write([]byte("OK!"))
    
}

func (hnd *handler) GetHandler(response http.ResponseWriter, request *http.Request) {
    values := request.URL.Query()
    log.Println(values)

    response.WriteHeader(http.StatusOK)
    response.Write([]byte("OK!!"))

}


