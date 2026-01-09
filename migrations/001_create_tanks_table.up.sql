-- Create tanks table
CREATE TABLE IF NOT EXISTS tanks (
    id SERIAL PRIMARY KEY,
    tank_id INTEGER UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    short_name VARCHAR(100) NOT NULL,
    nation VARCHAR(50) NOT NULL,
    tier INTEGER NOT NULL CHECK (tier >= 1 AND tier <= 10),
    type VARCHAR(50) NOT NULL,
    tag VARCHAR(100) NOT NULL,
    
    -- Premium and gift status
    is_premium BOOLEAN DEFAULT FALSE,
    is_premium_igr BOOLEAN DEFAULT FALSE,
    is_gift BOOLEAN DEFAULT FALSE,
    is_wheeled BOOLEAN DEFAULT FALSE,
    
    -- Pricing
    price_credit BIGINT,
    price_gold INTEGER,
    
    -- Description and images
    description TEXT,
    big_icon_url VARCHAR(500),
    
    -- Modules (stored as JSONB for flexibility)
    radios JSONB,
    suspensions JSONB,
    provisions JSONB,
    engines JSONB,
    guns JSONB,
    turrets JSONB,
    crew JSONB,
    
    -- Technical specifications
    modules_tree JSONB,
    next_tanks JSONB,
    prices_xp JSONB,
    
    -- Metadata
    multination VARCHAR(50),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for common queries
CREATE INDEX idx_tanks_tank_id ON tanks(tank_id);
CREATE INDEX idx_tanks_nation ON tanks(nation);
CREATE INDEX idx_tanks_tier ON tanks(tier);
CREATE INDEX idx_tanks_type ON tanks(type);
CREATE INDEX idx_tanks_is_premium ON tanks(is_premium);
CREATE INDEX idx_tanks_name ON tanks(name);
CREATE INDEX idx_tanks_short_name ON tanks(short_name);

-- Composite indexes for common filter combinations
CREATE INDEX idx_tanks_nation_tier ON tanks(nation, tier);
CREATE INDEX idx_tanks_type_tier ON tanks(type, tier);
CREATE INDEX idx_tanks_nation_type_tier ON tanks(nation, type, tier);

-- GIN index for JSONB columns to enable efficient querying

CREATE INDEX idx_tanks_provisions ON tanks USING GIN(provisions);

-- Create function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger to automatically update updated_at
CREATE TRIGGER update_tanks_updated_at
    BEFORE UPDATE ON tanks
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE tanks IS 'Stores World of Tanks vehicle encyclopedia data';
COMMENT ON COLUMN tanks.tank_id IS 'Unique vehicle ID from Wargaming API';
COMMENT ON COLUMN tanks.tier IS 'Vehicle tier (1-10)';
COMMENT ON COLUMN tanks.type IS 'Vehicle type: heavyTank, mediumTank, lightTank, AT-SPG, SPG';
COMMENT ON COLUMN tanks.nation IS 'Nation: ussr, germany, usa, france, uk, china, japan, czech, etc.';
