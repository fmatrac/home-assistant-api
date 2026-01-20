-- Skrypt wypełniający tabelę produkty popularnymi produktami
-- Uruchom: psql -U home_assistant_user -d home_assistant -f scripts/seed_produkty.sql

-- Nabiał
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Mleko 2%', 'Nabiał', true),
('Mleko 3.2%', 'Nabiał', false),
('Masło', 'Nabiał', true),
('Jajka', 'Nabiał', true),
('Ser żółty', 'Nabiał', true),
('Ser biały', 'Nabiał', false),
('Jogurt naturalny', 'Nabiał', true),
('Jogurt owocowy', 'Nabiał', false),
('Śmietana 18%', 'Nabiał', true),
('Śmietana 30%', 'Nabiał', false),
('Kefir', 'Nabiał', false),
('Serek wiejski', 'Nabiał', false),
('Twaróg', 'Nabiał', false),
('Mozzarella', 'Nabiał', false),
('Parmezan', 'Nabiał', false);

-- Pieczywo
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Chleb pszenny', 'Pieczywo', true),
('Chleb żytni', 'Pieczywo', false),
('Chleb tostowy', 'Pieczywo', true),
('Bułki', 'Pieczywo', true),
('Bagietka', 'Pieczywo', false),
('Rogaliki', 'Pieczywo', false),
('Chleb razowy', 'Pieczywo', false),
('Chałka', 'Pieczywo', false);

-- Mięso i wędliny
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Pierś z kurczaka', 'Mięso', true),
('Udka z kurczaka', 'Mięso', false),
('Mięso mielone wołowe', 'Mięso', true),
('Mięso mielone wieprzowe', 'Mięso', false),
('Schab', 'Mięso', false),
('Karkówka', 'Mięso', false),
('Boczek', 'Mięso', false),
('Szynka', 'Wędliny', true),
('Salami', 'Wędliny', false),
('Parówki', 'Wędliny', true),
('Kiełbasa', 'Wędliny', false),
('Kabanosy', 'Wędliny', false),
('Pasztet', 'Wędliny', false);

-- Warzywa
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Ziemniaki', 'Warzywa', true),
('Marchewka', 'Warzywa', true),
('Cebula', 'Warzywa', true),
('Czosnek', 'Warzywa', true),
('Pomidory', 'Warzywa', true),
('Ogórki', 'Warzywa', true),
('Papryka', 'Warzywa', false),
('Sałata', 'Warzywa', true),
('Kapusta', 'Warzywa', false),
('Brokuły', 'Warzywa', false),
('Kalafior', 'Warzywa', false),
('Pieczarki', 'Warzywa', true),
('Szpinak', 'Warzywa', false),
('Cukinia', 'Warzywa', false),
('Bakłażan', 'Warzywa', false),
('Por', 'Warzywa', false),
('Seler', 'Warzywa', false),
('Pietruszka', 'Warzywa', false),
('Rukola', 'Warzywa', false),
('Rzodkiewka', 'Warzywa', false);

-- Owoce
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Jabłka', 'Owoce', true),
('Banany', 'Owoce', true),
('Pomarańcze', 'Owoce', true),
('Cytryny', 'Owoce', true),
('Mandarynki', 'Owoce', false),
('Winogrona', 'Owoce', false),
('Truskawki', 'Owoce', false),
('Maliny', 'Owoce', false),
('Borówki', 'Owoce', false),
('Gruszki', 'Owoce', false),
('Brzoskwinie', 'Owoce', false),
('Śliwki', 'Owoce', false),
('Kiwi', 'Owoce', false),
('Ananas', 'Owoce', false),
('Mango', 'Owoce', false),
('Awokado', 'Owoce', false);

-- Napoje
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Woda mineralna', 'Napoje', true),
('Woda gazowana', 'Napoje', true),
('Sok pomarańczowy', 'Napoje', true),
('Sok jabłkowy', 'Napoje', false),
('Cola', 'Napoje', false),
('Sprite', 'Napoje', false),
('Fanta', 'Napoje', false),
('Kawa mielona', 'Napoje', true),
('Kawa rozpuszczalna', 'Napoje', false),
('Herbata czarna', 'Napoje', true),
('Herbata zielona', 'Napoje', false),
('Herbata owocowa', 'Napoje', false),
('Piwo', 'Napoje', false),
('Wino czerwone', 'Napoje', false),
('Wino białe', 'Napoje', false);

-- Produkty suche
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Makaron spaghetti', 'Produkty suche', true),
('Makaron penne', 'Produkty suche', false),
('Ryż biały', 'Produkty suche', true),
('Ryż basmati', 'Produkty suche', false),
('Kasza gryczana', 'Produkty suche', false),
('Kasza jęczmienna', 'Produkty suche', false),
('Mąka pszenna', 'Produkty suche', true),
('Cukier', 'Produkty suche', true),
('Sól', 'Produkty suche', true),
('Pieprz', 'Produkty suche', true),
('Płatki owsiane', 'Produkty suche', true),
('Corn flakes', 'Produkty suche', false),
('Musli', 'Produkty suche', false);

-- Przetwory i konserwy
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Pomidory krojone (puszka)', 'Konserwy', true),
('Passata pomidorowa', 'Konserwy', true),
('Ketchup', 'Konserwy', true),
('Majonez', 'Konserwy', true),
('Musztarda', 'Konserwy', true),
('Ogórki kiszone', 'Konserwy', false),
('Ogórki konserwowe', 'Konserwy', false),
('Groszek konserwowy', 'Konserwy', false),
('Kukurydza konserwowa', 'Konserwy', false),
('Fasola konserwowa', 'Konserwy', false),
('Tuńczyk w puszce', 'Konserwy', false),
('Dżem truskawkowy', 'Konserwy', false),
('Miód', 'Konserwy', true),
('Nutella', 'Konserwy', false);

-- Mrożonki
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Lody waniliowe', 'Mrożonki', false),
('Lody czekoladowe', 'Mrożonki', false),
('Mrożone warzywa mieszanka', 'Mrożonki', true),
('Mrożona pizza', 'Mrożonki', false),
('Frytki mrożone', 'Mrożonki', false),
('Ryba mrożona', 'Mrożonki', false),
('Pierogi mrożone', 'Mrożonki', false);

-- Słodycze i przekąski
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Czekolada mleczna', 'Słodycze', false),
('Czekolada gorzka', 'Słodycze', false),
('Ciastka', 'Słodycze', false),
('Wafelki', 'Słodycze', false),
('Chipsy', 'Przekąski', false),
('Paluszki', 'Przekąski', false),
('Orzeszki ziemne', 'Przekąski', false),
('Żelki', 'Słodycze', false);

-- Chemia domowa
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Płyn do naczyń', 'Chemia', true),
('Płyn do płukania', 'Chemia', true),
('Proszek do prania', 'Chemia', true),
('Płyn uniwersalny', 'Chemia', true),
('Płyn do WC', 'Chemia', false),
('Płyn do szyb', 'Chemia', false),
('Worki na śmieci', 'Chemia', true),
('Papier toaletowy', 'Chemia', true),
('Ręczniki papierowe', 'Chemia', true),
('Gąbki do naczyń', 'Chemia', false);

-- Higiena osobista
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Mydło', 'Higiena', true),
('Szampon', 'Higiena', true),
('Żel pod prysznic', 'Higiena', true),
('Pasta do zębów', 'Higiena', true),
('Szczoteczka do zębów', 'Higiena', false),
('Dezodorant', 'Higiena', true),
('Krem do rąk', 'Higiena', false),
('Chusteczki higieniczne', 'Higiena', false);

-- Oleje i tłuszcze
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Olej rzepakowy', 'Oleje', true),
('Oliwa z oliwek', 'Oleje', true),
('Olej słonecznikowy', 'Oleje', false),
('Olej kokosowy', 'Oleje', false);

-- Przyprawy
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Papryka słodka', 'Przyprawy', true),
('Oregano', 'Przyprawy', true),
('Bazylia', 'Przyprawy', true),
('Tymianek', 'Przyprawy', false),
('Rozmaryn', 'Przyprawy', false),
('Curry', 'Przyprawy', false),
('Kurkuma', 'Przyprawy', false),
('Cynamon', 'Przyprawy', false),
('Liść laurowy', 'Przyprawy', false),
('Ziele angielskie', 'Przyprawy', false);

-- Dla dzieci/niemowląt
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Pieluszki', 'Dla dzieci', false),
('Chusteczki nawilżane', 'Dla dzieci', false),
('Mleko modyfikowane', 'Dla dzieci', false),
('Kaszka dla dzieci', 'Dla dzieci', false);

SELECT 'Dodano ' || COUNT(*) || ' produktów' FROM home_assistant.produkty;

