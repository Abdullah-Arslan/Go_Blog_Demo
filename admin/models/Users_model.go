package models

//***  BURADA YAZILAN TÜM FONKSİYONLAR (FUNC) HEPSİ DATABASE (DB) İŞEMLERİ İÇİN YANİ DATABASE İÇERİSİNE VERİ EKLEM, GÜNCELLEME, SİLME vb. -
//*** BÜTÜN İŞLEMLERİNİ Users_model.go DOSYASI İÇERİSİNDEN YAPIYORUZ YAZIYORUZ..

//*** ÖNEMLİ: BURADA YAZILAN BÜTÜN FONKSİYONLARI main.go İÇERİSİNDE ÇAĞIRIP ÇALIŞTIYORUZ.

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// go get "gorm.io/gorm" paketini çağırıyoruz.
// Burada struct yapısı oluşturuyoruz. Ve gorm.io yu model olarak çağırıyoruz.
// Buradaki gorm paketi bizim database olacak olan yapılarımızı MODELLERİMİZİ tanımladığiımız kısımdır. Bunlar ""Title,Slug,Description, Content, Picture_url string" ve "CategoryID int"
type User struct {
	gorm.Model
	Username, Password string
}

// MSQL.OPEN ÇALIŞMASI İÇİN  go get "gorm.io/driver/mysql" edilmesi gerekiyor.
// MODEL TANIMLAMASI YAPTIKTAN SONRA TOPLO OLUŞUMU MIGRATE OLUŞTURMAMIZ GEREKİYOR. BUNUN YAPMA NEDENİMİZ TAPLONUN OLUŞMASI İÇİN
func (user User) Migrate() { //BURADA mysql.Open KISMINI YÖNLENDİRMESİNİ DNS OLARAK VERİLMELİ
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//AUTOMİGRATE İLE DE TABLOLARIMIZI OLUŞTURUYORUZ.  YUKARIDAKİ user U BU ŞEKİLDE GÖNDERİYORUZ.
	db.AutoMigrate(&user)
}

// BUNDAN SONRA ADMİN PANELİNE VE DATABASE EKELENCEK OLAN KISIMLARININ FONKSİYONLARINI (FUNC) YAZIYORUZ.
// BUNUN İLKİ ADD İLE BAŞLIYORUZ.
// BURADA "func (user User)" user User İLE YAZILAN METOT DUR. YANİ YUKARIDA OLUŞTURDUĞUMUZ STRUCT YAPISINI ÇAĞIRIYORUZ.
func (user User) Add() {
	//EKLEME İŞLEMİNİ YAPACAĞIZ AMA ÖNCE YUKARIDAKİ GİBİ Migrate de GİBİ DATABASE Zİ OLUŞTURUYORUZ.
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	//BURADA db.Create(&user) DİYEREK &user İÇERİSİNE BAŞKA BİR YERDE OLUŞTURULAN TABLOLAMAYI GÖNDERİYORUZ. VE &user İLE BU TAPLO OLUŞTURULUYOR.
	/* ÖRNEK OLARAK DAHA ÖNCE main.go İÇERİSİNDE OLUŞTURULAN VE Users_model.go İÇERİSİNDE ADD FONKSİYONU İÇERİNDE CREAT OLUŞTURULAN TOPLO ÖRNEĞİ
	admin_models.User{
			Title: "Go ile web programlama",
			Slug:  "go-ile-web-programlama",
		}.Add()
	*/
	db.Create(&user)
}

// ÇEŞİTLİ VERİLER GELEBİLİR YANİ İNT, STRİNG vb. DEĞİŞKEN SAYIDA VERİ GELEBİLİR TÜM VERİLERİ TAMAMEN HER HANGİBİR TANIMLAMA YAPMADAN ALABİLMEK İÇİN "Get(where ...interface{})" where ...interface si KULLANIYORUZ.
func (user User) Get(where ...interface{}) User {
	db, err := gorm.Open(mysql.Open((Dns)), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}
	////YUKARIDAKİ func (user User) İÇERİSİNDEKİ user u VERİYORUZ.
	db.First(&user, where...) //BURADA ... ÜÇ NOKTA YUKARIDAKİ WHERE İÇERİNDE DEĞİŞKEN SAYIDA VERİ ALIMINDA ... ÜÇ NOKTA İLE PAST EDİYORUZ.
	return user
}

func (user User) GetAll(where ...interface{}) []User {

	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	//BİZDEN ARRAY İSTİYOR FİND AŞAĞIDAKİ TANIMLA İLE BUNU GERÇEKLEŞTİRİYORUZ.
	var users []User          //BUR ARRAY BİZE POST TUTACAK YANİ POST YAPISINDAN VERİ TUTACAK
	db.Find(&users, where...) //HANGİ VERİLERİ ÇEKECEĞİNİ WHERE KOMUTU İLE VERİYORUZ VE &user İLE DE VERİLERİ POST İÇERİSİNE AKTARDI. Func GetAll TANIMLASINDA YAPILAN KOŞUL İLE ... İLEDE VERİLERİ POST TA GÖNDERİYORUZ.
	return users              //RETURN POST DİYEREK POST İÇERİSİNE GİDEN VERİLERİ ALIYORUZ.

}

// value interface BİZE GELECEK OLAN VERİ HER TİPTEN OLABİLİR ANLAMINDADIR.
func (user User) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return

	}
	//Modeli func içerisindeki usertan alıyoruz.
	db.Model(&user).Update(column, value) //UPDATE İŞLEMLERİNİ BU ŞEKİLDE GERÇEKLEŞTİRİYORUZ.

}

// DATABASE İÇERİSİNDE MULTİPLE BİR UPDATE İŞLEMİ İÇİN AŞAPIDAKİ KISMI KULLANIYORUZ. BİR METOT YAZIYORUZ.
func (user User) Updates(data User) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&user).Updates(data) //BURADAKİ UPDATES YANİ BİR POST OLMASI GEREKİYOR ONUDA YUKARIDAKİ data User İLE TANIMLIYORUZ.

}

// DELETE SİLME İŞLEMLERİNİ YANİ VERİTABININDAN VERİLERİ SİLMEK İÇİN KULLANIYORUZ.
func (user User) Delete() {
	db, err := gorm.Open(mysql.Open((Dns)), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&user, user.ID)

}
