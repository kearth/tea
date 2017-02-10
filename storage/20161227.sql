DROP TABLE IF EXISTS User;
CREATE TABLE IF NOT EXISTS User(
    id         bigint(20)      NOT NULL AUTO_INCREMENT,
    account    varchar(64)     NOT NULL DEFAULT '',
    password   varchar(64)     NOT NULL DEFAULT '',
    status     tinyint(3)      NOT NULL DEFAULT 0,
    nickname   varchar(64)     NOT NULL DEFAULT '',
    telphone   varchar(16)     NOT NULL DEFAULT '',
    email      varchar(128)    NOT NULL DEFAULT '',
    headimg    varchar(256)    NOT NULL DEFAULT '',
    tokens     varchar(512)    NOT NULL DEFAULT '',
    createtime timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    updatetime timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    exist      tinyint(3)      NOT NULL DEFAULT 0,
    PRIMARY KEY (id) 
);
