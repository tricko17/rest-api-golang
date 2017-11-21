package main

import (
  "dibo-go-api/core"
  "github.com/gin-gonic/gin"
  // "github.com/gin-contrib/cors"
  // "fmt"
  "dibo-go-api/models"
  "crypto/md5"
  "encoding/hex"
  "net/http"
  // "time"
  "strconv"
  // "log"
)

func GetMD5Hash(text string) string {
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func main(){
  db:= database.ConnectionPostgres()
  db.AutoMigrate(models.Consultant{})

  router := gin.Default()

  // router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

  v1 := router.Group("/api/v1/consultants")
  {
    v1.GET("/", FetchAllConsultant)
    v1.GET("/:id", FetchSingleConsultant)
    v1.POST("/", CreateConsultant)
    v1.PUT("/:id", UpdateConsultant)
    v1.DELETE("/:id", DeleteConsultant)
  }

	router.Run(":3000")
}

func CreateConsultant(c *gin.Context){
  password := GetMD5Hash(c.PostForm("password"))
  phone, _ := strconv.Atoi(c.PostForm("phone"))
  consultant := models.Consultant{
    Username: c.PostForm("username"),
    Password: password,
    Fullname: c.PostForm("fullname"),
    Email: c.PostForm("email"),
    Phone: phone,
    Description: c.PostForm("description"),
    Status: c.PostForm("status")};

  db:= database.ConnectionPostgres()
  db.Save(&consultant)
  c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Consultant created successfully"})
}

func FetchAllConsultant(c *gin.Context){
  var consultants []models.Consultant
  var _consultants []models.TransformedConsultant

  limit := 20
  // page := 1

  db:= database.ConnectionPostgres()

  if (len(c.Query("length")) > 0){
    limit, _ = strconv.Atoi(c.Query("length"))
  }

  person := make(map[string]string)

  // check username
  if(len(c.Query("lastname"))>0){
    person["username"] = c.Query("lastname")
  }

  // check email
  if(len(c.Query("email"))>0){
    person["email"] = c.Query("email")
  }

  db.Where(person).Limit(limit).Find(&consultants)
  if (len(consultants) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No consultants found!"})
    return
  }

  // remove password field
  for _, item := range consultants {
    _consultants = append(_consultants,
      models.TransformedConsultant{
        ID: item.ID,
        Username: item.Username,
        Email: item.Email,
        Fullname: item.Fullname,
        Phone: item.Phone,
        Description: item.Description,
        Status: item.Status})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _consultants})
}

func FetchSingleConsultant(c *gin.Context){
  var consultant models.Consultant

  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&consultant, consultant_id)

  if (consultant_id == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No consultant found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : consultant})
}

func UpdateConsultant(c *gin.Context){
  var consultant models.Consultant
  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&consultant, consultant_id)

  if (consultant_id == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No consultant found!"})
    return
  }else{
    phone, _ := strconv.Atoi(c.PostForm("phone"))

    db.Model(&consultant).Update("username", c.PostForm("username"))
    db.Model(&consultant).Update("phone", phone)
    db.Model(&consultant).Update("fullname", c.PostForm("fullname"))
    db.Model(&consultant).Update("description", c.PostForm("description"))
    db.Model(&consultant).Update("status", c.PostForm("status"))
    db.Model(&consultant).Update("email", c.PostForm("email"))

    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Consultant updated successfully!"})
  }
}

func DeleteConsultant(c *gin.Context){
  var consultant models.Consultant
  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&consultant, consultant_id)

  if (consultant_id == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No consultant found!"})
    return
  }

  db.Delete(&consultant)
  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Consultant deleted successfully!"})

}
