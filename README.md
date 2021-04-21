# fpl-discord-bot

A Discord bot that helps you to live Fantasy Premier League funnier in your Discord server.
## Usage

You can run it with Docker in your local.
### Docker
```
 docker build -t fpl-bot .
```
LEAGUE_ID= Your league's id.

TOKEN= Discord Bot token.
```
docker run -d --restart \
> -e LEAGUE_ID=""\
> -e TOKEN="" \
> fpl-bot
```
## Features

- League standings
- Manager performance
- Weekly points
- Fixture
  