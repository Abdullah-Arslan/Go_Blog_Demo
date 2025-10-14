package helpers

import (
	"fmt"
	"goblog/admin/models"
	"net/http"
)

func SetUser(w http.ResponseWriter, r *http.Request, username string, password string) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		fmt.Println(err)
	}

	session.Values["username"] = username
	session.Values["password"] = password

	return session.Save(r, w)
}

// SESSİON DAKİ VERİYİ KONTROL ETMEK İÇİN AŞAĞIDAKİ FONKSİYONU YAZIYORUZ.
func CheckUser(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		return false //KULLANICI ADI DOĞRU İSE BU RADA TRUE DÖN DEĞİLSE FALSE ANLAMINDA VERDİK
	}

	username := session.Values["username"]
	password := session.Values["password"]

	//VERİ TABANINDAN KONTROLÜ BU ŞEKİLDE VERİYORUZ.
	user := models.User{}.Get("username = ? AND password = ?", username, password)

	if user.Username == username && user.Password == password {
		return true
	}
	SetAlert(w, r, "Lütfen Giriş Yapınız")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	return false //BURAYA KADAR OLAN KISIMDA GİRİŞ İŞLEMLERİNİN BAŞARILI OLUP OLMADIĞINI BOOL DEĞERİ İLE ALDIK.

}

// GİRİŞ YAPTIKTAN SONRA SESSİON DAKİ VERİLERİ SİLME İŞLEMİ BU ŞEKİLDE YAPILIYOR.
func RemoveUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "blog-user")
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1

	return session.Save(r, w)
}
