-- Drop trigger
DROP TRIGGER IF EXISTS update_tanks_updated_at ON tanks;

-- Drop function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop indexes
DROP INDEX IF EXISTS idx_tanks_provisions;
DROP INDEX IF EXISTS idx_tanks_default_profile;
DROP INDEX IF EXISTS idx_tanks_nation_type_tier;
DROP INDEX IF EXISTS idx_tanks_type_tier;
DROP INDEX IF EXISTS idx_tanks_nation_tier;
DROP INDEX IF EXISTS idx_tanks_short_name;
DROP INDEX IF EXISTS idx_tanks_name;
DROP INDEX IF EXISTS idx_tanks_is_premium;
DROP INDEX IF EXISTS idx_tanks_type;
DROP INDEX IF EXISTS idx_tanks_tier;
DROP INDEX IF EXISTS idx_tanks_nation;
DROP INDEX IF EXISTS idx_tanks_tank_id;

-- Drop table
DROP TABLE IF EXISTS tanks;
