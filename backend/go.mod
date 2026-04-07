module backend

go 1.22.0

require (
	github.com/gin-gonic/gin v1.9.1
	gorm.io/driver/postgres v1.5.7
	gorm.io/gorm v1.25.7
)

// In a real environment you would run `go mod tidy` to add the transitive dependencies.
// For this scaffolding, these will be pulled during the Docker build sequence.
