package v1

import (
	"log"
	"net/http"

	requests "github.com/samayamnag/boilerplate/app/http/requests/v1"
	resources "github.com/samayamnag/boilerplate/app/http/resources/v1"
	"github.com/samayamnag/boilerplate/app/http/responses"
	"github.com/samayamnag/boilerplate/app/models/mongo"
	"github.com/samayamnag/boilerplate/app/repositories"
	"github.com/samayamnag/boilerplate/app/util"

	"github.com/gorilla/mux"	
	"github.com/urfave/negroni"
)

func userIndex(service repositories.UserUseCaseInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading users"
		var data []*mongo.User
		var err error
		name := r.URL.Query().Get("name")
		switch {
		case name == "":
			data, err = service.FindAll()
		default:
			data, err = service.Search(name)
		}
		if err != nil && err != mongo.ErrNotFound {
			responses.RespondWithMsg(w, http.StatusInternalServerError, errorMessage)
			return
		}

		if data == nil {
			responses.RespondWithMsg(w, http.StatusNotFound, errorMessage)
			return
		}
		users := resources.UserCollection{}
		for _, u := range data {
			users = append(users, formatResource(u))
		}

		responses.RespondJSON(w, http.StatusOK, users)
	})
}

func userAdd(service repositories.UserUseCaseInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		e := requests.CreateUserRequest(r)
		vErr := map[string]interface{}{"errors": e}
		errSize := len(e)
		if errSize == 0 {
			errorMessage := "Error adding user"
			var user mongo.User
			user.Email = r.FormValue("email")
			user.Password = r.FormValue("password")
			user.FullName = r.FormValue("full_name")
			user, err := service.Store(user)
			if err != nil {
				log.Println(err.Error())
				responses.RespondWithMsg(w, http.StatusInternalServerError, errorMessage)
				return
			}
			responses.RespondJSON(w, http.StatusCreated, formatResource(&user))
		} else {
			responses.RespondJSON(w, http.StatusCreated, vErr)
		}
	})
}

func userFind(service repositories.UserUseCaseInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading user"
		vars := mux.Vars(r)
		id := vars["id"]
		data, err := service.Find(mongo.StringToID(id))
		if err != nil && err != mongo.ErrNotFound {
			responses.RespondWithMsg(w, http.StatusInternalServerError, errorMessage)
			return
		}

		if data == nil {
			responses.RespondWithMsg(w, http.StatusNotFound, errorMessage)
			return
		}

		responses.RespondJSON(w, http.StatusOK, formatResource(data))
	})
}

func userUpdate(service repositories.UserUseCaseInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]
		errorMessage := "Error reading user"

		user, err := service.Find(mongo.StringToID(id))
		if err != nil && err != mongo.ErrNotFound {
			responses.RespondWithMsg(w, http.StatusInternalServerError, errorMessage)
			return
		}

		if user == nil {
			responses.RespondWithMsg(w, http.StatusNotFound, errorMessage)
			return
		}

		// Validate input
		e := requests.UpdateUserRequest(r)
		vErr := map[string]interface{}{"errors": e}
		errSize := len(e)
		if errSize == 0 {
			errorMessage := "Error updating user"
			user.Email = r.FormValue("email")
			user.Password = r.FormValue("password")
			user.FullName = r.FormValue("full_name")
			user, err := service.Update(user)
			if err != nil {
				log.Println(err.Error())
				responses.RespondWithMsg(w, http.StatusInternalServerError, errorMessage)
				return
			}
			responses.RespondJSON(w, http.StatusOK, formatResource(user))
		} else {
			responses.RespondJSON(w, http.StatusCreated, vErr)
		}
	})
}

func userDelete(service repositories.UserUseCaseInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing user"
		vars := mux.Vars(r)
		id := vars["id"]
		err := service.Delete(mongo.StringToID(id))
		if err != nil {
			responses.RespondWithMsg(w, http.StatusInternalServerError, errorMessage)
			return
		}
		responses.RespondWithMsg(w, http.StatusOK, "Deleted successfully")
	})
}

func formatResource(user *mongo.User) resources.UserResource {
	return resources.UserResource{
		Id:               user.ID.String(),
		Email:            user.Email,
		FullName:         user.FullName,
		Timestamp:        util.FormatMongoDate(user.CreatedAt),
		UpdatedTimestamp: util.FormatMongoDate(user.UpdatedAt),
	}
}

//MakeUserHandlers make url handlers
func MakeUserHandlers(r *mux.Router, n negroni.Negroni, service repositories.UserUseCaseInterface) {
	r.Handle("/users", n.With(
		negroni.Wrap(userIndex(service)),
	)).Methods("GET", "OPTIONS").Name("userIndex")

	r.Handle("/users", n.With(
		negroni.Wrap(userAdd(service)),
	)).Methods("POST", "OPTIONS").Name("userAdd")

	r.Handle("/users/{id}", n.With(
		negroni.Wrap(userFind(service)),
	)).Methods("GET", "OPTIONS").Name("userFind")

	r.Handle("/users/{id}", n.With(
		negroni.Wrap(userUpdate(service)),
	)).Methods("PUT", "OPTIONS").Name("userUpdate")

	r.Handle("/users/{id}", n.With(
		negroni.Wrap(userDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("userDelete")

	// v1 handlers
	r.Handle("/v1/users", n.With(
		negroni.Wrap(userIndex(service)),
	)).Methods("GET", "OPTIONS").Name("userIndex")

	r.Handle("/v1/users", n.With(
		negroni.Wrap(userAdd(service)),
	)).Methods("POST", "OPTIONS").Name("userAdd")

	r.Handle("/v1/users/{id}", n.With(
		negroni.Wrap(userFind(service)),
	)).Methods("GET", "OPTIONS").Name("userFind")

	r.Handle("/v1/users/{id}", n.With(
		negroni.Wrap(userUpdate(service)),
	)).Methods("PUT", "OPTIONS").Name("userUpdate")

	r.Handle("/v1/users/{id}", n.With(
		negroni.Wrap(userDelete(service)),
	)).Methods("DELETE", "OPTIONS").Name("userDelete")
}
