SELECT
	id,
	key,
	slug,
	name,
	created_at,
	updated_at
FROM organizations
WHERE id = $1
AND deleted_at IS NULL;
