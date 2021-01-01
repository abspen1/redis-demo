import redis
import schedule
import time
import os

"""
DB: 0 --
    BOTS Standings/Points
    Projects 
    RPS
    Todos (Not in use)
    twitter-bot (Not in use)
    tw-bot
    Eagle
    Falcon
    Austin Intercept Live
    Puma
DB: 1 --
DB: 2 --
    BOTS Southeast
DB: 3 --
    BOTS West
    Drake Intercept Live
DB: 4 --
    BOTS Midwest
DB: 5 --
    BOTS Northeast
    Modern-Warfare
DB: 6 --
    Bday-Reminder
DB: 7 --
DB: 8 --
DB: 9 --
    Discord-Bot
DB: 10 --
    BOTS Rosters && Players
DB: 11 --
DB: 12 --
    BOTS Trending
DB: 13 --
DB: 14 --
DB: 15 --
"""


def main():
    client = redis.Redis(
        host=os.getenv("REDIS_HOST"), port=6379, db=0, password=os.getenv("REDIS_PASS")
    )

    client.set("cloud_read", "2442554")
    print(client.dbsize())


if __name__ == "__main__":
    main()

schedule.every(5).days.do(main)

while True:
    try:
        schedule.run_pending()
        time.sleep(1)
    except Exception as identifier:
        print(identifier)
        time.sleep(1)