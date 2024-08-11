package infrastructure

import (
	"database/sql"
	"github.com/MadAppGang/httplog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	user_usecase "newproject/domain/usecases/user"
	user_repository "newproject/infrastructure/repositories/user"
	"newproject/settings"
	"newproject/view/user"
)

func Setup(settings settings.Settings, router *mux.Router) error {
	db, err := setupDataBase(settings.DataBase)
	if err != nil {
		log.Println("[Setup] Error setupDataBase", err)
		return err
	}

	setupModules(db, router)

	return nil
}

func setupDataBase(databaseSettings settings.DataBase) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseSettings.GetDBSource())
	if err != nil {
		log.Println("[setupDataBase] Error Open", err)
		return nil, err
	}

	return db, nil
}

func setupModules(db *sql.DB, router *mux.Router) {
	router.Use(ContentTypeMiddleware)
	router.Use(httplog.Logger)

	userRepository := user_repository.NewRepository(db)
	userUseCases := user_usecase.NewUseCases(userRepository)
	user.NewHttpUserModule(userUseCases).Setup(router)

}

// retornar json para todas as funções
func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
