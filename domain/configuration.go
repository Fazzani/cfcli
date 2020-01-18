package domain

type Configuration struct {
	Authentication Authentication `json:"authentication"`
	API            API            `json:"api"`
	Log            Log            `json:"log"`
}

type API struct {
	BaseURL string `json:"base_url"`
}

type Authentication struct {
	AccountID string `json:"account_id"`
	Email     string `json:"email"`
	Key       string `json:"key"`
}

type Log struct {
	FilePath string `json:"file_path"`
	Level    string `json:"level"`
}
