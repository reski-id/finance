

<h3 align="center">Rest API <br>
<h5 align="center" >Golang Echo Clean Architecture <h5>
<br>
</h4>
<p align="left">
<h2>
  Content <br></h2>
  • Key Features <br>
  • Installing Using Github<br>
  • End Point<br>
  • Technologi that i use<br>
  • Contact me<br>
</p>


## 📱 Features

* register
* login
* Costumer
* Limit


## ⚙️ Installing and Runing from Github

installing and running the app from github repository <br>
To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/dl/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/reski-id/news_restAPI.git

# Install dependencies
$ go get

# Run the app
$ go run main.go
```

> **Note**
> Make sure you allready create database for this app see local `.env` file.


## 📜 End Point  

User
| Methode       | End Point      | used for            
| ------------- | -------------  | -----------        
| `POST`        | /register      | Register           
| `POST`        | /login         | Login              
| `GET`         | /my            | Get my Profile     
| `GET`         | /datausers     | Get All Registered User     
| `DELETE`      | /users         | Delete my Profile  
| `PUT`         | /users         | Update my Profile  

Costumer
| Methode       | End Point       | used for                
| ------------- | -------------   | -----------            
| `POST`        | /costumer       | Add Costumer 
| `GET`         | /costumer       | Get All Costumer         
| `GET`         | /costumer/:id       | Get single Costumer
| `DELETE`      | /costumer/:id   | Delete Costumer   
| `PUT`         | /costumer/:id   | Update Costumer   

Limit
| Methode       | End Point       | used for                
| ------------- | -------------   | -----------            
| `POST`        | /limit       | Add Limit         
| `GET`         | /limit/:id       | Get single Limit
| `DELETE`      | /limit/:id   | Delete Limit   
| `PUT`         | /limit/:id   | Update Limit

Transaction
| Methode       | End Point       | used for                
| ------------- | -------------   | -----------            
| `POST`        | /trx       | Add Transaction         
| `GET`         | /trx/:id       | Get Transaction

## 🛠️ Technologi

This software uses the following Tech:

- [Golang](https://go.dev/dl/)
- [Echo Framework](https://echo.labstack.com/)
- [Gorm](https://gorm.io/index.html)
- [mysql](https://www.mysql.com/)
- [Linux](https://www.linux.com/)
- [Docker](https://www.docker.com/)
- [Dockerhub](https://hub.docker.com/u/programmerreski)
- [Git Repository](https://github.com/reski-id)
- [JSON Web Tokens (JWT)](https://jwt.io/)

## 📱 Contact me
feel free to contact me ... 
- Email programmer.reski@gmail.com 
- [Linkedin](https://www.linkedin.com/in/reski-id)
- [Github](https://github.com/reski-id)
- Whatsapp <a href="https://wa.me/+6281261478432?text=Hello">Send WhatsApp Message</a>
