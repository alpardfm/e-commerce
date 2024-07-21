package cart

const (
	readCart = `
	SELECT
		id,
		user_id,
		product_id,
		quantity,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		cart`

	createCart = `
	INSERT INTO cart (
		user_id,
		product_id,
		quantity,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:user_id,
		:product_id,
		:quantity,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateCart = `
	UPDATE
		cart
	SET
		user_id = :user_id,
		product_id = :product_id,
		quantity = :quantity,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	deleteCart = `
	UPDATE
		cart
	SET
		is_deleted = :is_deleted,
		deleted_at = :deleted_at,
		deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
