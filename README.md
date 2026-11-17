# Projet_Forum - Bluedit

## Features

### FT1 - Système d'inscription

### FT-2 – Système de connexion

### FT-3 – Création d’un fil de discussion

- Un fil de discussion doit posséder un état parmi les valeurs suivantes : ouvert, fermé ou archivé.
- Un fil ouvert permet aux utilisateurs authentifiés de publier des messages.
- Un fil fermé reste visible et consultable, mais ne permet plus de publier de nouveaux messages.
- Un fil archivé ne doit plus être affiché ni accessible aux utilisateurs.

### FT-4 – Consulter un fil de discussion

- Les fils ouverts peuvent être consultés.
- Les fils fermés peuvent être consultés.
- Les fils archivés ne doivent pas être visibles.
- Les messages associés au fil doivent être affichés selon les règles de tri définies.

### FT-5 Publication d’un message dans un fil de discussion

- Un message doit être publié dans un fil de discussion existant.
- Un message ne peut être publié que dans un fil ouvert.
- Un fil fermé ne doit pas accepter de nouveaux messages.
- Un fil archivé ne doit pas accepter de nouveaux messages.

### FT-6 – Réagir à un message d’un fil de discussion

- Seul un utilisateur authentifié peut réagir à un message.
- Un utilisateur ne peut pas liker et disliker le même message en même temps.
- Un utilisateur ne peut avoir qu’une seule réaction active sur un même message.
- Un like ajoute +1 au score de popularité du message.
- Un dislike retire -1 au score de popularité du message.
- Le score de popularité doit pouvoir être utilisé pour trier les messages.

### FT-7 – Gestion des messages et des fils de discussion

### FT-8 – Tri des messages d’un fil de discussion

- Les messages doivent être affichés par défaut du plus récent au plus ancien.
- L’utilisateur doit pouvoir afficher les messages par ordre chronologique.
- L’utilisateur doit pouvoir afficher les messages selon leur score de popularité.
- Le score de popularité est calculé à partir des likes et des dislikes.
- Un like ajoute +1 au score de popularité.
- Un dislike retire -1 au score de popularité.
- Le tri par popularité doit tenir compte du score calculé pour chaque message.

### FT-9 – Pagination des fils de discussion et des messages

- Les listes de fils de discussion doivent pouvoir être paginées.
- Les listes de messages doivent pouvoir être paginées.
- L’utilisateur doit pouvoir choisir un affichage par lots de 10, 20 ou 30 éléments.
- L’utilisateur doit pouvoir choisir d’afficher la totalité des éléments.
- Par défaut, les éléments doivent être affichés par lots de 10.

### FT-10 – Afficher les fils de discussion par tag/catégorie

- Seuls les fils visibles doivent apparaître dans les résultats.
- Les fils archivés ne doivent pas apparaître dans les résultats.
- Le filtrage doit conserver un affichage cohérent avec les règles de pagination.

### FT-11 – Rechercher un fil de discussion

- Seul un utilisateur authentifié peut effectuer une recherche.
- La recherche doit pouvoir porter sur le titre d’un fil de discussion.
- La recherche doit pouvoir porter sur un tag/catégorie.
- Le système doit déterminer automatiquement si la recherche correspond à un titre, un tag ou une catégorie.
- Les fils archivés ne doivent pas apparaître dans les résultats.
- Les résultats doivent respecter les règles de visibilité des fils de discussion.

### FT-12 – Gestion et administration de la plateforme

- Seuls les utilisateurs possédant le rôle administrateur peuvent accéder au tableau de bord d’administration.
- Un administrateur peut modifier l’état d’un fil de discussion.
- Un administrateur peut supprimer un fil de discussion.
- Un administrateur peut supprimer un message.
- Un administrateur peut bannir un compte utilisateur.
- Un administrateur peut agir sur un fil même s’il n’en est pas le propriétaire.
- Un administrateur peut agir sur un message même s’il n’en est pas l’auteur.
- Un utilisateur banni ne doit plus pouvoir accéder aux fonctionnalités nécessitant une authentification.

## Auteurs
