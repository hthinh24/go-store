module github.com/hthinh24/go-store/services/product

go 1.24.3

replace github.com/hthinh24/go-store/internal/pkg => ../../pkg

require (
	github.com/gin-gonic/gin v1.10.1
	github.com/hthinh24/go-store/internal/pkg v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.30.1
)
