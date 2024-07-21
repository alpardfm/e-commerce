package role

const (
	createRole = `
	INSERT INTO role (
		name,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:name,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateRole = `
	UPDATE
		role
	SET
		name = :name,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id
	`

	readRole = `
	SELECT
		id,
		name,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		role`

	deleteRole = `
	UPDATE
		role
	SET
		is_deleted = :is_deleted,
	   	deleted_at = :deleted_at,
	   	deleted_by = :deleted_by
	WHERE
		id = :id`
)
