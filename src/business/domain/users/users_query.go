package users

const (
	createUsers = `
	INSERT INTO users (
		username,
		email,
		password,
		pincode,
		role_id,
		is_active,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:username,
		:email,
		:password,
		:pincode,
		:role_id,
		:is_active,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateUsers = `
	UPDATE
		users
	SET
		username = :username,
		email = :email,
		password = :password,
		pincode = :pincode,
		role_id = :role_id,
		is_active = :is_active,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id
	`

	readUsers = `
	SELECT
		id,
		username,
		email,
		password,
		pincode,
		role_id,
		is_active,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		users
	`

	deleteUsers = `
	UPDATE
		users
	SET
		is_deleted = :is_deleted,
	   	deleted_at = :deleted_at,
	   	deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
