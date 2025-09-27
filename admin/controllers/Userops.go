package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"html/template"
	"net/http"
)

type Userops struct{}

// ADMİN GİRİŞ PANELİ İÇİN GEREKLİ OLAN KONTROLLERİN YAZILDIGI KISIM
func (userops Userops) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("userops/login")...)
	if err != nil {
		fmt.Println(err)
		return

	}
	view.ExecuteTemplate(w, "index", nil)

}

// ADMİN GİRİŞ EKRANI POST EDİLİYOR ODA userops altındaki login içindeki index.html İÇERİSİNDEKİ do_login İÇERİSİNE POST EDİLİYOR. BUNU ALMAK İÇİN AŞAĞIDAKİ KONTROLÜ YAZIYORUZ.
func (userops Userops) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//userops içerisindeki login içindeki index.html içerisinden uername ve password bilgilerini almak için aşağıdaki kısmı yazıyoruz.
	username := r.FormValue("username") //FORMVLUE İLE USERNAME VE
	password := r.FormValue("password") //PASSOWRD ALINIYOR.
	fmt.Println(username, password)
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther) //BU KISMI TAMAMLADIKTAN SONRA Routes.go İÇERİSİNDE ROUTE TANIMLAMASINI POST İLE YAPIYROUZ.

}
