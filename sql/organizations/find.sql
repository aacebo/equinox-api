SELECT
	id,
	key,
	slug,
	name,
	created_at,
	updated_at
FROM organizations
WHERE deleted_at IS NULL
AND name ILIKE $1
OFFSET $2
LIMIT $3;
