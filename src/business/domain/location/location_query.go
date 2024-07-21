package location

const (
	createLocation = `
	INSERT INTO location (
		lat,
		long,
		distance,
		secret,
		created_at,
		created_by,
		is_deleted
	)
	VALUES (
		:lat,
		:long,
		:distance,
		:secret,
		:created_at,
		:created_by,
		:is_deleted
	)`

	updateLocation = `
	UPDATE
		location
	SET
		lat = :lat,
		long = :long,
		distance = :distance,
		secret = :secret,
		updated_at = :updated_at,
		updated_by = :updated_by,
		is_deleted = :is_deleted
	WHERE
		id = :id
	`

	readLocation = `
	SELECT
		id,
		lat,
		long,
		distance,
		secret,
		COALESCE(updated_at, TIMESTAMP("01-01-0001")) as updated_at,
	    COALESCE(updated_by, "") as updated_by,
	    COALESCE(deleted_at, TIMESTAMP("01-01-0001")) as deleted_at,
	    COALESCE(deleted_by, "") as deleted_by,
	    is_deleted
	FROM
		location
	`

	deleteLocation = `
	UPDATE
		location
	SET
		is_deleted = :is_deleted,
	   	deleted_at = :deleted_at,
	   	deleted_by = :deleted_by
	WHERE
		id = :id
	`
)
