package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// go get "gorm.io/gorm" paketini çağırıyoruz.
// Burada struct yapısı oluşturuyoruz. Ve gorm.io yu model olarak çağırıyoruz.
// Buradaki gorm paketi bizim database olacak olan yapılarımızı MODELLERİMİZİ tanımladığiımız kısımdır. Bunlar ""Title,Slug,Description, Content, Picture_url string" ve "CategoryID int"
type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Picture_url string
	CategoryID                                     int
}

// MSQL.OPEN ÇALIŞMASI İÇİN  go get "gorm.io/driver/mysql" edilmesi gerekiyor.
// MODEL TANIMLAMASI YAPTIKTAN SONRA TOPLO OLUŞUMU MIGRATE OLUŞTURMAMIZ GEREKİYOR. BUNUN YAPMA NEDENİMİZ TAPLONUN OLUŞMASI İÇİN
func (post Post) Migrate() { //BURADA mysql.Open KISMINI YÖNLENDİRMESİNİ DNS OLARAK VERİLMELİ
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//AUTOMİGRATE İLE DE TABLOLARIMIZI OLUŞTURUYORUZ.  YUKARIDAKİ post U BU ŞEKİLDE GÖNDERİYORUZ.
	db.AutoMigrate(&post)
}

// BUNDAN SONRA ADMİN PANELİNE VE DATABASE EKELENCEK OLAN KISIMLARININ FONKSİYONLARINI (FUNC) YAZIYORUZ.
// BUNUN İLKİ ADD İLE BAŞLIYORUZ.
// BURADA "func (post Post)" post Post İLE YAZILAN METOT DUR. YANİ YUKARIDA OLUŞTURDUĞUMUZ STRUCT YAPISINI ÇAĞIRIYORUZ.
func (post Post) Add() {
	//EKLEME İŞLEMİNİ YAPACAĞIZ AMA ÖNCE YUKARIDAKİ GİBİ Migrate de GİBİ DATABASE Zİ OLUŞTURUYORUZ.
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//BURADA db.Create(&post) DİYEREK &post İÇERİSİNE BAŞKA BİR YERDE OLUŞTURULAN TABLOLAMAYI GÖNDERİYORUZ. VE &post İLE BU TAPLO OLUŞTURULUYOR.
	/* ÖRNEK OLARAK DAHA ÖNCE main.go İÇERİSİNDE OLUŞTURULAN VE Posts_model.go İÇERİSİNDE ADD FONKSİYONU İÇERİNDE CREAT OLUŞTURULAN TOPLO ÖRNEĞİ
	admin_models.Post{
			Title: "Go ile web programlama",
			Slug:  "go-ile-web-programlama",
		}.Add()
	 */
	db.Create(&post)
}
