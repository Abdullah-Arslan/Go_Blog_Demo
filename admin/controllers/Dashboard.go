package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"html/template"
	"net/http"
)

// TEMANIN ANASAYFASI OLACAK KISIM BURASIDIR.
// ÖNCE TYPE STRUCTLARI OLUŞTURUYORUZ.
// DASHBOARD STRUCT TI OLUŞTURUYORUZ. FUNC OLARAK TANIMLAMA YAPABİLMEK İÇİN OLUŞTURULMASI GEREKİYOR.
type Dashboard struct{}

// BURASI BİZİM CONTROLLERIMIZ YANİ VIEW,TEMPLATE EXECUTE ETTİĞİMİZ KISIMDIR. YANİ ONLARI ÇALIŞTIRIP KOMUT VERİLEN YERDİR.
func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil) //SİTENİN ÇALIŞTILMASI İÇİN EXECUTE ETMEK GEREKİYOR.

}
