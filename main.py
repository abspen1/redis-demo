import redis
import schedule
import time
import os


def main():
    client = redis.Redis(host=os.getenv('REDIS_HOST'), port=6379, db=7,
                        password=os.getenv("REDIS_PASS"))

    client.set('language', 'Python')

    print(client.get('language'))



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