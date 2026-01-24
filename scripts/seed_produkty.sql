-- Skrypt wypełniający tabelę produkty popularnymi produktami
-- Uruchom: psql -U home_assistant_user -d home_assistant -f scripts/seed_produkty.sql

-- Nabiał
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Mleko 2%', 'Nabiał', false),
('Mleko 3.2%', 'Nabiał', false),
('Masło', 'Nabiał', false),
('Jajka', 'Nabiał', false),
('Ser żółty', 'Nabiał', false),
('Ser biały', 'Nabiał', false),
('Jogurt naturalny', 'Nabiał', false),
('Jogurt owocowy', 'Nabiał', false),
('Śmietana 18%', 'Nabiał', false),
('Śmietana 30%', 'Nabiał', false),
('Kefir', 'Nabiał', false),
('Serek wiejski', 'Nabiał', false),
('Twaróg', 'Nabiał', false),
('Mozzarella', 'Nabiał', false),
('Parmezan', 'Nabiał', false);

-- Pieczywo
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Chleb pszenny', 'Pieczywo', false),
('Chleb żytni', 'Pieczywo', false),
('Chleb tostowy', 'Pieczywo', false),
('Bułki', 'Pieczywo', false),
('Bagietka', 'Pieczywo', false),
('Rogaliki', 'Pieczywo', false),
('Chleb razowy', 'Pieczywo', false),
('Chałka', 'Pieczywo', false);

-- Mięso i wędliny
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Pierś z kurczaka', 'Mięso', false),
('Udka z kurczaka', 'Mięso', false),
('Mięso mielone wołowe', 'Mięso', false),
('Mięso mielone wieprzowe', 'Mięso', false),
('Schab', 'Mięso', false),
('Karkówka', 'Mięso', false),
('Boczek', 'Mięso', false),
('Szynka', 'Wędliny', false),
('Salami', 'Wędliny', false),
('Parówki', 'Wędliny', false),
('Kiełbasa', 'Wędliny', false),
('Kabanosy', 'Wędliny', false),
('Pasztet', 'Wędliny', false);

-- Warzywa
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Ziemniaki', 'Warzywa', false),
('Marchewka', 'Warzywa', false),
('Cebula', 'Warzywa', false),
('Czosnek', 'Warzywa', false),
('Pomidory', 'Warzywa', false),
('Ogórki', 'Warzywa', false),
('Papryka', 'Warzywa', false),
('Sałata', 'Warzywa', false),
('Kapusta', 'Warzywa', false),
('Brokuły', 'Warzywa', false),
('Kalafior', 'Warzywa', false),
('Pieczarki', 'Warzywa', false),
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
('Jabłka', 'Owoce', false),
('Banany', 'Owoce', false),
('Pomarańcze', 'Owoce', false),
('Cytryny', 'Owoce', false),
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
('Woda mineralna', 'Napoje', false),
('Woda gazowana', 'Napoje', false),
('Sok pomarańczowy', 'Napoje', false),
('Sok jabłkowy', 'Napoje', false),
('Cola', 'Napoje', false),
('Sprite', 'Napoje', false),
('Fanta', 'Napoje', false),
('Kawa mielona', 'Napoje', false),
('Kawa rozpuszczalna', 'Napoje', false),
('Herbata czarna', 'Napoje', false),
('Herbata zielona', 'Napoje', false),
('Herbata owocowa', 'Napoje', false),
('Piwo', 'Napoje', false),
('Wino czerwone', 'Napoje', false),
('Wino białe', 'Napoje', false);

-- Produkty suche
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Makaron spaghetti', 'Produkty suche', false),
('Makaron penne', 'Produkty suche', false),
('Ryż biały', 'Produkty suche', false),
('Ryż basmati', 'Produkty suche', false),
('Kasza gryczana', 'Produkty suche', false),
('Kasza jęczmienna', 'Produkty suche', false),
('Mąka pszenna', 'Produkty suche', false),
('Cukier', 'Produkty suche', false),
('Sól', 'Produkty suche', false),
('Pieprz', 'Produkty suche', false),
('Płatki owsiane', 'Produkty suche', false),
('Corn flakes', 'Produkty suche', false),
('Musli', 'Produkty suche', false);

-- Przetwory i konserwy
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Pomidory krojone (puszka)', 'Konserwy', false),
('Passata pomidorowa', 'Konserwy', false),
('Ketchup', 'Konserwy', false),
('Majonez', 'Konserwy', false),
('Musztarda', 'Konserwy', false),
('Ogórki kiszone', 'Konserwy', false),
('Ogórki konserwowe', 'Konserwy', false),
('Groszek konserwowy', 'Konserwy', false),
('Kukurydza konserwowa', 'Konserwy', false),
('Fasola konserwowa', 'Konserwy', false),
('Tuńczyk w puszce', 'Konserwy', false),
('Dżem truskawkowy', 'Konserwy', false),
('Miód', 'Konserwy', false),
('Nutella', 'Konserwy', false);

-- Mrożonki
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Lody waniliowe', 'Mrożonki', false),
('Lody czekoladowe', 'Mrożonki', false),
('Mrożone warzywa mieszanka', 'Mrożonki', false),
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
('Płyn do naczyń', 'Chemia', false),
('Płyn do płukania', 'Chemia', false),
('Proszek do prania', 'Chemia', false),
('Płyn uniwersalny', 'Chemia', false),
('Płyn do WC', 'Chemia', false),
('Płyn do szyb', 'Chemia', false),
('Worki na śmieci', 'Chemia', false),
('Papier toaletowy', 'Chemia', false),
('Ręczniki papierowe', 'Chemia', false),
('Gąbki do naczyń', 'Chemia', false);

-- Higiena osobista
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Mydło', 'Higiena', false),
('Szampon', 'Higiena', false),
('Żel pod prysznic', 'Higiena', false),
('Pasta do zębów', 'Higiena', false),
('Szczoteczka do zębów', 'Higiena', false),
('Dezodorant', 'Higiena', false),
('Krem do rąk', 'Higiena', false),
('Chusteczki higieniczne', 'Higiena', false);

-- Oleje i tłuszcze
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Olej rzepakowy', 'Oleje', false),
('Oliwa z oliwek', 'Oleje', false),
('Olej słonecznikowy', 'Oleje', false),
('Olej kokosowy', 'Oleje', false);

-- Przyprawy
INSERT INTO home_assistant.produkty (nazwa, kategoria, ulubiony) VALUES
('Papryka słodka', 'Przyprawy', false),
('Oregano', 'Przyprawy', false),
('Bazylia', 'Przyprawy', false),
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

