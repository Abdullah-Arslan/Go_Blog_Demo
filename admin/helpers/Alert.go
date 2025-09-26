package helpers

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

// BURADA ALET İŞLEMLERİ YAPIYORUZ VE ONUNLA İLGİLİ KOMUTLARI YAZIYORUZ
var store = sessions.NewCookieStore([]byte("123456"))

// SESSİON A EKLENECEK BİLGİYİ DIŞARIDAN ALACAĞIZ ONUDA message string İLE YAPIYORUZ ONU DA session.AddFlash(message) TANITIYORUZ.
func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := store.Get(r, "alert-go")
	if err != nil {
		fmt.Println(err)
		return err
	}
	session.AddFlash(message) //ALERT İŞLEMLERİNİ FLASH İLE GÖSTERİYORUZ.

	return session.Save(r, w) //SAVE İLEDE BU İŞLEMİ KAYDEDİYORUZ.

}

// ALERT İŞLEMİNİ KONTROL EDEBİLMEK İÇİN session DAN BUNU ÇAPIRACAK FONKSİYONA İHTİYAÇ VAR ONUDA AŞAĞIDA YAZIYORUZ.
func GetAlert(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	session, err := store.Get(r, "alert-go")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	data := make(map[string]interface{})
	flashes := session.Flashes()

	if len(flashes) > 0 { //flashes 0 DAN BÜYÜKSE YANİ İÇERİSİNDE VERİ VARSA AŞAĞIDAKİ İŞLEMLERİ YAP DİYORUZ.
		//DATA İÇERİSİNDE VERİ VARSA BUNU YAP DİYORUZ. İLK ELEMAN MESAJ OLARAK DÖNÜYOR.
		data["is_alert"] = true
		data["message"] = flashes[0]
	} else { //DATA İÇERİSİNDE VERİ YOKSA ELSE Yİ YAP
		data["is_alert"] = false
		data["message"] = nil
	}

	session.Save(r, w)

	//FLASH I GERİ ÇAĞIRMAK İÇİN AŞAĞIDAKİ KOMUTU KULLANIYORUZ.
	return data

}
