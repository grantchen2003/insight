import concurrent, grpc, os

from source_code_parser.protobuf import source_code_parser_pb2_grpc
from source_code_parser.service.source_code_parser_service import SourceCodeParser


def start():
    server = grpc.server(
        concurrent.futures.ThreadPoolExecutor(),
        options=[("grpc.max_receive_message_length", -1)],
    )
    source_code_parser_pb2_grpc.add_SourceCodeParserServicer_to_server(SourceCodeParser(), server)
    address = f"{os.environ['DOMAIN']}:{os.environ['PORT']}"
    server.add_insecure_port(address)
    print(f"starting source code parser server on {address}")
    server.start()
    server.wait_for_termination()
