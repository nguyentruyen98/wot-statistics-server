-- Create modules table for World of Tanks vehicle components
CREATE TABLE IF NOT EXISTS modules (
    id SERIAL PRIMARY KEY,
    module_id INTEGER NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    type VARCHAR(50) NOT NULL, -- vehicleChassis, vehicleTurret, vehicleGun, vehicleEngine, vehicleRadio
    nation VARCHAR(50) NOT NULL,
    tier INTEGER NOT NULL CHECK (tier >= 1 AND tier <= 15), -- Some modules can have tier > 10
    weight INTEGER NOT NULL,
    price_credit INTEGER,
    image TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for fast queries
CREATE INDEX IF NOT EXISTS idx_modules_module_id ON modules(module_id);
CREATE INDEX IF NOT EXISTS idx_modules_type ON modules(type);
CREATE INDEX IF NOT EXISTS idx_modules_nation ON modules(nation);
CREATE INDEX IF NOT EXISTS idx_modules_tier ON modules(tier);
CREATE INDEX IF NOT EXISTS idx_modules_type_nation ON modules(type, nation);
CREATE INDEX IF NOT EXISTS idx_modules_type_tier ON modules(type, tier);

-- Create trigger to auto-update updated_at timestamp
CREATE OR REPLACE FUNCTION update_modules_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_modules_updated_at
    BEFORE UPDATE ON modules
    FOR EACH ROW
    EXECUTE FUNCTION update_modules_updated_at();

-- Add comments for documentation
COMMENT ON TABLE modules IS 'World of Tanks vehicle modules (chassis, turrets, guns, engines, radios)';
COMMENT ON COLUMN modules.module_id IS 'Unique module ID from Wargaming API';
COMMENT ON COLUMN modules.type IS 'Module type: vehicleChassis, vehicleTurret, vehicleGun, vehicleEngine, vehicleRadio';
COMMENT ON COLUMN modules.weight IS 'Module weight in kg';
COMMENT ON COLUMN modules.price_credit IS 'Purchase price in credits';
