package strings

import "fmt"

var (
	// Download Successful
	DownloadSuccessful = "Successfully downloaded template"

	// Download Errors
	ErrDownloadFailedGeneric = "Failed to download template"

	ErrTemplateAlreadyExists = func(templateName string) string {
		return fmt.Sprintf("%s: Template (%s) already exists. Use -f to overwrite", ErrDownloadFailedGeneric, templateName)
	}
)
