CREATE TABLE IF NOT EXISTS `category` (
    `id` bigint NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `create_at` bigint NOT NULL,
    `update_at` bigint NOT NULL,
    `name` varchar(255) NOT NULL,
    `key_picture` varchar(1024) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;