package main

import (
	"fmt"
	"learn/common"
	"learn/entity"
	"learn/handler"
	"learn/util"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm/clause"
)

func main() {
	var db = common.GetDB()
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Project{})
	var user = &entity.User{
		Name: util.Ref("dong"),
	}
	projectName := "xe hoi"
	db.Create(&user)
	fmt.Printf("%+v", *user)
	var project = &entity.Project{
		Name:   &projectName,
		UserId: user.ID,
		Attribute: &entity.ProjectAttribute{
			Tag: util.Ref("car"),
		},
	}
	db.Create(&project)
	fmt.Printf("%+v\n", *project)
	var users []entity.User

	db.Preload("Projects").Preload(clause.Associations).Find(&users)
	db.Preload("Projects").Find(&users)
	for _, user := range users {
		for _, project := range user.Projects {
			fmt.Printf("%+v\n", project.Attribute)
		}
	}
	fmt.Println("before update")
	fmt.Printf("%+v", project.Attribute.Tag)
	db.Model(&entity.Project{}).Where("id", project.ID).First(project)
	updateObj := make(map[string]interface{})
	updateObj["attribute"] = nil
	updateObj["name"] = "haha"
	db.Model(&project).Select([]string{"attribute", "name"}).Updates(updateObj)

	uById := entity.GetById(10)
	secretBytes := []byte("asecretkeyhaha")
	token, _ := util.SignWithClaims(entity.Claims{
		UserID: uById.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 100).Unix(),
		},
	}, secretBytes)
	fmt.Printf("TOKEN = %s", token)
	fmt.Printf("%+v", uById)

	handler.Run(8266)
}
