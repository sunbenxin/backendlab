--
CREATE DATABASE IF NOT EXISTS backend DEFAULT CHARSET utf8 COLLATE utf8_general_ci;

-- TABLE: order
DROP TABLE IF EXISTS `orders`;
-- pg
CREATE TABLE order (
    id              BIGSERIAL     PRIMARY KEY,
    description     TEXT,

    created_by      BIGINT        NOT NULL,
    created_at      TIMESTAMPTZ   NOT NULL,
    updated_by      BIGINT,
    updated_at      TIMESTAMPTZ,
    deleted_by      BIGINT,
    deleted_at      TIMESTAMPTZ
);

-- mysql
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`
(
    `id`         bigint       NOT NULL AUTO_INCREMENT,
    `description`   varchar(256) NOT NULL,

    `created_at` timestamp       NOT NULL COMMENT '创建时间',
    `created_by` bigint       NOT NULL COMMENT '',
    `updated_at` timestamp       NOT NULL COMMENT '更新时间',
    `updated_by` BIGINT,
    `deleted_at` timestamp       NULL COMMENT '',
    `deleted_by` BIGINT  NULL,
    PRIMARY KEY (`id`) USING BTREE
);
