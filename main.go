package main

import (
	"fmt"
	"learn/common"
	"learn/entity"
	"learn/handler"
	"learn/util"

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

	secretBytes := []byte("asecretkeyhaha")
	token, _ := util.SignWithClaims(secretBytes)
	fmt.Printf("TOKEN = %s", token)
	uById := entity.GetById(10)
	fmt.Printf("%+v", uById)

	handler.Run(8266)
}
