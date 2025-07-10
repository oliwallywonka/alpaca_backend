package shared

import (
	"database/sql/driver"
	"encoding/json"
)

type LocationField struct {
	Lat float64
	Lng float64
}

func (m *LocationField) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, m)
}

func (m LocationField) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type PriceField []struct {
	MinPersons  int
	MaxPersons  int
	Price       float64
	Currency    string
	IsPerPerson bool
}

func (m *PriceField) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, m)
}

func (m PriceField) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type ContactField []struct {
	Type  string
	Value string
}

func (m *ContactField) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, m)
}

func (m ContactField) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type PermissionField []string

func (m *PermissionField) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, m)
}

func (m PermissionField) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type ImageField []string

func (m *ImageField) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, m)
}

func (m ImageField) Value() (driver.Value, error) {
	return json.Marshal(m)
}

type LangField map[string]string

func (m *LangField) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, m)
}

func (m LangField) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func HasLangField(tableName string, columnName string) bool {
	langFieldColumns := map[string][]string{
		"destination": {"name", "description"},
		"tag":         {"title", "description"},
		"resource":    {"name", "description"},
		"tour": {
			"name", "slug", "transport", "accommodation", "team",
			"short_description", "long_description",
		},
		"tour_destination": {"description"},
	}

	return searchMap(langFieldColumns, tableName, columnName)
}

func HasImageField(tableName string, columnName string) bool {
	imageFieldColumns := map[string][]string{
		"tour":                {"images"},
		"resource":            {"images"},
		"section_description": {"images"},
	}
	return searchMap(imageFieldColumns, tableName, columnName)
}

func HasPermissionField(tableName string, columnName string) bool {
	imageFieldColumns := map[string][]string{
		"role": {"permissions"},
	}
	return searchMap(imageFieldColumns, tableName, columnName)
}

func HasContactField(tableName string, columnName string) bool {
	contactFieldColumns := map[string][]string{
		"customer": {"contacts"},
		"provider": {"contacts"},
		"user":     {"contacts"},
	}
	return searchMap(contactFieldColumns, tableName, columnName)
}

func HasPriceField(tableName string, columnName string) bool {
	priceFieldColumns := map[string][]string{
		"resource_provider": {"ref_prices"},
	}
	return searchMap(priceFieldColumns, tableName, columnName)
}

func HasLocationField(tableName string, columnName string) bool {
	locationFieldColumns := map[string][]string{
		"destination": {"location"},
		"resource":    {"location"},
	}
	return searchMap(locationFieldColumns, tableName, columnName)
}

func searchMap(m map[string][]string, tableName, columnName string) bool {
	_, ok := m[tableName]
	if !ok {
		return false
	}
	for _, column := range m[tableName] {
		if column == columnName {
			return true
		}
	}
	return false
}
