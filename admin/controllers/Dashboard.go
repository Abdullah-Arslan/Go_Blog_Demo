package controllers

import (
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"goblog/admin/helpers"
	"goblog/admin/models"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
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
	//GÖNDERİLEN FORMLARIN PANELDE LİSTENMESİ KISMINI BURADAN YAPIYORUZ.
	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll() //HEPSİNİ YANİ FORMDAN DOLDURULAN HERŞEYİ ÇEKİP EKRTANA EKLEYECEK
	data["Alert"] = helpers.GetAlert(w, r) //BİR ALERT VAR MI YOK MU BİR MESAJ VAR MI BİLGİSİ ARTIK BURADA TUTULUYOR.
	view.ExecuteTemplate(w, "index", data) //SİTENİN ÇALIŞTILMASI İÇİN EXECUTE ETMEK GEREKİYOR. ALT TEMALARA DA GİTMESİ İÇİN LİST İÇERİSİNDEKİ İNDEX.HTML İÇERİSİNDE  {{template "content" .}} KISIMINDA . NOKTA KOYUYORUZ.

}

// BURADA YAPTIĞIMIZ VİEWS İÇERİSİNDE DASHBOARD İÇERİSİNDE ADD KLASÖRÜ ALTINDAKİ HTML DOSYALARINI ÇAĞIRMA VE ÇALIŞMATA İŞLEMLERİ İÇİN FONKSİYON YAZIYORUZ.
// DASHBOARD İÇERİSİNDE YENİ BİR FONKSİYON YAZDIKTAN SONRA ROUTE TANIMLAMASI YAPIYORUZKİ ÇALIŞTIRABİLELİM.
func (dashboard Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//helpers.Include KULLANMAMIZIN NEDENİ BÜTÜN HTML DOSYALARINI INCLUDE.GO İÇERİSİNDE ÇAĞIRDIK ONUDA BURADA ÇAĞIRARAK ÇALIŞTIRIYORUZ.
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...) //DASHBOARD İÇERİNDE ADD KLASÖRÜ İÇERİNDEN HTML DOSYALARINI helpers.Include İÇERİNDEN AL DEMEKTİR.
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil) //DATA GÖNDERİMİ YAPMIYORUZ BU NEDENLE NİL OLARAK TANIMALMA YAPIYORUZ.

}

// add/content.html İÇERİSİNDEKİ  <form action="/admin/add" method="post" enctype="multipart/form-data"> "/admin/add" deki ADD i FONKSİYON OLARAK YAZIYORUZ
func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//OLUŞTURACAĞIMIZ KISIMDA FORMDAN GELECEK OLAN VERİLERİ ALACAĞIMIZ KISIMLARI OLUŞTURUYORUZ. BU KISMDA NELERİN OLDUGUNU add/content İÇERİNDE FORMDAKİ KISIMLARA BAKARAK OLUŞTURUYORUZ.
	title := r.FormValue("blog-title")
	slug := slug.Make(title) //SLUG ŞEKLİNDE OLUŞUM SAGLAMAK İÇİN BU KISIM KULLANILIYOR.
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category")) //strconv.Atoi BURADAKİ KISIM ŞU İŞE YARIYOR categoryID DATABASE DE int DEĞER ALIYOR FORMDAN GELEN VERİ İSE stinrg BİR DEĞERDİR. strconv.Atoi BİZE STRİNG DEĞERİ İNT E ÇEVİRMESİNİ SAĞLIYOR.
	content := r.FormValue("blog-content")

	//BURADA FORMDAKİ VERİLERİ ALACAĞIZ VE UPLOAD İŞLEMİ YAPACAĞIZ. BUNUN İÇİNDE uploads ADINDA KLASÖR OLUŞTURUYORUZ. ONUNDA İÇİNDE
	//Upload
	r.ParseMultipartForm(10 << 20) //BURADA YÜKLENECEK OLAN RESİMLERİN 10<<20 MB OLARAK BOYUTUNU BELİRLİYORUZ.
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	//AŞAĞIDA file İLE ALINAN DOSYAYININ İÇERİĞİNİ f İLE BAŞLAYAN BAŞKA BİR DOSYAYA AKTARIYORUZ.
	//BU KISIMDA OPENFİLE OLARAK DA YAPILACAK OLAN OKUMA,YAZMA VE OLUŞTURMA KISIMLARINI PİCTURE OLARAK YAPTIK
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = io.Copy(f, file) //BURADA file İÇERİSİNDEKİ DOSYAYI f İÇERİSİNE KOPYALA DİYORUZ.
	if err != nil {
		fmt.Println(err)
		return
	}
	//Upload End
	//UPLOAD İŞLEMERİ BURADA BİTTİ

	//VERİ TABANINA KAYDIMIZI EKLİYORUZ. YANİ YUKARIDA OLUŞTURULAN UPLOAD İŞLEMLERİ VE FORMDAN VERİ ALMA KISIMLARINI VERİ TABANINA GÖNDERME İŞLEMLERİNİ BURADAN YAPIYORUZ.
	models.Post{ //BU KISIMLARA FORMDAN GELECEK OLAN KISIMLARNIN TANIMLAMASINI YAPIYORUZ.
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Picture_url: "uploads/" + header.Filename,
		//BURADA YAZDIĞIMIZ KISIMLARIN HEPSİ VERİ TABANINA EKLECEK VE BU VERİLER BİZDEN FORMDAN ALACAK

	}.Add()                                           //TÜM BU İŞLEMLER BİTTİKTEN SONRA routes TANIMLAMASINI Routes.go İÇERİNE YAPIYORUZ.
	helpers.SetAlert(w, r, "Kayıt Başarıyla Eklendi") //YAPTIĞIMIZ ALERT İŞLEMİNİ BURADA TANIMLIYORUZ.
	//routes.go YA EKLEME İŞLEMİNİ YAPTIK AMA BURADAN ORAYA DÖNÜŞ YAPMAMIZ LAZIM AŞAĞIDAKİ KOD İLE DE DÖNÜŞ İŞLEMİNİ GERÇEKLEŞTİRİYORUZ.
	http.Redirect(w, r, "/admin", http.StatusSeeOther) //BU ŞEKİLDE TEKRAR İSTENEN GİRİLDİKTEN SONRA ANASAYFA YA DÖNÜŞ YAPILACAK. AKSİ TAKDİRDE BOŞ BİR SAYFA GÖRÜNÜR

}

// FROMDA EKLEDİĞİMİZ VERİLERİ SİLME İŞLEMİNİ BURADAKİ KISIM İLE GERÇEKLEŞTİRİYORUZ.
func (dashboard Dashboard) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	post := models.Post{}.Get(params.ByName("id")) //PARAMETRELERİ params.ByName İLE ALIYORUZ.
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}

// dashboar İÇERİSİNDE edit ADINDA YENİ BİR KLASÖR KOPYALIYORUZ BU KLASÖR add DEN KOPYALANDI VE AYNISI, BURADA YAPILACAK OLAN KISIM İSE EDİT İŞELMLERİ İÇİRİSİNDEKİ KISIMLARI KONTROLLÜ
func (dashboard Dashboard) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	//BURADAKİ İŞLEMİ ALABİŞLMEK İÇİŞN PARAMETRE ORAK EDİT GÖNDERİLMESİ GEREKİYOR ONUDA AŞAĞIDAKİ İŞLEMLER İLE YAPIYORUZ.
	data := make(map[string]interface{})
	data["Post"] = models.Post{}.Get(params.ByName("id")) //param.ByName ile EDİT KISMININ id si ALINIYOR.
	view.ExecuteTemplate(w, "index", data)
	//NOT: BURADA ÖNEMLİ OLAN YAZDIIMIZ EDİT FONKSİYONUNU edit/content.html İÇERİSİNDE DOGRU YERLERDE TANIMLAMAK MESELA "      <input type="text" name="blog-title" class="form-control" value="{{.Post.Title}}">" GİBİ DĞERLERİNİ DE AYNI ŞEKİLDE VERMEK GEREKİYOR

}
func (dashboard Dashboard) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//BURADA NEYİ UPDATE EDECEKSEK ONU ALIYORUZ ÖNCE
	post := models.Post{}.Get(params.ByName("id")) //BURADA FORMDAKİ HANGİ YERİ DEĞİŞTİRECEKSEK ONU BURAYA ALMIŞ OLDUK. HANGİ KAYDI GÜNCELLEYECEKSEK İD ÜZERİNDEN ALDIK GETİRDİK POST İÇERİSNE ATTIK
	title := r.FormValue("blog-title")
	slug := slug.Make(title) //SLUG ŞEKLİNDE OLUŞUM SAGLAMAK İÇİN BU KISIM KULLANILIYOR.
	description := r.FormValue("blog-desc")
	categoryID, _ := strconv.Atoi(r.FormValue("blog-category")) //strconv.Atoi BURADAKİ KISIM ŞU İŞE YARIYOR categoryID DATABASE DE int DEĞER ALIYOR FORMDAN GELEN VERİ İSE stinrg BİR DEĞERDİR. strconv.Atoi BİZE STRİNG DEĞERİ İNT E ÇEVİRMESİNİ SAĞLIYOR.
	content := r.FormValue("blog-content")
	is_selected := r.FormValue("is_selected")
	var picture_url string

	//BU KISIMDA BİZ YENİ BİR DOSYA SEÇTİK Mİ FİLE UPDATE YAPILACAK MI YAPILMAYACAK MI ONU İF DÖNGÜSÜ İLE BELİRLİYORUZ.
	if is_selected == "1" { //EĞER is_selected 1 RE EŞİTSE UPLAOD İŞLEMİ YAPMIŞIZDIR DEMEKTİR.
		//is_selected burada değişecek

		r.ParseMultipartForm(10 << 20) //BURADA YÜKLENECEK OLAN RESİMLERİN 10<<20 MB OLARAK BOYUTUNU BELİRLİYORUZ.
		file, header, err := r.FormFile("blog-picture")
		if err != nil {
			fmt.Println(err)
			return
		}
		//AŞAĞIDA file İLE ALINAN DOSYAYININ İÇERİĞİNİ f İLE BAŞLAYAN BAŞKA BİR DOSYAYA AKTARIYORUZ.
		//BU KISIMDA OPENFİLE OLARAK DA YAPILACAK OLAN OKUMA,YAZMA VE OLUŞTURMA KISIMLARINI PİCTURE OLARAK YAPTIK
		f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = io.Copy(f, file)                  //BURADA file İÇERİSİNDEKİ DOSYAYI f İÇERİSİNE KOPYALA DİYORUZ.
		picture_url = "uploads/" + header.Filename //picture_url BURADA GÜNCELLENİYOR
		//UPLOAD YAPTIK BUNU SİLMEMİZDE GEREKİYOR ONUN İÇİN AŞAĞIDAKİ KISMI KULLANIYORUZ
		os.Open(post.Picture_url) //BU BİZİM ÖNCEKİ GÖRSELİ SİLMİŞ OLACAK

	} else {
		//is_selected burada aynen kalacak
		picture_url = post.Picture_url
	}
	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  categoryID,
		Content:     content,
		Picture_url: picture_url,
	})
	http.Redirect(w, r, "/admin/edit/"+params.ByName("id"), http.StatusSeeOther)
}
