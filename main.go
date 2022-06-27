package main

import (
	"fmt"
	"learn/common"
	"learn/handler"
	"learn/model"
	"learn/util"

	"gorm.io/gorm/clause"
)

func main() {
	var db = common.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Project{})
	var user = &model.User{
		Name: util.Ref("dong"),
	}
	projectName := "xe hoi"
	db.Create(&user)
	fmt.Printf("%+v", *user)
	var project = &model.Project{
		Name:   &projectName,
		UserId: user.ID,
		Attribute: &model.ProjectAttribute{
			Tag: util.Ref("car"),
		},
	}
	db.Create(&project)
	fmt.Printf("%+v\n", *project)
	var users []model.User

	db.Preload("Projects").Preload(clause.Associations).Find(&users)
	// db.Preload("Projects").Find(&users)
	for _, user := range users {
		for _, project := range user.Projects {
			fmt.Printf("%+v\n", project.Attribute)
		}
	}
	fmt.Println("before update")
	fmt.Printf("%+v", project.Attribute.Tag)
	db.Model(&model.Project{}).Where("id", project.ID).First(project)
	updateObj := make(map[string]interface{})
	updateObj["attribute"] = nil
	updateObj["name"] = "haha"
	db.Model(&project).Select([]string{"attribute", "name"}).Updates(updateObj)

	secretBytes := []byte("asecretkeyhaha")
	token, _ := util.SignWithClaims(secretBytes)
	fmt.Printf("TOKEN = %s", token)

	handler.Run(8266)
}
