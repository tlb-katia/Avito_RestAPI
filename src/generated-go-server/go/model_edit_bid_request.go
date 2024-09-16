// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Tender Management API
 *
 * API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка).
 *
 * API version: 1.0
 */

package openapi

type EditBidRequest struct {

	// Полное название предложения
	Name string `json:"name,omitempty"`

	// Описание предложения
	Description string `json:"description,omitempty"`
}

// AssertEditBidRequestRequired checks if the required fields are not zero-ed
func AssertEditBidRequestRequired(obj EditBidRequest) error {
	return nil
}

// AssertEditBidRequestConstraints checks if the values respects the defined constraints
func AssertEditBidRequestConstraints(obj EditBidRequest) error {
	return nil
}
