#!/usr/bin/env python
import pika
import uuid
import json


class callFunc(object):

    def __init__(self):
        # username = "default_user_NJ3h3CGstRtDaLQ5PJG"
        # password = "5ejAO_rV0y0jHNHyVEhgQa2eJ79CUnsb"
        username = "{{RABBITMQ_USERNAME}}"
        password = "{{RABBITMQ_PASSWORD}}"
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(host="localhost", port=5672, credentials=pika.PlainCredentials(username, password)),
        )

        self.channel = self.connection.channel()

        result = self.channel.queue_declare(queue='rpc_queue', durable=True, auto_delete=False)
        self.callback_queue = result.method.queue

        self.channel.basic_consume(
            queue=self.callback_queue,
            on_message_callback=self.on_response,
            auto_ack=True)

        self.response = None
        self.corr_id = None

    def on_response(self, ch, method, props, body):
        if self.corr_id == props.correlation_id:
            self.response = body

    def call(self, body):
        self.response = None
        self.corr_id = str(uuid.uuid4())
        self.channel.basic_publish(
            exchange='',
            routing_key='rpc_queue',
            properties=pika.BasicProperties(
                content_type= 'application/json',
                reply_to=self.callback_queue,
                correlation_id=self.corr_id,
            ),
            body=json.dumps(body))
        while self.response is None:
            self.connection.process_data_events(time_limit=None)
        return self.response


def {{FUNCTION_NAME}}(body):
    return callFunc().call(body)