from source_code_parser.protobuf import source_code_parser_pb2, source_code_parser_pb2_grpc

class SourceCodeParser(source_code_parser_pb2_grpc.SourceCodeParserServicer):
    def SyntaxParse(self, requests, context):
        print('yo')
        for request in requests:
            user_id = request.user_id
            file_path = request.file_path
            print(user_id, file_path)
            
            response = source_code_parser_pb2.SyntaxParseResponse(success = False)
            yield response
