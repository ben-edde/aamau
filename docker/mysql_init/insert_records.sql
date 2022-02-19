use aamauDB;
insert into Cake values (1,"Apple cake", 1, 100);
insert into Cake values (2,"Buccellato", 2, 200);
insert into Cake values (3,"Cassata", 3, 300);
-- insert into Cake values (4,"Dundee cake", 4, 400);
-- insert into Cake values (5,"Esterházy torte", 5, 500);
-- insert into Cake values (6,"Flourless chocolate cake", 6, 600);

insert into Ingredient values (1,"Apple", 5, 50);
insert into Ingredient values (2,"Sugar", 0.5, 300);
insert into Ingredient values (3,"Flour", 0.7, 500);
insert into Ingredient values (4,"Butter", 10, 100);
insert into Ingredient values (5,"Egg", 3, 200);

insert into Recipe values (1,1, 1, 3);
insert into Recipe values (2,1, 2, 1);
insert into Recipe values (3,1, 3, 1);
insert into Recipe values (4,2, 2, 2);
insert into Recipe values (5,2, 3, 1);
insert into Recipe values (6,2, 5, 2);
insert into Recipe values (7,3, 1, 1);
insert into Recipe values (8,3, 4, 3);
insert into Recipe values (9,3, 3, 2);

insert into User values (1,"Alpha", "12345678", "Alpha@test.com", "Alpha Home");
insert into User values (2,"Bravo", "22345678", "Bravo@test.com", "Bravo Home");
insert into User values (3,"Charlie", "32345678", "Charlie@test.com", "Charlie Home");
insert into User values (4,"Delta", "42345678", "Delta@test.com", "Delta Home");
insert into User values (5,"Echo", "52345678", "Echo@test.com", "Echo Home");

insert into Orders values (NULL, "2022-01-01", "2022-01-02", 1, 1, 2, 200);
insert into Orders values (NULL, "2022-01-11", "2022-01-14", 2, 3, 1, 300);
insert into Orders values (NULL, "2022-01-21", "2022-01-23", 3, 2, 5, 1000);

