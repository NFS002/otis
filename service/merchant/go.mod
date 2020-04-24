module gitlab.com/otis_team/backend/service/merchant

go 1.13

replace (
//	gitlab.com/otis_team/backend/api/merchant => ../../api/merchant/ //_LOCAL
//	gitlab.com/otis_team/backend/api/user => ../../api/user //_LOCAL
//	gitlab.com/otis_team/backend/db => ../../db //_LOCAL
//	gitlab.com/otis_team/backend/dtypes => ../../dtypes //_LOCAL
//	gitlab.com/otis_team/backend/service/merchant => ../../service/merchant //_LOCAL
//	gitlab.com/otis_team/backend/service/transaction => ../../service/transaction //_LOCAL
//	gitlab.com/otis_team/backend/service/user => ../../service/user //_LOCAL
)

replace github.com/coreos/etcd v3.3.17+incompatible => github.com/coreos/etcd v3.3.4+incompatible

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro v1.18.0
	github.com/ugorji/go v1.1.7 // indirect
	gitlab.com/otis_team/backend/db v0.0.0-00010101000000-000000000000
	gitlab.com/otis_team/backend/dtypes v0.0.0-00010101000000-000000000000
)
