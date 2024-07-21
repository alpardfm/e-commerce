package payments

const (
	readPayments = `
	SELECT
		id,
		order_id,
		payment_method,
		payment_status,
		transaction_id,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		payments
	`

	createPayments = `
	INSERT INTO payments (
		order_id,
		payment_method,
		payment_status,
		transaction_id,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:order_id,
		:payment_method,
		:payment_status,
		:transaction_id,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updatePayments = `
	UPDATE
		payments
	SET
		order_id = :order_id,
		payment_method = :payment_method,
		payment_status = :payment_status,
		transaction_id = :transaction_id,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	deletePayments = `
	UPDATE
		payments
	SET
		is_deleted = :is_deleted,
		deleted_at = :deleted_at,
		deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
