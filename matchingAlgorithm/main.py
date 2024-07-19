from kafka import KafkaConsumer, KafkaProducer
import json
from similar import *

consumer = KafkaConsumer('input_topic', bootstrap_servers=['localhost:9092'],
                         value_deserializer=lambda m: json.loads(m.decode('ascii')))
producer = KafkaProducer(bootstrap_servers=['localhost:9092'], value_serializer=lambda v: json.dumps(v).encode('ascii'))

for message in consumer:
    data = message.value
    username = data['user_name']
    playlist_id = data['playlist_id']
    user_id = data['user_id']

    print(username, playlist_id)

    result = process_user_data(username, playlist_id, user_id)

    producer.send('output_topic', {'result': result})
