SELECT
	id,
	slug,
	name,
	created_at,
	updated_at
FROM organizations
WHERE slug = $1
AND deleted_at IS NULL;
