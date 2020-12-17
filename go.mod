module github.com/shodhangk/go-auth

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.1
)

replace (
	github.com/shodhangk/go-auth/controllers => /Volumes/Development/Development/workspace/go/go-auth/controllers
	github.com/shodhangk/go-auth/models => /Volumes/Development/Development/workspace/go/go-auth/models
	github.com/shodhangk/go-auth/routes => /Volumes/Development/Development/workspace/go/go-auth/routes
	github.com/shodhangk/go-auth/utils/auth => /Volumes/Development/Development/workspace/go/go-auth/utils/auth
)
