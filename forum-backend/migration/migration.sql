USE bluedit;

INSERT INTO users (name, email, password, role) VALUES
('GamerPro_99', 'gamerpro@bluedit.com', '$2a$12$ExAmPlEhAsHeDkWq789234789234789234789234789234', 'admin'),
('RetroLover', 'retro@bluedit.com', '$2a$12$ExAmPlEhAsHeDkWq111222333444555666777888999000', 'user'),
('Ezio_Du_93', 'ezio@bluedit.com', '$2a$12$ExAmPlEhAsHeDkWq999888777666555444333222111000', 'user'),
('IndieFanatic', 'indie@bluedit.com', '$2a$12$ExAmPlEhAsHeDkWq123123123123123123123123123123', 'user'),
('NoobMaster', 'noob@bluedit.com', '$2a$12$ExAmPlEhAsHeDkWq456456456456456456456456456456', 'user');

INSERT INTO categories (name, description) VALUES
('Rétro', 'Pour les nostalgiques des pixels, de la Super Nintendo à la PS1. Viens souffler dans ta cartouche !'),
('Indé', 'Les pépites cachées, le pixel art et les chefs-d''œuvre des petits studios indépendants (Hollow Knight, Hades, Celeste...).'),
('Assassin''s Creed', 'Le repaire des Assassins. Théories, lore, configurations PC et débats sur le meilleur opus de la licence.'),
('Général', 'Discussions libres sur l''actualité du jeu vidéo, les setups, ou juste pour tailler le bout de gras.');

INSERT INTO threads (name, content, fk_id_category, fk_id_user) VALUES
-- Threads Rétro (Catégorie 1)
('La meilleure console de salon de tous les temps ?', 'Pour moi c''est la PlayStation 2, la ludothèque était juste légendaire. Vous en pensez quoi ?', 1, 2),
('Votre pire souvenir sur un jeu NES ?', 'Ce niveau de la tortue dans Teenage Mutant Ninja Turtles m''a traumatisé à vie. Qui a réussi à le passer ici ?', 1, 5),

-- Threads Indé (Catégorie 2)
('Hollow Knight : Des news de Silksong ?', 'Je commence à perdre patience, ça fait des années qu''on attend. Quelqu''un a vu passer des leaks crédibles récemment ?', 2, 4),
('Celeste : Le chapitre 9 est une torture', 'Je bloque dessus depuis 3 jours, mes doigts saignent. Donnez-moi des conseils s''il vous plaît !', 2, 5),

-- Threads Assassin's Creed (Catégorie 3)
('Pourquoi AC II reste le sommet de la saga', 'L''ambiance de la Renaissance Italienne, la musique de Jesper Kyd, le développement d''Ezio... Rien n''a égalé ça depuis.', 3, 3),

-- Threads Général (Catégorie 4)
('Bienvenue sur Bluedit ! Présentez-vous !', 'Salut à tous les gamers ! Bienvenue sur le forum. Lâchez vos meilleurs setups et vos jeux du moment en commentaire.', 4, 1);

INSERT INTO messages (content, fk_id_user, fk_id_thread) VALUES
-- Réponses au thread PS2 (Thread 1)
('La Super Nintendo rigole au fond de la pièce. Chrono Trigger, Super Metroid... imbattable !', 4, 1),
('Totalement d''accord pour la PS2, rien que pour San Andreas et GoW 2.', 1, 1),
('La GameCube était sous-cotée par contre, le catalogue de fou.', 3, 1),

-- Réponses au thread TMNT NES (Thread 2)
('Mdr le niveau de la bombe dans l''eau ? Une horreur sans nom !', 2, 2),
('Je l''ai passé une fois en screed, j''ai perdu toutes mes vies juste après.', 1, 2),

-- Réponses au thread Silksong (Thread 3)
('Silksong n''existe pas, c''est une hallucination collective.', 2, 3),
('Il sortira quand on sera tous à la retraite à mon avis.', 5, 3),

-- Réponses au thread AC II (Thread 4)
('AC Origins était pas mal aussi pour le renouveau, mais Ezio reste le roi.', 1, 5),
('Rien que d''entendre "Ezio''s Family" j''ai des frissons.', 2, 5),

-- Réponses au thread Bienvenue (Thread 6)
('Trop lourd le design sombre de Bluedit, hâte de voir la suite du dev !', 3, 6),
('Salut ! Setup full Corsair ici, je ponce Elden Ring en ce moment.', 4, 6);