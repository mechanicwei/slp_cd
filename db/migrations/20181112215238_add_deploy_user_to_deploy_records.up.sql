ALTER TABLE deploy_records
ADD deploy_user jsonb DEFAULT '{}'::jsonb;
