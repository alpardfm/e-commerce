package order_items

const (
	createOrderItems = `
	INSERT INTO order_items (
		order_id,
		product_id,
		quantity,
		price,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:order_id,
		:product_id,
		:quantity,
		:price,
		:created_at,
		:created_by,
		:is_deleted
	)`

	readOrderItems = `
	SELECT
		id,
		order_id,
		product_id,
		quantity,
		price,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		order_items
	`

	updateOrderItems = `
	UPDATE
		order_items
	SET
		order_id = :order_id,
		product_id = :product_id,
		quantity = :quantity,
		price = :price,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	deleteOrderItems = `
	UPDATE
		order_items
	SET
		is_deleted = :is_deleted,
		deleted_at = :deleted_at,
		deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
