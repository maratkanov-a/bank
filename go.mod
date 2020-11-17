module github.com/maratkanov-a/bank

go 1.14

require (
	github.com/go-chi/chi v3.3.4+incompatible
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.1
	github.com/gojuno/minimock/v3 v3.0.8
	github.com/jmoiron/sqlx v1.2.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.8.0
	github.com/maratkanov-a/bank/pkg/accounts v0.0.0-00010101000000-000000000000
	github.com/maratkanov-a/bank/pkg/payments v0.0.0-00010101000000-000000000000
	github.com/mattn/go-sqlite3 v1.10.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/shopspring/decimal v1.2.0
	github.com/sirupsen/logrus v1.4.2
	github.com/stretchr/testify v1.6.1
	github.com/utrack/clay/doc/example v0.0.0-20201026130614-4706f034cfde
	github.com/utrack/clay/v2 v2.4.9
	google.golang.org/grpc v1.33.2
)

replace (
	github.com/maratkanov-a/bank/pkg/accounts => ./pkg/accounts
	github.com/maratkanov-a/bank/pkg/payments => ./pkg/payments
)
