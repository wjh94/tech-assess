TRUNCATE user_roles RESTART IDENTITY;

-- This migration file is used to remove all entries from the user_roles table.
-- It should be run when rolling back the migration that seeded the user_roles table.
-- Ensure that the user_roles table exists before running this migration.
-- This will reset the table, allowing for a fresh start if needed.