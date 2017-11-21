package models

import (
  "github.com/jinzhu/gorm"
_ "github.com/jinzhu/gorm/dialects/postgres"
  // "time"
)

/*
*   Users Models
*/

type Users struct {
	gorm.Model
	Username     string  `json:"username" gorm:"type:varchar(255);unique_index"`
	Password     string  `json:"password"`
  Fullname     string  `json:"fullname" gorm:"type:varchar(255);index"`
  Email        string  `json:"email" gorm:"unique_index"`
  Phone        int     `json:"phone" gorm:"size:11"`
  Description  string  `json:"description"`
  Status       string  `json:"status" sql:"size:1;index"`
  Type         string  `json:"type" gorm:"type:varchar(255);index"`
  Uid          int     `json:"uid" gorm:"index"`
}

type TransformedUsers struct {
	ID           uint
  Username     string
	Fullname     string
  Email        string
  Phone        int
  Description  string
  Status       string
}

/* End Of users models part*/

/*
*   Brands Models
*/

type Brands struct {
  gorm.Model
  Name        string    `json:"name" gorm:"type:varchar(255);index"`
  Description string    `json:"description"`
  Status      string    `json:"status" sql:"size:1;index"`
  Client      int       `json:"client" gorm:"index"`
}

type TransformedBrands struct{
  ID          uint
  Name        string
  Description string
  Status      string
  Client      int
}

/* End Of brands parts */

/*
*   Projects Models
*/

type Projects struct {
  gorm.Model
  Code        string    `json:"code" gorm:"type:varchar(255);unique_index"`
  Name        string    `json:"name" gorm:"type:varchar(255);index"`
  Description string    `json:"description"`
  Start       string `json:"start" gorm:"index"`
  End         string `json:"start" gorm:"index"`
  Status      string    `json:"status" sql:"size:1;index"`
  Brand       int       `json:"brand" gorm:"index"`
}

type TransformedProjects struct {
  gorm.Model
  ID          uint
  Code        string
  Name        string
  Description string
  Status      string
  Start       string
  End         string
  Brand       int
}
/* End Of Projects */
