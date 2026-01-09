-- Remove module support from profiles table
ALTER TABLE profiles DROP CONSTRAINT IF EXISTS check_entity_type;
ALTER TABLE profiles DROP CONSTRAINT IF EXISTS fk_profiles_module;

DROP INDEX IF EXISTS idx_profiles_module_id_lookup;
DROP INDEX IF EXISTS idx_profiles_module_id;
DROP INDEX IF EXISTS idx_profiles_tank_id;

ALTER TABLE profiles DROP COLUMN IF EXISTS module_id;

-- Restore tank_id as NOT NULL and UNIQUE
ALTER TABLE profiles ALTER COLUMN tank_id SET NOT NULL;
ALTER TABLE profiles ADD CONSTRAINT profiles_tank_id_key UNIQUE (tank_id);
