## Architecture

L'application **Bluedit** (`forum-bluedit`) est structurée en monorepo Go (`go.mod` unique) et déployée sous forme d'un seul processus web sur **Scalingo** :

- **Backend API & serveur principal (`forum-backend`)** : expose l'API REST (threads, messages, catégories, utilisateurs, likes, authentification), la route de santé `GET /health` et rattache le routeur client (`RouterClient`) ainsi que les templates/assets du frontend sur un unique port (`$PORT` en production).
- **Frontend (`forum-frontend`)** : génère et sert les pages/templates côté client via `RouterClient` et `helper.Load()`, et communique avec l'API.
- **Base de données** : instance **MySQL** provisionnée via l'add-on Scalingo, reliée au processus web via les variables d'environnement `DB_*` (chaîne DSN au format `user:pass@tcp(host:port)/dbname?parseTime=true`).
- **CI/CD** : dépôt hébergé sur GitHub avec un workflow **GitHub Actions** (`.github/workflows/go.yml`, ignoré uniquement pour `**.md`) exécutant `gofmt -l`, `go vet ./...` et `go build ./...`, suivi du déploiement automatique sur Scalingo depuis la branche `main` une fois les checks au vert.

---

## Procédure de déploiement reproductible

### 1. Prérequis
- Compte **GitHub** et compte **Scalingo** (CLI `scalingo` installée et connectée).
- Application Scalingo créée (ex. `forum-bluedit`) avec l'add-on **MySQL** attaché et un `Procfile` lançant `./bin/forum-backend` sur le processus `web`.

### 2. Configuration des variables d'environnement sur Scalingo
Définir les variables requises sur l'application (sans mettre de secrets dans le dépôt) :

```bash
scalingo --app forum-bluedit env-set JWT_SECRET="<secret_fort>"
# Vérifier que DB_HOST, DB_PORT, DB_USER, DB_PWD, DB_NAME sont présents
scalingo --app forum-bluedit env
```

### 3. Liaison GitHub et CI/CD automatique
1. Pousser le code et le workflow CI (`.github/workflows/go.yml`) sur la branche `main` de GitHub.
2. Vérifier que le workflow GitHub Actions (`gofmt -l`, `go vet ./...`, `go build ./...`) passe au vert.
3. Dans le tableau de bord Scalingo (`forum-bluedit`), lier le dépôt GitHub, sélectionner la branche `main`, activer **Auto Deploy** et activer l'attente des statuts/checks GitHub avant déploiement.

### 4. Déploiement manuel (si besoin)
```bash
git push scalingo main
scalingo --app forum-bluedit logs
```

### 5. Vérification post-déploiement

Contrôler la route de santé : GET https:///health doit retourner un statut 200.

Vérifier dans les logs Scalingo (scalingo --app forum-bluedit logs) que la connexion MySQL et le démarrage du serveur sur $PORT ont réussi sans erreur.

### 6. Variables d'environnement

Les variables suivantes sont lues par l'application (aucune valeur ne doit être versionnée dans Git) :

| Variable | Description | Obligatoire |
| --- | --- | --- |
| PORT | Port d'écoute HTTP fourni dynamiquement par Scalingo (repli local si non défini) | Oui (en prod) |
| DB_HOST | Hôte du serveur MySQL | Oui |
| DB_PORT | Port du serveur MySQL | Oui |
| DB_USER | Nom d'utilisateur de la base MySQL | Oui |
| DB_PWD | Mot de passe de la base MySQL | Oui |
| DB_NAME | Nom de la base de données MySQL | Oui |
| JWT_SECRET | Clé secrète utilisée pour signer/vérifier les jetons d'authentification | Oui |

### 7. Coût mensuel : estimé vs réel

À compléter avec les tarifs exacts de ton compte / période de facturation (selon le plan du conteneur web et le plan de l'add-on MySQL sur Scalingo).

Poste | Coût mensuel estimé | Coût mensuel réel | Notes
| --- | --- | --- | --- |
Conteneur web Scalingo (web) | 14.40€ | 14.40€ | Selon le plan choisi (M)
Add-on base de données MySQL | 7.20€ | 7.20€ | Selon le plan MySQL Scalingo
GitHub / GitHub Actions | Gratuit (Dépôt public) | Gratuit (Dépôt public) | Selon le forfait du compte / quota de minutes
Total | 21.60€ | 21.60€ | Comparer l'estimation initiale au relevé de facturation réel