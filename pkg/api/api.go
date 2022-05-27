package api

import (
	"encoding/json"
	"errors"
	cfg "github.com/StepanShevelev/task/pkg/config"
	mydb "github.com/StepanShevelev/task/pkg/db"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func InitBackendApi(config *cfg.Config) {
	http.HandleFunc("/API/create_user", apiCreateUser)
	http.HandleFunc("/API/create_category", apiCreateCategory)
	http.HandleFunc("/API/create_pet", apiCreatePet)

	http.HandleFunc("/API/update_user", apiUpdateUser)
	http.HandleFunc("/API/update_category", apiUpdateCategory)
	http.HandleFunc("/API/update_pet", apiUpdatePet)

	http.HandleFunc("/API/get_user", apiGetUser)
	http.HandleFunc("/API/get_category", apiGetCategory)
	http.HandleFunc("/API/get_pet", apiGetPet)

	http.HandleFunc("/API/get_users", apiGetAllUsers)
	http.HandleFunc("/API/get_categories", apiGetAllCategories)
	http.HandleFunc("/API/get_pets", apiGetAllPets)

	http.HandleFunc("/API/delete_user", apiDeleteUser)
	http.HandleFunc("/API/delete_category", apiDeleteCategory)
	http.HandleFunc("/API/delete_pet", apiDeletePet)

	http.HandleFunc("/API/user_add_category", apiUserAddCategory)
	http.HandleFunc("/API/user_delete_category", apiUserDeleteCategory)

	http.HandleFunc("/API/user_add_pet", apiUserAddPet)
	http.HandleFunc("/API/user_delete_pet", apiUserDeletePet)

	http.HandleFunc("/API/pharaoh_user", apiPharaohUser)

}

func apiCreateUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}

	CreateUser(r)
}

func CreateUser(r *http.Request) {

	usr := &mydb.User{}
	err := json.NewDecoder(r.Body).Decode(usr)
	if err != nil {
		return
	}
	defer r.Body.Close()
	mydb.Client.Save(usr)

}

func apiCreateCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}
	CreateCategory(r)
}

func CreateCategory(r *http.Request) {

	cat := &mydb.Category{}
	err := json.NewDecoder(r.Body).Decode(cat)
	if err != nil {
		return
	}
	defer r.Body.Close()
	mydb.Client.Save(cat)
}

func apiCreatePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodPOST(w, r) {
		return
	}
	CreatePet(r)
}

func CreatePet(r *http.Request) {

	pet := &mydb.Pet{}
	err := json.NewDecoder(r.Body).Decode(pet)
	if err != nil {
		return
	}
	defer r.Body.Close()
	mydb.Client.Save(pet)
}

func apiUpdateUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := parseId(w, r)
	if !okId {
		return
	}

	usr, okUsr := getUserById(id, w)
	if !okUsr {
		return
	}

	name := ""
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		return
	}

	usr.Name = name

	mydb.Client.Save(&usr)
}

func apiUpdateCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := parseId(w, r)
	if !okId {
		return
	}

	cat, okCat := getCategoryById(id, w)
	if !okCat {
		return
	}

	name := ""
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		return
	}

	cat.Name = name

	mydb.Client.Save(&cat)
}

func apiUpdatePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodPUT(w, r) {
		return
	}

	id, okId := parseId(w, r)
	if !okId {
		return
	}

	pet, okPet := getUserById(id, w)
	if !okPet {
		return
	}

	name := ""
	err := json.NewDecoder(r.Body).Decode(&name)
	if err != nil {
		return
	}

	pet.Name = name

	mydb.Client.Save(&pet)
}

func apiGetUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	userId, okId := parseId(w, r)
	if !okId {
		return
	}

	user, okUser := getUserById(userId, w)
	if !okUser {
		return
	}
	sendData(user, w)
}

func getUserById(userId int, w http.ResponseWriter) (*mydb.User, bool) {

	user, err := mydb.Client.DbGetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return user, true
}

func apiGetCategory(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	userId, okId := parseId(w, r)
	if !okId {
		return
	}

	user, okCategory := getCategoryById(userId, w)
	if !okCategory {
		return
	}
	sendData(user, w)
}

func getCategoryById(categoryId int, w http.ResponseWriter) (*mydb.Category, bool) {

	cat, err := mydb.Client.DbGetCategoryByID(categoryId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return cat, true
}

func apiGetPet(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	petId, okId := parseId(w, r)
	if !okId {
		return
	}

	pet, okPet := getPetById(petId, w)
	if !okPet {
		return
	}
	sendData(pet, w)
}

func getPetById(petId int, w http.ResponseWriter) (*mydb.Pet, bool) {

	pet, err := mydb.Client.DbGetPetByID(petId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return pet, true
}

func apiGetAllUsers(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	users, okUsers := getUsers(w)
	if !okUsers {
		return
	}
	sendData(users, w)
}

func getUsers(w http.ResponseWriter) ([]mydb.User, bool) {

	users, err := mydb.Client.DbGetUsers()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return users, true
}

func apiGetAllCategories(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	categories, okCategories := getCategories(w)
	if !okCategories {
		return
	}
	sendData(categories, w)
}
func getCategories(w http.ResponseWriter) ([]mydb.Category, bool) {

	categories, err := mydb.Client.DbGetCategories()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return categories, true
}

func apiGetAllPets(w http.ResponseWriter, r *http.Request) {
	if !isMethodGET(w, r) {
		return
	}
	users, okUsers := getPets(w)
	if !okUsers {
		return
	}
	sendData(users, w)
}
func getPets(w http.ResponseWriter) ([]mydb.Pet, bool) {

	pets, err := mydb.Client.DbGetPets()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return nil, false
	} else if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return nil, false
	}
	return pets, true
}

func apiDeleteUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	usr, okUsr := getUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Client.Unscoped().Delete(&usr)
	w.WriteHeader(http.StatusOK)
}

func apiDeleteCategory(w http.ResponseWriter, r *http.Request) {

	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	cat, okCat := getCategoryById(id, w)
	if !okCat {
		return
	}

	mydb.Client.Unscoped().Delete(&cat)
	w.WriteHeader(http.StatusOK)
}

func apiDeletePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	pet, okPet := getPetById(id, w)
	if !okPet {
		return
	}

	mydb.Client.Unscoped().Delete(&pet)
	w.WriteHeader(http.StatusOK)
}

func apiUserAddCategory(w http.ResponseWriter, r *http.Request) {

}

func apiUserDeleteCategory(w http.ResponseWriter, r *http.Request) {

}

func apiUserAddPet(w http.ResponseWriter, r *http.Request) {

}

func apiUserDeletePet(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	usr, okUsr := getUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Client.Where("user_id = ?", usr.ID).Unscoped().Delete(&mydb.Pet{})
	w.WriteHeader(http.StatusOK)
}

func apiPharaohUser(w http.ResponseWriter, r *http.Request) {
	if !isMethodDELETE(w, r) {
		return
	}
	id, okId := parseId(w, r)
	if !okId {
		return
	}

	usr, okUsr := getUserById(id, w)
	if !okUsr {
		return
	}

	mydb.Client.Unscoped().Delete(&usr)
	mydb.Client.Where("user_id = ?", usr.ID).Unscoped().Delete(&mydb.Pet{})
	w.WriteHeader(http.StatusOK)
}

func parseId(w http.ResponseWriter, r *http.Request) (int, bool) {
	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "arguments params are missing"}`))
		return 0, false
	}
	userId, err := strconv.Atoi(keys[0])
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "can't pars id"}`))
		return 0, false
	}
	return userId, true
}

func isMethodGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func isMethodDELETE(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func isMethodPOST(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func isMethodPUT(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func sendData(data interface{}, w http.ResponseWriter) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "can't marshal json"}`))
		return
	}
	w.Write(b)
	w.WriteHeader(http.StatusOK)
}
