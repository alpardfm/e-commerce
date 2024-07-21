package orders

const (
	readOrders = `
	SELECT
		id,
		user_id,
		total_price,
		status,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		orders`

	createOrders = `
	INSERT INTO orders (
		user_id,
		total_price,
		status,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:user_id,
		:total_prices,
		:status,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateOrders = `
	UPDATE
		orders
	SET
		user_id = :user_id,
		total_price = :total_price,
		status = :status,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	deleteOrders = `
	UPDATE
		orders
	SET
		is_deleted = :is_deleted,
		deleted_at = :deleted_at,
		deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
