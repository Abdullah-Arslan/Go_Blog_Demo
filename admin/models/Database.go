package models

// Burada yapılan işlem gorm.io sitesinden dns olacak yapıyı alıyoruz. Local de kullanılan dns ne ise onun root ve user işlemleri yapıp database adını veriyoruz.
var Dns string = "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
