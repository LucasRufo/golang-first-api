package handler

import "fmt"

func errorParamIsRequired(paramName, paramType string) error {
	return fmt.Errorf("param: %s (type %s) is required", paramName, paramType)
}

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errorParamIsRequired("role", "string")
	}

	if r.Company == "" {
		return errorParamIsRequired("company", "string")
	}

	if r.Location == "" {
		return errorParamIsRequired("location", "string")
	}

	if r.Link == "" {
		return errorParamIsRequired("link", "string")
	}

	if r.Remote == nil {
		return errorParamIsRequired("remote", "bool")
	}

	if r.Salary <= 0 {
		return errorParamIsRequired("salary", "int64")
	}

	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}

	return fmt.Errorf("at least one valid field must be provided")
}
