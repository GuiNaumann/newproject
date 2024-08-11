package user

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"newproject/domain/entities"
	"newproject/domain/usecases/user"
	"newproject/view"
	"newproject/view/http_error"
)

type newHttpUserModule struct {
	useCases user.UseCases
}

func NewHttpUserModule(useCases user.UseCases) view.HttpModule {
	return &newHttpUserModule{
		useCases: useCases,
	}
}

func (n newHttpUserModule) Setup(router *mux.Router) {
	router.HandleFunc("/users", n.create).Methods("POST")
}
func (n newHttpUserModule) create(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http_error.HandleError(w, http_error.NewBadRequestError("Ocorreu um erro."))
		return
	}

	log.Println("bytes", b)

	var user entities.User
	if err = json.Unmarshal(b, &user); err != nil {
		log.Println("[create] Error json.Unmarshal", err)
		http_error.HandleError(w, http_error.NewBadRequestError("jogador não é valido."))
		return
	}
	log.Println(user)

	err = n.useCases.Create(r.Context(), user)
	if err != nil {
		log.Println("[create] Error Create", err)
		http_error.HandleError(w, err)
		return
	}

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Println("[create] Error Write", err)
		http_error.HandleError(w, http_error.NewInternalServerError(http_error.UnexpectedError))
		return
	}

}
