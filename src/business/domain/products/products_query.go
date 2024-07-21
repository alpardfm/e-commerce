package products

const (
	readProducts = `
	SELECT
		id,
		category_id,
		name,
		description,
		discount_price,
		price,
		stock,
		image_url,
		created_at,
	    created_by,
	    COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		products`

	createProducts = `
	INSERT INTO products (
		category_id,
		name,
		description,
		discount_price,
		price,
		stock,
		image_url,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:category_id,
		:name,
		:description,
		:discount_price,
		:price,
		:stock,
		:image_url,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateProducts = `
	UPDATE
		products
	SET
		category_id = :category_id,
		name = :name,
		description = :description,
		discount_price = :discount_price,
		price = :price,
		stock = :stock,
		image_url = :image_url,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id`

	deleteProducts = `
	UPDATE
		products
	SET
		is_deleted = :is_deleted,
		deleted_at = :deleted_at,
		deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
