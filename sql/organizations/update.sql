UPDATE organizations SET
slug = $2, name = $3, created_at = $4, updated_at = $5
WHERE id = $1;
