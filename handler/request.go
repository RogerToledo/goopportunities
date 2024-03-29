package handler

import (
	"fmt"
	"strings"
)

type UpsertOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpsertOpeningRequest) ValidateCreate() map[string]string {
	var invalidFields = make(map[string]string)

	if r.Role == "" {
		invalidFields["role"] = "string"
	}

	if r.Company == "" {
		invalidFields["company"] = "string"
	}

	if r.Location == "" {
		invalidFields["location"] = "string"
	}

	if r.Remote == nil {
		invalidFields["remote"] = "bool"
	}

	if r.Link == "" {
		invalidFields["link"] = "string"
	}

	if r.Salary <= 0 {
		invalidFields["salary"] = "int64"
	}

	return invalidFields
}

func (r *UpsertOpeningRequest) ValidateUpdate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be filled")
}

func ErrParamRequired(invalidFields map[string]string) string {
	var msg []string
	for k, v := range invalidFields {
		msg = append(msg, fmt.Sprintf("%s is required and should be a %s", k, v))
	}

	oneMsg := strings.Join(msg, " | ")

	return oneMsg
}
