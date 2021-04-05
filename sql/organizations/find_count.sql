SELECT
	COUNT(id)
FROM organizations
WHERE deleted_at IS NULL
AND name ILIKE $1;
