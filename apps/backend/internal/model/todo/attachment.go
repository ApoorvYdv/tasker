package todo

import "github.com/ApoorvYdv/go-tasker/internal/model"

type Attachment struct {
	model.Base
	TodoID      string  `json:"todoId" db:"todo_id"`
	Name        string  `json:"name" db:"name"`
	UploadedBy  string  `json:"uploadedBy" db:"uploaded_by"`
	DownloadKey string  `json:"downloadKey" db:"download_key"`
	FileSize    *int64  `json:"fileSize" db:"file_size"`
	MimeType    *string `json:"mimeType" db:"mime_type"`
}
