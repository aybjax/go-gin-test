package data_structures

// InsertData is struct with data
//		of authorized users
type InsertData struct {
	User    int    `json:"user_id" binding:"required_with=device_id"`
	Device  string `json:"device_id" binding:"required"`
	Token   string `json:"token" binding:"required"`
	Os      string `json:"os" binding:"required"`
	Version string `json:"version" binding:"required"`
}
