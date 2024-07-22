package refund

const (
	createRefund = `
	INSERT INTO refund (
		user_id,
		order_id,
		reason,
		status,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:user_id,
		:order_id,
		:reason,
		:status,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateRefund = `
	UPDATE
		refund
	SET
		user_id = :user_id,
		order_id = :order_id,
		reason = :reason,
		status = :status,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	readRefund = `
	SELECT
		id,
		user_id,
		order_id,
		reason,
		status,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		refund`

	deleteRefund = `
	UPDATE
		refund
	SET
		is_deleted = :is_deleted,
	   	deleted_at = :deleted_at,
	   	deleted_by = :deleted_by
	WHERE
		id = :id`
)
