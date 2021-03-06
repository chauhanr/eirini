# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: ingress.proto for package 'loggregator.v2'

require 'grpc'
require 'ingress_pb'

module Loggregator
  module V2
    module Ingress
      class Service

        include GRPC::GenericService

        self.marshal_class_method = :encode
        self.unmarshal_class_method = :decode
        self.service_name = 'loggregator.v2.Ingress'

        rpc :Sender, stream(Envelope), IngressResponse
        rpc :BatchSender, stream(EnvelopeBatch), BatchSenderResponse
        rpc :Send, EnvelopeBatch, SendResponse
      end

      Stub = Service.rpc_stub_class
    end
  end
end
