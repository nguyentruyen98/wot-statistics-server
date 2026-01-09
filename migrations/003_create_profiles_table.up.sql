-- Create profiles table
CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    tank_id INTEGER UNIQUE NOT NULL REFERENCES tanks(tank_id) ON DELETE CASCADE,
    
    -- Basic characteristics
    hp INTEGER,
    hull_hp INTEGER,
    hull_weight INTEGER,
    max_ammo INTEGER,
    max_weight INTEGER,
    speed_backward DECIMAL(10, 2),
    speed_forward DECIMAL(10, 2),
    weight INTEGER,
    
    -- Ammo characteristics (stored as JSONB array)
    ammo JSONB,
    
    -- Armor characteristics
    armor_hull_front INTEGER,
    armor_hull_rear INTEGER,
    armor_hull_sides INTEGER,
    armor_turret_front INTEGER,
    armor_turret_rear INTEGER,
    armor_turret_sides INTEGER,
    
    -- Engine characteristics
    engine_fire_chance DECIMAL(5, 4),
    engine_name VARCHAR(255),
    engine_power INTEGER,
    engine_tag VARCHAR(255),
    engine_tier INTEGER,
    engine_weight INTEGER,
    
    -- Gun characteristics
    gun_aim_time DECIMAL(10, 2),
    gun_caliber INTEGER,
    gun_dispersion DECIMAL(10, 4),
    gun_fire_rate DECIMAL(10, 2),
    gun_move_down_arc INTEGER,
    gun_move_up_arc INTEGER,
    gun_name VARCHAR(255),
    gun_reload_time DECIMAL(10, 2),
    gun_tag VARCHAR(255),
    gun_tier INTEGER,
    gun_traverse_speed INTEGER,
    gun_weight INTEGER,
    
    -- Mounted modules
    modules_engine_id INTEGER,
    modules_gun_id INTEGER,
    modules_radio_id INTEGER,
    modules_suspension_id INTEGER,
    modules_turret_id INTEGER,
    
    -- Radio characteristics
    radio_name VARCHAR(255),
    radio_signal_range INTEGER,
    radio_tag VARCHAR(255),
    radio_tier INTEGER,
    radio_weight INTEGER,
    
    -- Rapid mode (for wheeled vehicles) stored as JSONB
    rapid JSONB,
    
    -- Siege mode characteristics stored as JSONB
    siege JSONB,
    
    -- Suspension characteristics
    suspension_load_limit INTEGER,
    suspension_name VARCHAR(255),
    suspension_steering_lock_angle INTEGER,
    suspension_tag VARCHAR(255),
    suspension_tier INTEGER,
    suspension_traverse_speed INTEGER,
    suspension_weight INTEGER,
    
    -- Turret characteristics
    turret_hp INTEGER,
    turret_name VARCHAR(255),
    turret_tag VARCHAR(255),
    turret_tier INTEGER,
    turret_traverse_left_arc INTEGER,
    turret_traverse_right_arc INTEGER,
    turret_traverse_speed INTEGER,
    turret_view_range INTEGER,
    turret_weight INTEGER,
    
    -- Metadata
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for common queries
CREATE INDEX idx_profiles_tank_id ON profiles(tank_id);
CREATE INDEX idx_profiles_tier ON profiles(engine_tier);
