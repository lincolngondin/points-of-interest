package config

type configs struct {
	DBDataSourceName string
	DBDriverName     string
}

func New() *configs {
	return &configs{
		DBDriverName:     "postgres",
		DBDataSourceName: "host=localhost port=5432 user=postgres password=lincoln dbname=gps-service sslmode=disable",
	}
}
