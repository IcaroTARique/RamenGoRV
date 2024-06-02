-- Criação das tabelas
CREATE TABLE broth (
                       id INTEGER PRIMARY KEY,
                       image_inactive TEXT,
                       image_active TEXT,
                       name TEXT,
                       description TEXT,
                       price INTEGER
);

CREATE TABLE protein (
                         id INTEGER PRIMARY KEY,
                         image_inactive TEXT,
                         image_active TEXT,
                         name TEXT,
                         description TEXT,
                         price INTEGER
);

CREATE TABLE answer(
                       id INTEGER PRIMARY KEY,
                       broth_id INTEGER,
                       protein_id INTEGER,
                       description TEXT,
                       image TEXT,
                       FOREIGN KEY (broth_id) REFERENCES broth(id),
                       FOREIGN KEY (protein_id) REFERENCES protein(id)
);

CREATE TABLE orders(
                       id INTEGER PRIMARY KEY AUTO_INCREMENT,
                       broth_id INTEGER,
                       protein_id INTEGER,
                       FOREIGN KEY (broth_id) REFERENCES broth(id),
                       FOREIGN KEY (protein_id) REFERENCES protein(id)
);

-- Inserção dos dados

INSERT INTO broth (id, image_inactive, image_active, name, description, price) VALUES
                                                                                   (1, 'https://tech.redventures.com.br/icons/salt/inactive.svg', 'https://tech.redventures.com.br/icons/salt/active.svg', 'Salt', 'A sliced flavourful pork meat with a selection of season vegetables.',   10),
                                                                                   (2, 'https://tech.redventures.com.br/icons/shoyu/inactive.svg', 'https://tech.redventures.com.br/icons/shoyu/active.svg', 'Shoyu', 'A delicious vegetarian lamen with a selection of season vegetables.',  10),
                                                                                   (3, 'https://tech.redventures.com.br/icons/miso/inactive.svg', 'https://tech.redventures.com.br/icons/miso/active.svg', 'Miso', 'Three units of fried chicken, moyashi, ajitama egg and other vegetables.', 12);

INSERT INTO protein (id, image_inactive, image_active, name, description, price) VALUES
                                                                                     (1, 'https://tech.redventures.com.br/icons/pork/inactive.svg', 'https://tech.redventures.com.br/icons/pork/active.svg', 'Chasu', 'Simple like the seawater, nothing more', 10),
                                                                                     (2, 'https://tech.redventures.com.br/icons/yasai/inactive.svg', 'https://tech.redventures.com.br/icons/yasai/active.svg', 'Yasai Vegetarian', 'The good old and traditional soy sauce', 10),
                                                                                     (3, 'https://tech.redventures.com.br/icons/chicken/inactive.svg', 'https://tech.redventures.com.br/icons/chicken/active.svg', 'Karaague', 'Paste made of fermented soybeans', 12);


INSERT INTO answer (id, broth_id, protein_id, description, image) VALUES
                                                                      (1, 1, 1, 'Salt and Chasu Ramen','https://tech.redventures.com.br/icons/ramen/ramenChasu.png'),
                                                                      (2, 1, 2, 'Salt and Yasai Vegetarian Ramen','https://tech.redventures.com.br/icons/ramen/ramenYasai Vegetarian.png'),
                                                                      (3, 1, 3, 'Salt and Karaague Ramen','https://tech.redventures.com.br/icons/ramen/ramenKaraague.png'),
                                                                      (4, 2, 1, 'Shoyu and Chasu Ramen','https://tech.redventures.com.br/icons/ramen/ramenChasu.png'),
                                                                      (5, 2, 2, 'Shoyu and Yasai Vegetarian Ramen','https://tech.redventures.com.br/icons/ramen/ramenYasai Vegetarian.png'),
                                                                      (6, 2, 3, 'Shoyu and Karaague Ramen','https://tech.redventures.com.br/icons/ramen/ramenKaraague.png'),
                                                                      (7, 3, 1, 'Miso and Chasu Ramen','https://tech.redventures.com.br/icons/ramen/ramenChasu.png'),
                                                                      (8, 3, 2, 'Miso and Yasai Vegetarian Ramen','https://tech.redventures.com.br/icons/ramen/ramenYasai Vegetarian.png'),
                                                                      (9, 3, 3, 'Miso and Karaague Ramen','https://tech.redventures.com.br/icons/ramen/ramenKaraague.png');