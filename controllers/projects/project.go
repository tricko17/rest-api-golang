package project

import (

  "github.com/gin-gonic/gin"

  // "fmt"
  "dibo-go-api/core"
  "dibo-go-api/models"
  "net/http"
  "strconv"
)

func CreateProject(c *gin.Context){
  brand, _ := strconv.Atoi(c.PostForm("brand"))
  projects := models.Projects{
    Code: c.PostForm("code"),
    Name: c.PostForm("name"),
    Description: c.PostForm("description"),
    Status: c.PostForm("status"),
    Start: c.PostForm("start"),
    End: c.PostForm("end"),
    Brand: brand};

  db := database.ConnectionPostgres()
  db.Save(&projects)
  c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Project successfully created"})
}

func FetchAllProject(c *gin.Context){
  var projects []models.Projects
  var _projects []models.TransformedProjects

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

  db.Limit(limit).Offset(offset).Find(&projects)
  if (len(projects) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No projects found!"})
    return
  }

  // remove password field
  for _, item := range projects {
    _projects = append(_projects,
      models.TransformedProjects{
        ID: item.ID,
        Code: item.Name,
        Name: item.Client,
        Description: item.Description,
        Status: item.Status,
        Start: item.Start,
        End: item.End,
        Brand: item.Brand})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _projects})
}

func FetchProjectBaseBrand(c *gin.Context){
  var projects []models.Projects
  var _projects []models.TransformedProjects

  brand_id, _ := strconv.Atoi(c.Param("id"))
  limit := 20
  page := 1

  result := map[string]interface{}{
    "brand": brand_id}

  db:= database.ConnectionPostgres()
  if (len(c.Query("length")) > 0){
    limit, _ = strconv.Atoi(c.Query("length"))
  }

  if (len(c.Query("page")) > 0){
    page, _ = strconv.Atoi(c.Query("page"))
  }

  // calculate offset
  offset := (page - 1) * limit

  db.Where(result).Limit(limit).Offset(offset).Find(&projects)
  if (len(projects) <= 0) {
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No projects found!"})
    return
  }

  // remove password field
  for _, item := range projects {
    _projects = append(_projects,
      models.TransformedProjects{
        ID: item.ID,
        Code: item.Name,
        Name: item.Client,
        Description: item.Description,
        Status: item.Status,
        Start: item.Start,
        End: item.End,
        Brand: item.Brand})
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _projects})
}

func FetchSingleProject(c *gin.Context){
  var projects models.Projects
  brand_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&projects, brand_id)

  if (projects.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No projects found!"})
    return
  }

  item := projects

  result := map[string]interface{}{
    "ID": item.ID,
    "Code": item.Code,
    "Name": item.Name,
    "Description": item.Description,
    "Status": item.Status,
    "Start": item.Start,
    "End": item.End,
    "Brand": item.Brand,
  }

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : result})
}

func UpdateProject(c *gin.Context){
  var projects models.Projects
  project_id, _ := strconv.Atoi(c.Param("id"))


  db:= database.ConnectionPostgres()
  db.First(&projects, project_id)

  if (projects.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No brand found!"})
    return
  }

  brand, _ := strconv.Atoi(c.PostForm("brand"))
  db.Model(&projects).Update("code", c.PostForm("code"))
  db.Model(&projects).Update("name", c.PostForm("name"))
  db.Model(&projects).Update("description", c.PostForm("description"))
  db.Model(&projects).Update("start", c.PostForm("start"))
  db.Model(&projects).Update("end", c.PostForm("end"))
  db.Model(&projects).Update("status", c.PostForm("status"))
  db.Model(&projects).Update("brand", brand)

  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Projects successfully updated!"})
}

func DeleteProject(c *gin.Context){
  var projects models.Projects
  project_id, _ := strconv.Atoi(c.Param("id"))

  db:= database.ConnectionPostgres()
  db.First(&projects, project_id)

  if (projects.ID == 0){
    c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No projects found!"})
    return
  }

  db.Delete(&projects)
  c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Projects successfully deleted!"})
}
