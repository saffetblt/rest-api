# TANIM
Uygulama REST API ile endpointlere gelen HTTP isteklerine database üzerinde işlem yaparak cevap vermektedir.

## Kullanılan Teknolojiler
Bu uygulamada **Gin Web Framework** ile **Gorm ORM Kütüphanesi** kullanılmaktadır.

## Bağımlılıklar
```
go get -u github.com/gin-gonic/gin
go get -u github.com/jinzhu/gorm
go get github.com/mattn/go-sqlite3

```
## Uygulamanın başlatılması
Uygulama /backend klasörü içinde aşağıdaki komut ile başlatılır.
```
go run main.go
```

## Uygulamanın Çalıştığı Port
Uygulama **8085** portunu dinlemektedir.

## Uygulamanın Dinlediği Uç Noktalar

| Method | URI         | Fonksiyon        |
|--------|-------------|------------------|
| GET    | /user/    | main.GetUsers   |
| GET    | /user/:id | main.GetUser    |
| POST   | /user     | main.CreateUser |
| DELETE | /user/:id | main.DeleteUser |
| PUT    | /user/:id | main.UpdateUser |

## Örnek bir JSON çıktısı

```
{
    "ID": 1,
    "CreatedAt": "2019-04-20T22:02:04.648969408+03:00",
    "UpdatedAt": "2019-04-20T22:02:04.648969408+03:00",
    "DeletedAt": null,
    "username": "saffetblt",
    "name": "Saffet",
    "surname": "BULUT",
    "password": "12345",
    "email": "saffetblt@gmail.com",
    "phone": "555 444 33 22",
    "age": 25,
    "country": "Turkey",
    "city": "Ankara"
}
```

## Örnek HTTP istekleri
GET
```
192.168.0.12:8085/user/
192.168.0.12:8085/user/5
```
POST
```
192.168.0.12:8085/user/
```
DELETE
```
192.168.0.12:8085/user/4
```
PUT
```
192.168.0.12:8085/user/3
```

## Öneriler
Http istekleri için Postman programı kullanılmaktadır.

.db uzantılı database dosyasını görüntülemek için DB Browser for SQLite kullanılmaktadır.


