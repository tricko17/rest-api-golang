package brand

import (

  "github.com/gin-gonic/gin"

  // "fmt"
  "dibo-go-api/core"
  "dibo-go-api/models"
  "strconv"
  "net/http"
)

func CreateBrand(c *gin.Context){
  client, _ := strconv.Atoi(c.PostForm("client"))
  brands := models.Brands{
    Name: c.PostForm("name"),
    Description: c.PostForm("description"),
    Status: c.PostForm("status"),
    Client: client};

  db := database.ConnectionPostgres()
  db.Save(&brands)
  c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Brand successfully created"})
}

func FetchAllBrand(c *gin.Context){
  var brands []models.Brands
  var _brands []models.TransformedBrands

  limit := 20
  page := 1


  db:= database.ConnectionPostgres()
  if (len(c.Query("length")) > 0){
    limit, _ = strconv.Atoi(c.Query("length"))
  }

  if (len(c.Query("page")) > 0){
    page, _ = strconv.Atoi(c.Query("page"))
  }

  // calculate offset
  offset := (page - 1) * limit

  db.Limit(limit).Offset(offset).Find(&brands)
  if (len(brands) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No brands found!"})
    return
  }

  // remove password field
  for _, item := range brands {
    _brands = append(_brands,
      models.TransformedBrands{
        ID: item.ID,
        Name: item.Name,
        Client: item.Client,
        Description: item.Description,
        Status: item.Status})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _brands})
}

func FetchBrandBaseClient(c *gin.Context){
  var brands []models.Brands
  var _brands []models.TransformedBrands

  client_id, _ := strconv.Atoi(c.Param("id"))
  limit := 20
  page := 1

  result := map[string]interface{}{
    "client": client_id}

  db:= database.ConnectionPostgres()
  if (len(c.Query("length")) > 0){
    limit, _ = strconv.Atoi(c.Query("length"))
  }

  if (len(c.Query("page")) > 0){
    page, _ = strconv.Atoi(c.Query("page"))
  }

  // calculate offset
  offset := (page - 1) * limit

  db.Where(result).Limit(limit).Offset(offset).Find(&brands)
  if (len(brands) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No brands found!"})
    return
  }

  // remove password field
  for _, item := range brands {
    _brands = append(_brands,
      models.TransformedBrands{
        ID: item.ID,
        Name: item.Name,
        Client: item.Client,
        Description: item.Description,
        Status: item.Status})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _brands})
}

func FetchSingleBrand(c *gin.Context){
  var brands models.Brands
  brand_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&brands, brand_id)

  if (brands.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No brand found!"})
    return
  }

  item := brands

  result := map[string]interface{}{
    "ID": item.ID,
    "Name": item.Name,
    "Client": item.Client,
    "Description": item.Description,
    "Status": item.Status,
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : result})
}

func UpdateBrand(c *gin.Context){
  var brands models.Brands
  brand_id, _ := strconv.Atoi(c.Param("id"))


  db:= database.ConnectionPostgres()
  db.First(&brands, brand_id)

  if (brands.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No brand found!"})
    return
  }

  client, _ := strconv.Atoi(c.PostForm("client"))

  db.Model(&brands).Update("name", c.PostForm("name"))
  db.Model(&brands).Update("description", c.PostForm("description"))
  db.Model(&brands).Update("status", c.PostForm("status"))
  db.Model(&brands).Update("client", client)

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Brands successfully updated!"})
}

func DeleteBrand(c *gin.Context){
  var brand models.Brands
  brand_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&brand, brand_id)

  if (brand.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No brands found!"})
    return
  }

  db.Delete(&brand)
  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Brands successfully deleted!"})
}
