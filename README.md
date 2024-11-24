# Typhero Server

A Server for my little school project [Typehero](https://github.com/oxelf/typehero), a typing test.

This Server provides basic functionality such as saving typing test
results and calculating a leaderboard from them.

Data is stored in a small sqlite instance.

Endpoints:

GET /healthcheck (Responds with 200 if the server is up, but it should almost never
be down as its deployed through K8s)

POST /result (The results are posted to this endpoint and are associated with users
through a random uuid which will be generated and stored on the clientside)

GET /leaderboard?mode=${mode}&language=${language}&wordAmount=${wordAmount} (Returns a Leaderboard for given params)

mode = "words" || "quote"
language = "english" || "german"
wordAmount = "10" || "25" || "50" || "75" or "0" incase of mode being "quote"
