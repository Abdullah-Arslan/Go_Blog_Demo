package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	//Post_model İÇERİSİNDEKİ MODEL DOSYALARIMIZI ÇALIŞTIRMAK İÇİN MAİN İÇERİSİNDE TANIMLAMA YAPIYORUZ. VE BU ŞEKİLDE DOSYALARI AÇĞIRARAK OLUŞTURUYORUZ.
	admin_models.Post{}.Migrate()
	admin_models.Post{
		Title: "Go ile web programlama",
		Slug:  "go-ile-web-programlama",
	}.Add()

	//MODEL VIEW CONTROLLER
	//helpers.Include() //HELPERS İÇERİSİNDE OLUŞTURDUGUMUZ Include.html MAİN.GO İÇERİSİNE AKTARIMI
	http.ListenAndServe(":8080", config.Routes())
}
