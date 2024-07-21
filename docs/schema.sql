DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `username` VARCHAR(50) NOT NULL UNIQUE,
    `email` VARCHAR(100) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `role` ENUM('customer', 'admin') DEFAULT 'customer',
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(100) NOT NULL UNIQUE,
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `category_id` INT,
    `name` VARCHAR(100) NOT NULL,
    `description` TEXT,
    `price` DECIMAL(10, 2) NOT NULL,
    `stock` INT NOT NULL,
    `image_url` VARCHAR(255),
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT,
    `total_price` DECIMAL(10, 2) NOT NULL,
    `status` ENUM('pending', 'paid', 'shipped', 'completed', 'canceled') DEFAULT 'pending',
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT,
    `product_id` INT,
    `quantity` INT NOT NULL,
    `price` DECIMAL(10, 2) NOT NULL,

    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT,
    `product_id` INT,
    `quantity` INT NOT NULL,
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `reviews`;
CREATE TABLE `reviews` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `user_id` INT,
    `product_id` INT,
    `rating` INT CHECK (rating >= 1 AND rating <= 5),
    `comment` TEXT,
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

DROP TABLE IF EXISTS `payments`;
CREATE TABLE `payments` (
    `id` INT AUTO_INCREMENT PRIMARY KEY,
    `order_id` INT,
    `payment_method` VARCHAR(50),
    `payment_status` ENUM('pending', 'completed', 'failed') DEFAULT 'pending',
    `transaction_id` VARCHAR(100),
    
    -- Utility columns
    `created_at` TIMESTAMP(6) NOT NULL,
    `created_by` VARCHAR(50) NOT NULL,
    `updated_at` TIMESTAMP(6) NULL,
    `updated_by` VARCHAR(50) NULL,
    `is_deleted` TINYINT NOT NULL,
    `deleted_at` TIMESTAMP(6) NULL,
    `deleted_by` VARCHAR(50) NULL
);

