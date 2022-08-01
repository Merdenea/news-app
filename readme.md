## news-app

App that fetches articles from rss feeds, caches them(in-memory) and exposes http endpoints to retrieve them.

Many, many things unfinished or untouched, including this readme.

To BUILD && RUN: \
    1. `docker compose up --build` \
    2. `Talk to me.`

Working endpoints: (app will run on localhost, port 8080): \
    `GET /sources`
    `GET /news`
