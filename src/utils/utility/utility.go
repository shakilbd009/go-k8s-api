package utility

import "fmt"

var (
	ErrDatabase = fmt.Errorf("database_error")
)

func GetInt32Pointer(i int32) *int32 {
	return &i
}
