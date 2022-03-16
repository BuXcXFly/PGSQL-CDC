package main

// 如果使用JSON解码插件，这里是用于Decode的Schema
type Payload struct {
	Change []struct {
		Kind         string        `json:"kind"`
		Schema       string        `json:"schema"`
		Table        string        `json:"table"`
		ColumnNames  []string      `json:"columnnames"`
		ColumnTypes  []string      `json:"columntypes"`
		ColumnValues []interface{} `json:"columnvalues"`
		OldKeys      struct {
			KeyNames  []string      `json:"keynames"`
			KeyTypes  []string      `json:"keytypes"`
			KeyValues []interface{} `json:"keyvalues"`
		} `json:"oldkeys"`
	} `json:"change"`
}
