package helpers

import (
	"path/filepath"
)

// BU KISIM BU ŞEKİLDE HEM TEPMLATE ALTINDAKİ DOSYALARI LİSTE YAPIP BİZE VERİYOR. HEM O ANDA LİST İN ALTINDAKİ HTML DOSYALARINIDA BİZE BİR LİSTE HALİNDE DÖNDÜRÜYOR.
func Include(path string) []string {
	//Glob bizim verdigimiz dizinde istenen kriterlere uyan dosyaları bir string olarak dizisini bize döndürüyor.
	files, _ := filepath.Glob("admin/views/templates/*.html")         //BAŞI FARK ETMEZ SONU HTML OLANLARI BANA DÖNDÜR DEMEKTİR. *html İLE TANIMLIYORUZ.
	path_files, _ := filepath.Glob("admin/views/" + path + "/*.html") //admin/views/dashboard/list/*html BURADA admin/views/ İLE *html HEPSİNDE ORTAK YOL /dashboard/list KSIMINI GÖNDERMEK YETERLİ OLUYOR
	for _, file := range path_files {
		files = append(files, file)

	}
	return files
}
