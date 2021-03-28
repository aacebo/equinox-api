CREATE TABLE IF NOT EXISTS organizations (
	id UUID 		PRIMARY KEY,
	slug 				VARCHAR UNIQUE NOT NULL,
	name 				VARCHAR NOT NULL,
	created_at	TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at	TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	deleted_at	TIMESTAMPTZ
);

