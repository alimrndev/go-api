package utils

import (
	"reflect"
)

// UpdateStruct menggunakan refleksi untuk melakukan update pada sebuah struktur data
func UpdateStruct(target interface{}, update interface{}) {
	targetType := reflect.TypeOf(target).Elem() // Mengambil tipe dari nilai yang ditunjuk oleh pointer
	targetValue := reflect.ValueOf(target).Elem()
	updateValue := reflect.ValueOf(update)

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldName := field.Name
		if fieldName == "ID" || fieldName == "CreatedAt" || fieldName == "UpdatedAt" {
			continue // Langsung lanjut jika field adalah ID, CreatedAt, atau UpdatedAt
		}
		updateFieldValue := updateValue.FieldByName(fieldName)
		if updateFieldValue.IsValid() && !updateFieldValue.IsZero() && updateFieldValue.Interface() != "" {
			targetValue.Field(i).Set(updateFieldValue)
		}
	}
}
