"""
 Copyright 2022 Jo Shinonome

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
"""
import logging


import grpc
import data_pb2
import data_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel("localhost:1897") as channel:
        stub = data_pb2_grpc.DataServiceStub(channel)
        response = stub.GetTrade(data_pb2.TradeRequest(sym="a"))
    for trade in response.trades:
        print(trade)


if __name__ == "__main__":
    logging.basicConfig()
    run()
