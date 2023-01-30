DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('My Everything', 'Ariana Grande', 56.99),
  ('Sounds good feels good', '5 Seconds of Summer', 17.99),
  ('Midgnight Memories', 'One Direction', 39.99);
