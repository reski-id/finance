package mysql

import (
	"finance/config"
	"fmt"
	"log"

	costumerData "finance/feature/costumer/data"
	limitData "finance/feature/limit/data"
	tranData "finance/feature/transaction/data"
	userData "finance/feature/user/data"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.Username, cfg.Password, cfg.Address, cfg.Port, cfg.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}
	return db
}

func MigrateData(db *gorm.DB) {
	db.AutoMigrate(userData.User{}, costumerData.Costumer{}, limitData.Limit{}, tranData.Transaction{})
}

func SeedUsers(db *gorm.DB) {
	users := []userData.User{
		{
			Username: "jhon",
			Email:    "jhon@gmail.com",
			Password: "123456",
			FullName: "Jhon Wick",
			Role:     "user",
		},
		{
			Username: "jhane",
			Email:    "jhane@gmail.com",
			Password: "jhane123",
			FullName: "Jhane Do",
			Role:     "user",
		},
		{
			Username: "reski",
			Email:    "programmer.reski@gmail.com",
			Password: "reski1234",
			FullName: "Ahmad Reski",
			Role:     "admin",
		},
	}

	for _, user := range users {
		var existingUser userData.User
		result := db.Where("email = ?", user.Email).First(&existingUser)

		if result.Error != nil {
			// If there is no error, it means the user already exists in the database
			if result.RowsAffected > 0 {
				continue
			}

			// If the user doesn't exist, hash the password and insert the data
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal("Error hashing password: ", err)
			}
			user.Password = string(hashedPassword)

			result := db.Create(&user)
			if result.Error != nil {
				log.Fatal("Error seeding user: ", result.Error)
			}
		}
	}
}
