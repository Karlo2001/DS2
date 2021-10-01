/*
 * Course API
 *
 * REST API exposing data about courses at ITU
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Course struct {

	Id int32 `json:"id"`

	Name string `json:"name"`

	SatisfactionRating float32 `json:"satisfaction_rating"`
}
