// Package client provides the client-side functionalities for interacting with the Rabbtix API.
package client

// Importing necessary libraries.
import (
	"rabbitx-client/model"
	"reflect"
	"time"
)

// TILL_EXPIRATION is the duration before the API secret's expiration when it is considered close to expiring.
const TILL_EXPIRATION = time.Hour * 2

// isApiSecretExpired checks if the provided API secret is expired.
// It returns true if the API secret is nil or its expiration time is less than or equal to the current time.
// Otherwise, it returns false.
func isApiSecretExpired(apiSecret *model.APISecret) bool {
	if apiSecret == nil {
		return true
	}

	if int64(apiSecret.Expiration) <= time.Now().Unix() {
		return true
	}

	return false
}

// isCloseToExpired checks if the provided API secret is close to expiring.
// It returns true if the API secret is nil or its expiration time is less than or equal to the current time plus TILL_EXPIRATION.
// Otherwise, it returns false.
func isCloseToExpired(apiSecret *model.APISecret) bool {
	if apiSecret == nil {
		return true
	}

	if time.Now().Add(TILL_EXPIRATION).Unix() >= int64(apiSecret.Expiration) {
		return true
	}

	return false
}

// makeQueryParams converts the fields of the provided itemValue into a map of query parameters.
// It only considers fields of type string.
func makeQueryParams(itemValue reflect.Value) map[string]string {
	queryParams := make(map[string]string)

	for i := 0; i < itemValue.NumField(); i++ {
		if itemValue.Field(i).Kind() == reflect.String {
			queryParams[itemValue.Type().Field(i).Name] = itemValue.Field(i).String()
		}
	}

	return queryParams
}
