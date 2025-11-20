import time
import json
import requests
import pika
import os

LAT = "-19.76"
LON = "-44.05"
RABBITMQ_HOST = "rabbitmq" 
QUEUE_NAME = "weather_queue"

def get_weather():
    url = f"https://api.open-meteo.com/v1/forecast?latitude={LAT}&longitude={LON}&current=temperature_2m,relative_humidity_2m,wind_speed_10m,weather_code"
    try:
        response = requests.get(url)
        return response.json()
    except Exception as e:
        print(f"Erro ao buscar clima: {e}")
        return None

def send_to_queue(data):
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters(host=RABBITMQ_HOST))
        channel = connection.channel()
        
        channel.queue_declare(queue=QUEUE_NAME, durable=True)
        
        channel.basic_publish(exchange='', routing_key=QUEUE_NAME, body=message)
        print(f" [x] Dados enviados: {data['current']['temperature_2m']}°C")
        
        connection.close()
    except Exception as e:
        print(f"Erro ao conectar no RabbitMQ (tentando novamente em breve): {e}")

if __name__ == "__main__":
    print("Iniciando Coletor de Clima...")
    while True:
        weather_data = get_weather()
        if weather_data:
            send_to_queue(weather_data)
        
        time.sleep(60)