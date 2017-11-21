package consultant

import (

  "github.com/gin-gonic/gin"
  "github.com/appleboy/gin-jwt"
  // "github.com/gin-contrib/cors"
  // "fmt"
  "dibo-go-api/core"
  "dibo-go-api/library"
  "dibo-go-api/models"
  "net/http"
  // "time"
  "strconv"
)

/*
*   Consultant Part
*/
func CreateConsultant(c *gin.Context){
  password := encryptor.GetMD5Hash(c.PostForm("password"))
  phone, _ := strconv.Atoi(c.PostForm("phone"))
  consultant := models.Users{
    Username: c.PostForm("username"),
    Password: password,
    Fullname: c.PostForm("fullname"),
    Email: c.PostForm("email"),
    Phone: phone,
    Description: c.PostForm("description"),
    Status: c.PostForm("status"),
    Type: "consultant",
    Uid: 1};

  db:= database.ConnectionPostgres()
  db.Save(&consultant)
  c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "User successfully created"})
}

func FetchAllConsultant(c *gin.Context){
  var consultants []models.Users
  var _consultants []models.TransformedUsers

  limit := 20
  page := 1


  db:= database.ConnectionPostgres()
  result := map[string]interface{}{
    "type": "consultant",
  }


  if (len(c.Query("length")) > 0){
    limit, _ = strconv.Atoi(c.Query("length"))
  }

  if (len(c.Query("page")) > 0){
    page, _ = strconv.Atoi(c.Query("page"))
  }

  // calculate offset
  offset := (page - 1) * limit

  // check username
  if(len(c.Query("username"))>0){
    result["username"] = c.Query("username")
  }

  // check email
  if(len(c.Query("username"))>0){
    result["email"] = c.Query("email")
  }

  db.Where(result).Limit(limit).Offset(offset).Find(&consultants)
  if (len(consultants) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
    return
  }

  // remove password field
  for _, item := range consultants {
    _consultants = append(_consultants,
      models.TransformedUsers{
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
  var consultant models.Users
  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&consultant, consultant_id)

  if (consultant.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
    return
  }

  item := consultant

  result := map[string]interface{}{
    "ID": item.ID,
    "Username": item.Username,
    "Email": item.Email,
    "Fullname": item.Fullname,
    "Phone": item.Phone,
    "Description": item.Description,
    "Status": item.Status,
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : result})
}

func UpdateConsultant(c *gin.Context){
  var consultant models.Users
  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&consultant, consultant_id)

  if (consultant_id == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
    return
  }else{
    phone, _ := strconv.Atoi(c.PostForm("phone"))

    db.Model(&consultant).Update("username", c.PostForm("username"))
    db.Model(&consultant).Update("phone", phone)
    db.Model(&consultant).Update("fullname", c.PostForm("fullname"))
    db.Model(&consultant).Update("description", c.PostForm("description"))
    db.Model(&consultant).Update("email", c.PostForm("email"))

    c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Users successfully updated!"})
  }
}

func DeleteConsultant(c *gin.Context){
  var consultant models.Users
  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&consultant, consultant_id)

  if (consultant.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
    return
  }

  db.Delete(&consultant)
  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Users successfully deleted!"})
}



func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World.",
	})
}
