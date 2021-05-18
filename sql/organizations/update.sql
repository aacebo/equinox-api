UPDATE organizations SET
key = $2, slug = $3, name = $4, created_at = $5, updated_at = $6
WHERE id = $1;
