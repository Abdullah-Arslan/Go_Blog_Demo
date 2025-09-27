package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
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
	//BUNLARI KULLANABİLMEK İÇİN ÖNCE Users_model.go MODEL DOSYASINI OLUŞTURUYORUZ.
	username := r.FormValue("username")                                           //FORMVLUE İLE USERNAME VE
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(r.FormValue("password")))) //PASSOWRD ALINIYOR. BURADA YAPILAN İŞLEM GİRİLEN PARALOYI ŞİFRELİ OLARAK BİZE VERMESİ İÇİN SHA256 İLE ŞİFRELEME YAPIYORUZ.

	user := models.User{}.Get("username =? AND password=?", username, password)

	//ŞİMDİ BURADA FORMDAN GELEN VERİNİN KONTROL EDİLMESİ İŞLEMİNİ YAPACAĞIZ. BUNUDA İF KOMUTU YAPIYORUZ.
	if user.Username == username && user.Password == password {
		//LOGİN GİRİŞ YAPILAN KISIM
		http.Redirect(w, r, "/admin", http.StatusSeeOther) //BU KISIM GİRİŞ YAPTIKTAN SONRA BİZİ ADMİN ANASAYFASINA YÖNLENDİRİYOR.
	} else {
		//DENIED ŞİFRE YADA KULLANICI ADI YANLIŞ OLURSA GİRMESİNE İZİN VERİLMEYEN YER YANİ admin/login e GERİ GÖNDERİLECEK YER.
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther) //BU KISIM BİZİ YANLIŞ ŞİFRE YADA KULLANICI ADI GİRDİKTEN SONRA TEKRAR GİRİŞ PANEL SAYFASINA YÖNLENDİRİYOR.
	}

	//http.Redirect(w, r, "/admin/login", http.StatusSeeOther) //BU KISMI TAMAMLADIKTAN SONRA Routes.go İÇERİSİNDE ROUTE TANIMLAMASINI POST İLE YAPIYROUZ.

}
