module github.com/jonhorner/swgoh-api

go 1.22.2

replace swgoh-api/units => ./units

replace swgoh-api/guild => ./guild

replace swgoh-api/db => ./db

require github.com/joho/godotenv v1.5.1

require (
	swgoh-api/guild v0.0.0-00010101000000-000000000000
	swgoh-api/units v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-lambda-go v1.47.0 // indirect
	github.com/aws/aws-sdk-go v1.52.2 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	swgoh-api/db v0.0.0-00010101000000-000000000000 // indirect
)
