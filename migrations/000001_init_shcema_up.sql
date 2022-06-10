CREATE TABLE IF NOT EXISTS `grants`
(
    `oauth_uuid`      VARCHAR(50)  NOT NULL PRIMARY KEY,
    `acces_token`     VARCHAR(100) NOT NULL,
    `expires_date`    DATETIME     NOT NULL,
    `token_type`      VARCHAR(50)  NOT NULL
)