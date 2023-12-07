generate-swagger:
	swag init

# Define the paths for the input Swagger file and the output OpenAPI file
SWAGGER_FILE = ./docs/swagger.yaml
OPENAPI_FILE = ./docs/openapi.yaml

# Makefile target to convert Swagger 2.0 to OpenAPI 3.0+
convert-openapi:
	npm install -g swagger2openapi
	swagger2openapi $(SWAGGER_FILE) -o $(OPENAPI_FILE)
	@echo "Conversion complete. OpenAPI 3.0+ spec generated at $(OPENAPI_FILE)"

# Default target
all: generate-swagger convert-openapi
