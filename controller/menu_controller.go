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

// MenuController adalah struct untuk user controller
type MenuController struct{}

// NewMenuController adalah fungsi pembuat untuk MenuController
func NewMenuController() *MenuController {
	return &MenuController{}
}

func (mc *MenuController) CreateMenu(w http.ResponseWriter, r *http.Request) {
	var menu model.MenuItem
	err := json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simpan user ke database
	err = config.DB.Create(&menu).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Hit CreateMenu Success")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menu)
}

func (mc *MenuController) GetAllMenus(w http.ResponseWriter, r *http.Request) {
	var pagination utils.Pagination
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))
	pagination.Page = page
	pagination.PerPage = perPage

	db := config.DB.Model(&model.MenuItem{})
	db, err := utils.Paginate(db, pagination)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var menus []model.MenuItem
	var count int64

	// Ambil semua data pengguna tanpa paginasi jika page=1 dan perPage=10 (default)
	if pagination.Page == 1 && pagination.PerPage == 10 {
		err = db.Find(&menus).Error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// Query untuk mengambil data pengguna dengan paginasi
		err = db.Find(&menus).Error
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
		"data":    menus,
	}

	log.Println("Hit GetAllMenus Success")
	json.NewEncoder(w).Encode(response)
}

func (mc *MenuController) GetMenuByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid menu ID", http.StatusBadRequest)
		return
	}

	menu := model.MenuItem{}
	err = config.DB.First(&menu, id).Error
	if err != nil {
		http.Error(w, "Menu not found", http.StatusNotFound)
		return
	}

	log.Println("Hit GetMenuByID Success")
	json.NewEncoder(w).Encode(menu)
}

func (mc *MenuController) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid menu ID", http.StatusBadRequest)
		return
	}

	var updateMenu model.MenuItem
	err = json.NewDecoder(r.Body).Decode(&updateMenu)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Dapatkan menu dari database berdasarkan ID
	var menu model.MenuItem
	err = config.DB.First(&menu, id).Error
	if err != nil {
		http.Error(w, "Menu not found", http.StatusNotFound)
		return
	}

	// Update struct menu dengan data dari updateMenu
	utils.UpdateStruct(&menu, updateMenu)

	// Simpan perubahan ke database
	err = menu.Update()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Hit UpdateMenu: ", updateMenu)
	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menu)
}

func (mc *MenuController) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Dapatkan menu dari database berdasarkan ID
	var menu model.MenuItem
	err = config.DB.First(&menu, id).Error
	if err != nil {
		http.Error(w, "Menu not found", http.StatusNotFound)
		return
	}

	// Hapus menu dari database
	err = menu.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Hit DeleteUser Success")
	// Kirim respons berhasil
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Menu deleted successfully"})
}
