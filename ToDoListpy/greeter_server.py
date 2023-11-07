# Copyright 2015 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""The Python implementation of the GRPC helloworld.Greeter server."""

from concurrent import futures
import logging

import grpc
import todo_list_pb2
import todo_list_pb2_grpc
import ToDoListMicroservise

class Greeter(todo_list_pb2_grpc.GreeterServicer):
    def SayHello(self, request, context):
        # return todo_list_pb2.HelloReply(message="Hello, %s!" % request.name)
        print("request.user_input:", request.user_input)
        print("request.task_description:", request.task_description)
        return todo_list_pb2.HelloReply(message=ToDoListMicroservise.ToDoService(request.user_input, request.task_description))


def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    todo_list_pb2_grpc.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()
