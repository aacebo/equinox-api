SELECT
	id,
	slug,
	name,
	created_at,
	updated_at
FROM organizations
WHERE deleted_at IS NULL;
