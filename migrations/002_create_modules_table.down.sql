-- Drop modules table and related objects
DROP TRIGGER IF EXISTS trigger_update_modules_updated_at ON modules;
DROP FUNCTION IF EXISTS update_modules_updated_at();
DROP TABLE IF EXISTS modules CASCADE;
