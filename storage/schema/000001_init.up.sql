CREATE TABLE drinks
(
    id               serial not null unique,
    name             varchar(20) not null,
    price            int not null,
    alcoholic        varchar(10) not null,
    ice              varchar(10) not null,
    flavour          varchar(10) not null,
    primary_type     varchar(10) not null,
    secondary_type   varchar(10),
    recipe           varchar(150),
    shortcut         varchar(60),
    description      varchar(150)
);


INSERT INTO drinks VALUES (1, 'Bad Touch', 250, 'Yes', 'Yes', 'Sour', 'Classy', 'Vintage', '2 Bronson Extract, 2 Powdered Delta, 2 Flanergide and 4 Karmotrine. All on the rocks and mixed.', '2xW, 2xE, 2xR, 4xT, A, all mixed.', 'We''re nothing but mammals after all.');
INSERT INTO drinks VALUES (2, 'Beer', 200, 'Yes', 'No', 'Bubbly', 'Classic', 'Vintage', '1 Adelhyde, 2 Bronson Extract, 1 Powdered Delta, 2 Flanergide and 4 Karmotrine. All mixed.', '1xQ, 2xW, 1xE, 2xR, 4xT, all mixed.', 'Traditionally brewed beer has become a luxury, but this one''s pretty close to the real deal...');
INSERT INTO drinks VALUES (3, 'Bleeding Jane', 200, 'No', 'No', 'Spicy', 'Classic', 'Sobering', '1 Bronson Extract, 3 Powdered Delta and 3 Flanergide. All blended.', '1xW, 3xE, 3xR, all blended.', 'Say the name of this drink three times in front of a mirror and you''ll look like a fool.');
INSERT INTO drinks VALUES (4, 'Bloom Light', 230, 'Yes', 'Yes', 'Spicy', 'Promo', 'Bland', '4 Adelhyde, 1 Powdered Delta, 2 Flanergide and 3 Karmotrine. All aged, on the rocks and mixed.', '4xQ, 1xE, 2xR, 3xT, A, S, all mixed.', 'It''s so unnecessarily brown....');
INSERT INTO drinks VALUES (5, 'Blue Fairy', 170, 'Optional', 'No', 'Sweet', 'Girly', 'Soft', '4 Adelhyde, 1 Flanergide and optional Karmotrine. All aged and mixed.', '4xQ, 1xR, optional T, S, all mixed.', 'One of these will make all your teeth turn blue. Hope you brushed them well.');
INSERT INTO drinks VALUES (6, 'Brandtini', 250, 'Yes', 'No', 'Sweet', 'Classy', 'Happy', '6 Adelhyde, 3 Powdered Delta and 1 Karmotrine. All aged and mixed.', '6xQ, 3xE, 1xT, S, all mixed.', '8 out of 10 smug assholes would recommend it but they''re too busy being smug assholes.');
INSERT INTO drinks VALUES (7, 'Cobalt Velvet', 280, 'Yes', 'Yes', 'Bubbly', 'Classy', 'Burning', '2 Adelhyde, 3 Flanergide and 5 Karmotrine. All on the rocks and mixed.', '2xQ, 3xR, 5xT, A, all mixed.', 'It''s like champaigne served on a cup that had a bit of cola left.');
INSERT INTO drinks VALUES (8, 'Crevice Spike', 140, 'Optional', 'No', 'Sour', 'Manly', 'Sobering', '2 Powdered Delta, 4 Flanergide and optional Karmotrine. All blended.', '2xE, 4xR, optional T, all blended.', 'It will knock the drunkenness out of you or knock you out cold.');
INSERT INTO drinks VALUES (9, 'Fluffy Dream', 170, 'Optional', 'No', 'Sour', 'Girly', 'Soft', '3 Adelhyde, 3 Powdered Delta and optional Karmotrine. All aged and mixed.', '3xQ, 3xE, optional T, S, all mixed.', 'A couple of these will make your tongue feel velvet-y. More of them and you''ll be sleeping soundly.');
INSERT INTO drinks VALUES (10, 'Fringe Weaver', 260, 'Yes', 'No', 'Bubbly', 'Classy', 'Strong', '1 Adelhyde and 9 Karmotrine. All aged and mixed.', '1xQ, 9xT, S, all mixed.', 'It''s like drinking ethylic alcohol with a spoonful of sugar.');
INSERT INTO drinks VALUES (11, 'Frothy Water', 150, 'No', 'No', 'Bubbly', 'Classic', 'Bland', '1 Adelhyde, 1 Bronson Extract, 1 Powdered Delta and 1 Flanergide. All aged and mixed.', '1xQ, 1xW, 1xE, 1xR, S, all mixed.', 'PG-rated shows'' favorite Beer ersatz since 2040.');
INSERT INTO drinks VALUES (12, 'Grizzly Temple', 220, 'Yes', 'No', 'Bitter', 'Promo', 'Bland', '3 Adelhyde, 3 Bronson Extract, 3 Powdered Delta and 1 Karmotrine. All blended.', '3xQ, 3xW, 3xE, 1xT, all blended.', 'This one''s kinda unbearable. It''s mostly for fans of the movie it was used on.');
INSERT INTO drinks VALUES (13, 'Gut Punch', 80, 'Optional', 'No', 'Bitter', 'Manly', 'Strong', '5 Bronson Extract, 1 Flanergide and optional Karmotrine. All aged and mixed.', '5xW, 1xR, optional T, S, all mixed.', 'It''s supposed to mean "a punch made of innards", but the name actually described what you feel while drinking it.');
INSERT INTO drinks VALUES (14, 'Marsblast', 170, 'Yes', 'No', 'Spicy', 'Manly', 'Strong', '6 Bronson Extract, 1 Powdered Delta, 4 Flanergide and 2 Karmotrine. All blended.', '6xW, 1xE, 4xR, 2xT, all blended.', 'One of these is enough to leave your face red like the actual planet.');
INSERT INTO drinks VALUES (15, 'Mercuryblast', 250, 'Yes', 'Yes', 'Sour', 'Classy', 'Burning', '1 Adelhyde, 1 Bronson Extract, 3 Powdered Delta, 3 Flanergide and 2 Karmotrine. All on the rocks and blended.', '1xQ, 1xW, 3xE, 3xR, 2xT, A, all blended.', 'No thermometer was harmed in the creation of this drink.');
INSERT INTO drinks VALUES (16, 'Moonblast', 180, 'Yes', 'Yes', 'Sweet', 'Girly', 'Happy', '6 Adelhyde, 1 Powdered Delta, 1 Flanergide and 2 Karmotrine. All on the rocks and blended.', '6xQ, 1xE, 1xR, 2xT, A, all blended.', 'No relation to the Hadron cannon you can see on the moon for one week every month.');
INSERT INTO drinks VALUES (17, 'Piano Man', 320, 'Yes', 'Yes', 'Sour', 'Promo', 'Strong', '2 Adelhyde, 3 Bronson Extract, 5 Powdered Delta, 5 Flanergide and 3 Karmotrine. All on the rocks and mixed.', '2xQ, 3xW, 5xE, 5xR, 3xT, A, all mixed.', 'This drink does not represent the opinions of the Bar Pianists Union or its associates.');
INSERT INTO drinks VALUES (18, 'Piano Woman', 320, 'Yes', 'No', 'Sweet', 'Promo', 'Happy', '5 Adelhyde, 5 Bronson Extract, 2 Powdered Delta, 3 Flanergide and 3 Karmotrine. All aged and mixed.', '5xQ, 5xW, 2xE, 3xR, 3xT, S, all mixed.', 'It was originally called Pretty Woman, but too many people complained there should be a Piano Woman if there was a Piano Man.');
INSERT INTO drinks VALUES (19, 'Pile Driver', 160, 'Yes', 'No', 'Bitter', 'Manly', 'Burning', '3 Bronson Extract, 3 Flanergide and 4 Karmotrine. All mixed.', '3xW, 3xR, 4xT, all mixed.', 'It doesn''t burn as hard on the tongue but you better not have a sore throat when drinking it...');
INSERT INTO drinks VALUES (20, 'Sparkle Star', 150, 'Optional', 'No', 'Sweet', 'Girly', 'Happy', '2 Adelhyde, 1 Powdered Delta and optional Karmotrine. All aged and mixed.', '2xQ, 1xE, optional T, S, all mixed.', 'They used to actually sparkle, but too many complaints about skin problem made them redesign the drink without sparkling.');
INSERT INTO drinks VALUES (21, 'Sugar Rush', 150, 'Optional', 'No', 'Sweet', 'Girly', 'Happy', '2 Adelhyde, 1 Powdered Delta and optional Karmotrine. All mixed.', '2xQ, 1xE, optional T, all mixed.', 'Sweet, light and fruity. As girly as it gets.');
INSERT INTO drinks VALUES (22, 'Sunshine Cloud', 150, 'Optional', 'Yes', 'Bitter', 'Girly', 'Soft', '2 Adelhyde, 2 Bronson Extract and optional Karmotrine. All on the rocks and blended.', '2xQ, 2xW, optional T, A, all blended.', 'Tastes like old chocolate milk with its good smell intact. Some say it tastes like caramel too...');
INSERT INTO drinks VALUES (23, 'Suplex', 160, 'Yes', 'Yes', 'Bitter', 'Manly', 'Burning', '4 Bronson Extract, 3 Flanergide and 3 Karmotrine. All on the rocks and mixed.', '4xW, 3xR, 3xT, A, all mixed.', 'A small twist on the Piledriver, putting more emphasis on the tongue burning and less on the throat burning.');
INSERT INTO drinks VALUES (24, 'Zen Star', 210, 'Yes', 'Yes', 'Sour', 'Promo', 'Bland', '4 Adelhyde, 4 Bronson Extract, 4 Powdered Delta, 4 Flanergide and 4 Karmotrine. All on the rocks and mixed.', '4xQ, 4xW, 4xE, 4xR, 4xT, A, all mixed.', 'You''d think something so balanced would actually taste nice... you''d be dead wrong.');
INSERT INTO drinks VALUES (25, 'Flaming Moai', 150, 'Yes', 'No', 'Sour', 'Classy', 'Classic', '1 Adelhyde, 1 Bronson Extract, 2 Powdered Delta, 3 Flanergide and 5 Karmotrine. All mixed.', '1xQ, 1xW, 2xE, 3xR, 5xT, all mixed.', 'N/A');
INSERT INTO drinks VALUES (26, 'Absinthe', 500, 'Yes', 'No', 'N/A', 'Bottled', 'N/A', 'N/A', 'N/A', 'N/A');
INSERT INTO drinks VALUES (27, 'A Fedora', 500, 'Yes', 'No', 'N/A', 'Bottled', 'N/A', 'N/A', 'N/A', 'N/A');
INSERT INTO drinks VALUES (28, 'Mulan Tea', 500, 'Yes', 'No', 'N/A', 'Bottled', 'N/A', 'N/A', 'N/A', 'N/A');
INSERT INTO drinks VALUES (29, 'Rum', 500, 'Yes', 'No', 'N/A', 'Bottled', 'N/A', 'N/A', 'N/A', 'N/A');
