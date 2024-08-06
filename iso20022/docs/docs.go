package docs

import _ "embed"

//go:embed templates/swagger/swagger.html
var SwaggerTemplate []byte

//go:embed swagger/swagger.json
var SwaggerJson []byte
