package main

import (
	"fmt"
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	//Post_model İÇERİSİNDEKİ MODEL DOSYALARIMIZI ÇALIŞTIRMAK İÇİN MAİN İÇERİSİNDE TANIMLAMA YAPIYORUZ. VE BU ŞEKİLDE DOSYALARI AÇĞIRARAK OLUŞTURUYORUZ.
	admin_models.Post{}.Migrate()
	/*
		admin_models.Post{
			Title: "Go ile web programlama",
			Slug:  "go-ile-web-programlama",
		}.Add()

	*/
	//BURADA POST_MODEL KISMINDA OLUŞTURDUGUMUZ GET METODUNA VERİLERİ POST ETME İŞLEMLERİNİ GÖRDÜK.
	//VE DATABASE KISMINDA VERİLERİ GET İLE NASIL ÇEKİLİR VE TERMİNALE YAZDIRILIR AŞAĞIDAKİ ÖRNEK İLE MEVCUT.
	post := admin_models.Post{}.Get("description= ?", "deneme")
	fmt.Println(post.Title)

	//MODEL VIEW CONTROLLER
	//helpers.Include() //HELPERS İÇERİSİNDE OLUŞTURDUGUMUZ Include.html MAİN.GO İÇERİSİNE AKTARIMI
	http.ListenAndServe(":8080", config.Routes())
}
