package main

import (
  "dibo-go-api/core"
  "dibo-go-api/models"
)

func main(){

  db:= database.ConnectionPostgres()
  db.AutoMigrate(models.Users{}, models.Brands{}, models.Projects{})
}
