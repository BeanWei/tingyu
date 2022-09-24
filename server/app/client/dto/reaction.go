//lint:file-ignore SA5008 .
package dto

type Reaction struct {
	SubjectId int64  `json:"-" db:"subject_id"`
	Code      string `json:"code" db:"code"`
	Count     int    `json:"count" db:"count"`
	Active    bool   `json:"active" db:"active"`
}
