SELECT * FROM organizations WHERE slug = $1 AND deleted_at IS NULL;
