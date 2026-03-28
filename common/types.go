package common

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type PaginationRequest struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"pagesize,omitempty"`
}

type Pagination struct {
	Page      int `json:"page,omitempty"`
	PageSize  int `json:"pagesize,omitempty"`
	Total     int `json:"total,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
}

type IDResponse struct {
	ID string `json:"id"`
}

type OperationResult struct {
	Success    []string          `json:"success,omitempty"`
	Fail       []string          `json:"fail,omitempty"`
	FailReason map[string]string `json:"fail_reason,omitempty"`
	Message    string            `json:"message,omitempty"`
}

type IDsOperationResponse struct {
	Success []string `json:"success"`
	Fail    []string `json:"fail"`
}

type Text string

func (t *Text) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*t = ""
		return nil
	}

	var value any
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	switch v := value.(type) {
	case string:
		*t = Text(v)
	case float64:
		*t = Text(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		*t = Text(strconv.FormatBool(v))
	default:
		return fmt.Errorf("unsupported text value type %T", value)
	}

	return nil
}
