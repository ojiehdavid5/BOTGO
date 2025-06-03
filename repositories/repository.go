package repositories




import (
 "github.com/chuks/BOTGO/database"

 "gorm.io/gorm"
)

var DB *gorm.DB = database.Init()