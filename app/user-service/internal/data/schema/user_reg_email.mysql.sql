-- srv_user_reg_email 用户注册邮箱
CREATE TABLE srv_user_reg_email
(
    id           BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
    created_time DATETIME(3)             NOT NULL COMMENT '创建时间',
    updated_time DATETIME(3)             NOT NULL COMMENT '最后修改时间',
    deleted_time BIGINT       DEFAULT 0 NULL COMMENT '删除时间',
    user_id      BIGINT       DEFAULT 0  NOT NULL COMMENT 'uid',
    user_email   varchar(255) DEFAULT '' NOT NULL COMMENT '邮箱',
    PRIMARY KEY (id),
    UNIQUE KEY (user_email),
    KEY (user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT = '用户注册邮箱';

