-- Rollback: Add default_profile column and its GIN index back to tanks table
ALTER TABLE tanks ADD COLUMN IF NOT EXISTS default_profile JSONB;
CREATE INDEX IF NOT EXISTS idx_tanks_default_profile ON tanks USING GIN(default_profile);
