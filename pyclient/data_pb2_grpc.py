# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import data_pb2 as data__pb2


class DataServiceStub(object):
    """The data service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetTrade = channel.unary_unary(
                '/api.DataService/GetTrade',
                request_serializer=data__pb2.TradeRequest.SerializeToString,
                response_deserializer=data__pb2.TradeResponse.FromString,
                )


class DataServiceServicer(object):
    """The data service definition.
    """

    def GetTrade(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DataServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetTrade': grpc.unary_unary_rpc_method_handler(
                    servicer.GetTrade,
                    request_deserializer=data__pb2.TradeRequest.FromString,
                    response_serializer=data__pb2.TradeResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.DataService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class DataService(object):
    """The data service definition.
    """

    @staticmethod
    def GetTrade(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.DataService/GetTrade',
            data__pb2.TradeRequest.SerializeToString,
            data__pb2.TradeResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)