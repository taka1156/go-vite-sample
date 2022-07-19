
-- +migrate Up
CREATE TABLE IF NOT EXISTS `todos` (
    id bigint AUTO_INCREMENT NOT NULL,
    task_name VARCHAR(255),
    task_status VARCHAR(255),
    --  timestamp
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    -- 主キー
    PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS `todos`;
