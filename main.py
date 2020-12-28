import redis
import os


client = redis.Redis(host=os.getenv('REDIS_HOST'), port=6379, db=7,
                     password=os.getenv("REDIS_PASS"))

client.set('language', 'Python')

print(client.get('language'))