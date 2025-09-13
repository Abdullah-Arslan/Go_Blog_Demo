package models

//***  BURADA YAZILAN TÜM FONKSİYONLAR (FUNC) HEPSİ DATABASE (DB) İŞEMLERİ İÇİN YANİ DATABASE İÇERİSİNE VERİ EKLEM, GÜNCELLEME, SİLME vb. -
//*** BÜTÜN İŞLEMLERİNİ Posts_model.go DOSYASI İÇERİSİNDEN YAPIYORUZ YAZIYORUZ..

//*** ÖNEMLİ: BURADA YAZILAN BÜTÜN FONKSİYONLARI main.go İÇERİSİNDE ÇAĞIRIP ÇALIŞTIYORUZ.

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

// ÇEŞİTLİ VERİLER GELEBİLİR YANİ İNT, STRİNG vb. DEĞİŞKEN SAYIDA VERİ GELEBİLİR TÜM VERİLERİ TAMAMEN HER HANGİBİR TANIMLAMA YAPMADAN ALABİLMEK İÇİN "Get(where ...interface{})" where ...interface si KULLANIYORUZ.
func (post Post) Get(where ...interface{}) Post {
	db, err := gorm.Open(mysql.Open((Dns)), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}
	////YUKARIDAKİ func (post Post) İÇERİSİNDEKİ post u VERİYORUZ.
	db.First(&post, where...) //BURADA ... ÜÇ NOKTA YUKARIDAKİ WHERE İÇERİNDE DEĞİŞKEN SAYIDA VERİ ALIMINDA ... ÜÇ NOKTA İLE PAST EDİYORUZ.
	return post
}

func (post Post) GetAll(where ...interface{}) []Post {

	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//BİZDEN ARRAY İSTİYOR FİND AŞAĞIDAKİ TANIMLA İLE BUNU GERÇEKLEŞTİRİYORUZ.
	var posts []Post          //BUR ARRAY BİZE POST TUTACAK YANİ POST YAPISINDAN VERİ TUTACAK
	db.Find(&posts, where...) //HANGİ VERİLERİ ÇEKECEĞİNİ WHERE KOMUTU İLE VERİYORUZ VE &post İLE DE VERİLERİ POST İÇERİSİNE AKTARDI. Func GetAll TANIMLASINDA YAPILAN KOŞUL İLE ... İLEDE VERİLERİ POST TA GÖNDERİYORUZ.
	return posts              //RETURN POST DİYEREK POST İÇERİSİNE GİDEN VERİLERİ ALIYORUZ.

}

// value interface BİZE GELECEK OLAN VERİ HER TİPTEN OLABİLİR ANLAMINDADIR.
func (post Post) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return

	}
	//Modeli func içerisindeki posttan alıyoruz.
	db.Model(&post).Update(column, value) //UPDATE İŞLEMLERİNİ BU ŞEKİLDE GERÇEKLEŞTİRİYORUZ.

}

// DATABASE İÇERİSİNDE MULTİPLE BİR UPDATE İŞLEMİ İÇİN AŞAPIDAKİ KISMI KULLANIYORUZ. BİR METOT YAZIYORUZ.
func (post Post) Updates(data Post) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Updates(data) //BURADAKİ UPDATES YANİ BİR POST OLMASI GEREKİYOR ONUDA YUKARIDAKİ data Post İLE TANIMLIYORUZ.

}

// DELETE SİLME İŞLEMLERİNİ YANİ VERİTABININDAN VERİLERİ SİLMEK İÇİN KULLANIYORUZ.
func (post Post) Delete() {
	db, err := gorm.Open(mysql.Open((Dns)), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&post, post.ID)

}
