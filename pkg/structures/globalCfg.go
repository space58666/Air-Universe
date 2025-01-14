package structures

type BaseConfig struct {
	Log   Log   `json:"log"`
	Panel Panel `json:"panel"`
	Proxy Proxy `json:"proxy"`
	Sync  Sync  `json:"sync"`
}
type Log struct {
	LogLevel string `json:"log_level"`
}
type Panel struct {
	Type    string   `json:"type"`
	URL     string   `json:"url"`
	Key     string   `json:"key"`
	NodeIDs []uint32 `json:"node_ids"`
}
type Proxy struct {
	Type            string    `json:"type"`
	AlertID         uint32    `json:"alert_id"`
	AutoGenerate    bool      `json:"auto_generate"`
	InTags          []string  `json:"in_tags"`
	APIAddress      string    `json:"api_address"`
	APIPort         uint32    `json:"api_port"`
	LogPath         string    `json:"log_path"`
	Cert            Cert      `json:"cert"`
	SpeedLimitLevel []float32 `json:"speed_limit_level"`
}
type Sync struct {
	Interval       uint32 `json:"interval"`
	FailDelay      uint32 `json:"fail_delay"`
	Timeout        uint32 `json:"timeout"`
	PostIPInterval uint32 `json:"post_ip_interval"`
}
type Cert struct {
	CertPath string `json:"cert_path"`
	KeyPath  string `json:"key_path"`
}
