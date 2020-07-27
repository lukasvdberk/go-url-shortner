# TODO Maybe setup with the .env file
use go_url_shortner;

CREATE TABLE ShortUrl
(
    id INTEGER AUTO_INCREMENT,
    realUrl VARCHAR(1024),
    PRIMARY KEY (id)
);