package storage

import "os"

func NuevoStorage() *PostgresStorage {
	dsn := ConstruirDSN(
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "archibase"),
		getEnv("DB_PASSWORD", "archibase"),
		getEnv("DB_NAME", "archibase"),
		getEnv("DB_SSLMODE", "disable"),
	)
	return NuevoPostgresStorage(dsn)
}

func getEnv(clave, valorPorDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return valorPorDefecto
}
