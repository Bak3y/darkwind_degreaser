# darkwind_degreaser

## Description
Exports news feed from Enjin and converts the data into WordPress posts, maintaining author, date, and contents from Enjin.

---
## Required env vars
* ENJIN_API_KEY
* ENJIN_API_URL
* ENJIN_SITE_ID
* ENJIN_USER_ID
* ENJIN_POST_LIMIT
* WP_URL

The docker-compose.yml and the .vscode/launch.json file expects these to be in .env at the root of this repo

## Running the app using docker-compose
Start up your local docker daemon and simply call

`docker-compose up --build`
