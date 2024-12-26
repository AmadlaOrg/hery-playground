package template

// NewTemplateService to set up the model service
func NewTemplateService(templatePath string) ITemplate {
	return &STemplate{
		templatePath: templatePath,
	}
}
