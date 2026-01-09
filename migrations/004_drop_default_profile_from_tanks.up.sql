-- Drop default_profile column and its GIN index from tanks table
DROP INDEX IF EXISTS idx_tanks_default_profile;
ALTER TABLE tanks DROP COLUMN IF EXISTS default_profile;
