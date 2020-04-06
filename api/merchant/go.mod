module gitlab.com/otis-team/backend/api/merchant

go 1.13

require (
	github.com/golang/protobuf v1.3.5
	github.com/micro/examples v0.2.0
	github.com/micro/go-micro v1.18.0
	gitlab.com/otis-team/backend/service/merchant v0.0.0-00010101000000-000000000000
	gitlab.com/otis-team/backend/service/transaction v0.0.0-00010101000000-000000000000
)

replace (
	gitlab.com/otis-team/backend/api/merchant => /Usets/noah/backend/api/user
	gitlab.com/otis-team/backend/api/user => /Users/noah/backend/api/user
	gitlab.com/otis-team/backend/db => /Users/noah/backend/db
	gitlab.com/otis-team/backend/service/merchant => /Users/noah/backend/service/merchant
	gitlab.com/otis-team/backend/service/transaction => /Users/noah/backend/service/transaction
	gitlab.com/otis-team/backend/service/user => /Users/noah/backend/service/user
)
