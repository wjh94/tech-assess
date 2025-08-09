INSERT INTO user_roles ("name", "status")
VALUES
('Editor', 1),
('Admin', 1);

-- This migration file is used to seed the user_roles table with initial data.
-- It should be run after creating the user_roles table.
-- Ensure that the user_roles table exists before running this migration.
-- The roles 'Editor' and 'Admin' are added with status 1 (active).