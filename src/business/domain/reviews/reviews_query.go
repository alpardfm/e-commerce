package reviews

const (
	readReviews = `
	SELECT
		id,
		user_id,
		product_id,
		rating,
		comment,
		created_at,
		created_by,
		COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
		COALESCE(updated_by, "") as updated_by,
		COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
		COALESCE(deleted_by, "") as deleted_by,
		is_deleted
	FROM
		reviews
	`

	createReviews = `
	INSERT INTO reviews (
		user_id,
		product_id,
		rating,
		comment,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:user_id,
		:product_id,
		:rating,
		:comment,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateReviews = `
	UPDATE
		reviews
	SET
		user_id = :user_id,
		product_id = :product_id,
		rating = :rating,
		comment = :comment,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	deleteReviews = `
	UPDATE
		reviews
	SET
		is_deleted = :is_deleted,
		deleted_at = :deleted_at,
		deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
