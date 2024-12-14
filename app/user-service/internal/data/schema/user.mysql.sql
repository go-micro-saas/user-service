-- srv_user 用户
CREATE TABLE srv_user
(
    id             BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
    created_time   DATETIME(3)             NOT NULL COMMENT '创建时间',
    updated_time   DATETIME(3)             NOT NULL COMMENT '最后修改时间',
    deleted_time   BIGINT       DEFAULT 0 NULL COMMENT '删除时间',
    user_id        BIGINT       DEFAULT 0  NOT NULL COMMENT 'uid',
    user_phone     varchar(255) DEFAULT '' NOT NULL COMMENT '手机',
    user_email     varchar(255) DEFAULT '' NOT NULL COMMENT '邮箱',
    user_nickname  varchar(255) DEFAULT '' NOT NULL COMMENT '昵称',
    user_avatar    varchar(255) DEFAULT '' NOT NULL COMMENT '头像',
    user_gender    varchar(255) DEFAULT '' NOT NULL COMMENT '性别',
    register_type  varchar(255) DEFAULT '' NOT NULL COMMENT '注册类型',
    user_status    integer      DEFAULT 0  NOT NULL COMMENT '状态',
    disable_time   DATETIME(3)             NULL COMMENT '禁用时间',
    blacklist_time DATETIME(3)             NULL COMMENT '黑名单时间',
    password_hash  varchar(255) DEFAULT '' NOT NULL COMMENT '密码',
    PRIMARY KEY (id),
    UNIQUE KEY (user_id),
    KEY (deleted_time)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '用户表';

