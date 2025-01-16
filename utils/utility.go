package utils

import (
	"mime"
	"path/filepath"
	"strings"
)

func GetMediaType(filePath string) string {
	// Extract the file extension
	ext := filepath.Ext(filePath)

	// Get the MIME type based on the file extension
	mimeType := mime.TypeByExtension(ext)

	// Define media types based on MIME type
	switch mimeType {
	case "image/jpeg", "image/png", "image/gif", "image/bmp", "image/webp":
		return "Image"
	case "application/pdf":
		return "PDF"
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "application/vnd.ms-excel":
		return "Excel"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/msword":
		return "Word"
	default:
		return "Unknown Media Type"
	}
}

func GetMimeType(extension string) string {
	// Convert the extension to lowercase and ensure it starts with a dot
	extension = strings.ToLower(extension)
	if !strings.HasPrefix(extension, ".") {
		extension = "." + extension
	}

	// Map file extensions to their MIME types
	mimeTypes := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".bmp":  "image/bmp",
		".webp": "image/webp",
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}

	// Look up the MIME type in the map
	mimeType, exists := mimeTypes[extension]
	if exists {
		return mimeType
	}

	return "application/octet-stream" // Default MIME type for unknown extensions
}
