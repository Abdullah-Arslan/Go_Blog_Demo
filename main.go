package main

import (
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
	//post := admin_models.Post{}.Get("description= ?", "deneme")
	//fmt.Println(post.Title)
	//fmt.Println(admin_models.Post{}.GetAll())//Posts_model.go İÇERİSNDEKİ Func GetAll KISMINDAKİ VERİLERİ YAZDIRMAK İÇİN KULLANILAN KISIM
	//post := admin_models.Post{}.Get(1) //"admin_models.Post{}.Get()" .Get ile DEĞER ALIYORUZ. post.Update İLE ALDIĞIMIZ DEĞERİ GÜNCELLİYORUZ. "where 1" BUNUN ANLAMI ID Sİ 1 OLANI GETİR DEMEKTİR.
	//post.Update("title", "BASLIK KISMI DEĞİŞTİ")
	//post := admin_models.Post{}.Get(1)
	//post.Updates(admin_models.Post{Title: "Ptyton ile web programlama", Description: "Test"})
	//post := admin_models.Post{}.Get(1) //"where 1" İLE ID ATAMASI YAPIYORUZ YANİ Post_model.go İÇERİSİNDE "func Delete" FONKSİYONU İLE OLŞUTURULAN SİLME İŞLEMİNİN ID 1 DİR DİYEREK ID ATAMASI YAPIYORUZ.
	//post.Delete()

	//MODEL VIEW CONTROLLER
	//helpers.Include() //HELPERS İÇERİSİNDE OLUŞTURDUGUMUZ Include.html MAİN.GO İÇERİSİNE AKTARIMI
	http.ListenAndServe(":8080", config.Routes())
}
