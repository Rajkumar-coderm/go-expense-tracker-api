package models

type CSVImportResponse struct {
	TotalRows    int           `json:"total_rows"`
	SuccessCount int           `json:"success_count"`
	FailedCount  int           `json:"failed_count"`
	FailedRows   []CSVRowError `json:"failed_rows,omitempty"`
}

type CSVRowError struct {
	RowNumber int    `json:"row_number"`
	Error     string `json:"error"`
}
