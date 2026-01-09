-- Add module_id column to profiles table
ALTER TABLE profiles ADD COLUMN module_id INTEGER UNIQUE;

-- Make tank_id nullable (was UNIQUE NOT NULL before)
ALTER TABLE profiles ALTER COLUMN tank_id DROP NOT NULL;

-- Drop the existing UNIQUE constraint on tank_id
ALTER TABLE profiles DROP CONSTRAINT IF EXISTS profiles_tank_id_key;

-- Add unique constraints for tank_id and module_id separately
CREATE UNIQUE INDEX idx_profiles_tank_id ON profiles(tank_id) WHERE tank_id IS NOT NULL;
CREATE UNIQUE INDEX idx_profiles_module_id ON profiles(module_id) WHERE module_id IS NOT NULL;

-- Add CHECK constraint to ensure either tank_id or module_id is provided (not both, not none)
ALTER TABLE profiles ADD CONSTRAINT check_entity_type 
    CHECK ((tank_id IS NOT NULL AND module_id IS NULL) OR (tank_id IS NULL AND module_id IS NOT NULL));

-- Add foreign key constraint for module_id
ALTER TABLE profiles ADD CONSTRAINT fk_profiles_module 
    FOREIGN KEY (module_id) REFERENCES modules(module_id) ON DELETE CASCADE;

-- Add index for module_id
CREATE INDEX idx_profiles_module_id_lookup ON profiles(module_id);

-- Add comment
COMMENT ON COLUMN profiles.module_id IS 'Module ID - either tank_id or module_id must be set, but not both';
