CREATE TABLE IF NOT EXISTS `todos` (
    `id` VARCHAR(36) NOT NULL COMMENT 'TODO ID'
    , `content` VARCHAR(1000) NOT NULL COMMENT 'TODO 内容'
    , `created_at` DATETIME(6) NOT NULL COMMENT 'TODO 作成日時'
    , `updated_at` DATETIME(6) NOT NULL COMMENT 'TODO 更新日時'
    , `completed_at` DATETIME(6) NULL COMMENT 'TODO 完了日時'
, PRIMARY KEY (`id`)
)