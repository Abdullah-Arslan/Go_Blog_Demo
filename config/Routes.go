package config

import (
	"github.com/julienschmidt/httprouter"
	admin "goblog/admin/controllers" //BURADA CONTROLLERS KISIMLARININ IMPORTLARINI AYIRMAK İÇİN BAŞLARINA TANIMLAMA YAPIYOURZ.
	"net/http"
)

// BURADAKİ *httprouter BİZE ROUTER DÖNECEK AMA POINTER OLARAK YÖNLENDİRMELERİN TAMAMINI BİZ BURADAN YAPACAĞIZ.
func Routes() *httprouter.Router {
	r := httprouter.New()

	//ADMIN
	//BLOGPOST
	r.GET("/admin", admin.Dashboard{}.Index)
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem) //DAHSBOARD İÇERİNDE OLUŞTURDUĞUMUZ NewItem FONKSİYONUNU BURADA BU  ŞEKİLDE ÇAĞIRIYORUZ. /admin/yeni-ekle",admin.Dashboard{}.NewItem BU ŞU DEMEK
	//admin içerisinde yeni-ekle KISMINA TIKLANDIĞINDA BUNU SEN admin.Dashboard{} içerinden NewItem a gideceksin demektir.

	// SERVE FILES

	//BURADA YAPILAN index.html SAYFASI İÇERİNDEKİ CSS VE JS DOSYALARINI ÇEKMESİ İÇİN ROUTE TANIMLAMASI YAPIYORUZ.
	//BU DİZİN GELİRSE /admin/assets/ SEN GİT http.Dir("admin/assets")) ADMİN ALTINDAN  ASSETS KLASÖRÜNE O DİZİNE YÖNLENDİR DİYORUZ.
	///admin/assets TAM HALİ İNDEX.HTML İÇERİSİNDE MEVCUT /admin/assets/css/sb-admin-2.min.css BUNUN GİBİ CSS VE JS HALİDE MEVCUT.
	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets"))

	return r

}
