CREATE TABLE pricebtc
(
    price TEXT,
    id SERIAL PRIMARY KEY,
    coinname TEXT

);

INSERT INTO pricebtc(price,coinname)
VALUES
    ('20000','bitcoin'),
    ('2000','ethereum')

