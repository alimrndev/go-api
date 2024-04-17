package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/alimrndev/go-api/config"
	"github.com/alimrndev/go-api/model"
	"github.com/alimrndev/go-api/utils"
	"github.com/gorilla/mux"
)

// UserController adalah struct untuk user controller
type UserController struct{}

// NewUserController adalah fungsi pembuat untuk UserController
func NewUserController() *UserController {
	return &UserController{}
}

// CreateUser adalah handler untuk membuat user baru
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set password
	err = user.SetPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Simpan user ke database
	err = config.DB.Create(&user).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Hit CreateUser")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers adalah handler untuk mendapatkan semua user
func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var pagination utils.Pagination
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	pagination.Page = page
	pagination.PerPage = perPage

	db := config.DB.Model(&model.User{})
	db, err := utils.Paginate(db, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var users []model.User
	var count int64

	// Ambil semua data pengguna tanpa paginasi jika page=1 dan perPage=10 (default)
	if pagination.Page == 1 && pagination.PerPage == 10 {
		err = db.Find(&users).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Query untuk mengambil data pengguna dengan paginasi
		err = db.Find(&users).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Hitung jumlah total data untuk paginasi
		err = config.DB.Model(&model.User{}).Count(&count).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Menyiapkan respons
	response := map[string]interface{}{
		"page":    pagination.Page,
		"perPage": pagination.PerPage,
		"count":   count,
		"data":    users,
	}

	log.Println("Hit GetAllUsers")
	json.NewEncoder(w).Encode(response)
}

func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user := model.User{}
	err = config.DB.First(&user, id).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	log.Println("Hit GetUserByID", user.Password)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updateUser model.User
	err = json.NewDecoder(r.Body).Decode(&updateUser)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validasi format email
	if updateUser.Email != "" {
		if !utils.IsValidEmail(updateUser.Email) {
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}
	}

	// Dapatkan user dari database berdasarkan ID
	var user model.User
	err = config.DB.First(&user, id).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update struct user dengan data dari updateUser
	utils.UpdateStruct(&user, updateUser)

	// Set password jika ada password yang disediakan dalam permintaan
	if updateUser.Password != "" {
		err = user.SetPassword(updateUser.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Simpan perubahan ke database
	err = user.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Hit UpdateUser: ", updateUser)
	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Dapatkan user dari database berdasarkan ID
	var user model.User
	err = config.DB.First(&user, id).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Hapus user dari database
	err = user.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Hit DeleteUser")
	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
