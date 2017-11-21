package client

import (

  "github.com/gin-gonic/gin"

  // "fmt"
  "dibo-go-api/core"
  "dibo-go-api/library"
  "dibo-go-api/models"
  "net/http"
  "strconv"
)

/*
*   Client Part
*/
func CreateClient(c *gin.Context){
  cns, _ := strconv.Atoi(c.PostForm("consultant"))
  password := encryptor.GetMD5Hash(c.PostForm("password"))
  phone, _ := strconv.Atoi(c.PostForm("phone"))
  client := models.Users{
    Username: c.PostForm("username"),
    Password: password,
    Fullname: c.PostForm("fullname"),
    Email: c.PostForm("email"),
    Phone: phone,
    Description: c.PostForm("description"),
    Status: c.PostForm("status"),
    Type: "client",
    Uid: cns};

  db:= database.ConnectionPostgres()
  db.Save(&client)
  c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Client created successfully"})
}

func FetchAllClient(c *gin.Context){
  var clients []models.Users
  var _clients []models.TransformedUsers

  limit := 20
  page := 1


  db:= database.ConnectionPostgres()
  result := map[string]interface{}{
    "type": "client",
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

  db.Where(result).Limit(limit).Offset(offset).Find(&clients)
  if (len(clients) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
    return
  }

  // remove password field
  for _, item := range clients {
    _clients = append(_clients,
      models.TransformedUsers{
        ID: item.ID,
        Username: item.Username,
        Email: item.Email,
        Fullname: item.Fullname,
        Phone: item.Phone,
        Description: item.Description,
        Status: item.Status})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _clients})
}

func FetchAllClientBaseConsultant(c *gin.Context){
  var clients []models.Users
  var _clients []models.TransformedUsers

  limit := 20
  page := 1
  consultant_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  result := map[string]interface{}{
    "type": "client",
    "uid": consultant_id}


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

  db.Where(result).Limit(limit).Offset(offset).Find(&clients)
  if (len(clients) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No users found!"})
    return
  }

  // remove password field
  for _, item := range clients {
    _clients = append(_clients,
      models.TransformedUsers{
        ID: item.ID,
        Username: item.Username,
        Email: item.Email,
        Fullname: item.Fullname,
        Phone: item.Phone,
        Description: item.Description,
        Status: item.Status})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _clients})
}
