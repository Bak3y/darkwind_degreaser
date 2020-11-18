# darkwind_degreaser

## Description
Exports news feed from Enjin and converts the data into WordPress posts, maintaining author, date, and contents from Enjin.

---
## Required env vars
Place these in your .env file, one per line
* ENJIN_API_KEY - Your Enjin API key
* ENJIN_API_URL - Your Enjin site API url, something like yoursite.com/api/v1/api.php
* ENJIN_SITE_ID - Your Enjin siite ID
* ENJIN_USER_ID - Your Enjin user ID
* ENJIN_POST_LIMIT - a integer representing how many posts you want to pull out of Enjin and put into Wordpress
* WP_URL - full API URl for post creation, including `/wp/v2/posts`
* WP_AUTH - base64 encrypted username:password to auth to wordpress

The docker-compose.yml and the .vscode/launch.json file expects the .env file to be in the root of this repo

## Running the app using docker-compose
Start up your local docker daemon and simply call

`docker-compose up --build`
