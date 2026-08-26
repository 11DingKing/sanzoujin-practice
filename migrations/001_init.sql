-- Schema is embedded in internal/storage/migrate.go so startup can run atomically.
-- This versioned file documents the same migration for operators and review.
CREATE TABLE IF NOT EXISTS schema_migrations(version INTEGER PRIMARY KEY, applied_at TEXT NOT NULL);
