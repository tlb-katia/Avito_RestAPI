// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Tender Management API
 *
 * API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка).
 *
 * API version: 1.0
 */

package openapi

// BidReview - Отзыв о предложении
type BidReview struct {

	// Уникальный идентификатор отзыва, присвоенный сервером.
	Id string `json:"id"`

	// Описание предложения
	Description string `json:"description"`

	// Серверная дата и время в момент, когда пользователь отправил отзыв на предложение. Передается в формате RFC3339.
	CreatedAt string `json:"createdAt"`
}

// AssertBidReviewRequired checks if the required fields are not zero-ed
func AssertBidReviewRequired(obj BidReview) error {
	elements := map[string]interface{}{
		"id":          obj.Id,
		"description": obj.Description,
		"createdAt":   obj.CreatedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertBidReviewConstraints checks if the values respects the defined constraints
func AssertBidReviewConstraints(obj BidReview) error {
	return nil
}
