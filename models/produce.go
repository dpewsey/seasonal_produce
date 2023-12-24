// models/produce.go
package models

type Produce struct {
    ID      int    `json:"id" gorm:"primaryKey"`
    Name    string `json:"name"`
    Season  string `json:"season"`
    Group   string `json:"group"` // think citrus 
    Specific_Group string: `json:"specific_group"` // think apple/oranages
    Start_Month  string `json:"start month"`
    End_Month    string `json:"end month"`
    Type    string `json:type` // fruit or veg 
}
