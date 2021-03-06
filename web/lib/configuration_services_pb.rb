# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: configuration.proto for package 'soapbox'

require 'grpc'
require 'configuration_pb'

module Soapbox
  module Configurations
    class Service

      include GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'soapbox.Configurations'

      rpc :ListConfigurations, ListConfigurationRequest, ListConfigurationResponse
      rpc :GetLatestConfiguration, GetLatestConfigurationRequest, Configuration
      rpc :CreateConfiguration, CreateConfigurationRequest, Configuration
      rpc :DeleteConfiguration, DeleteConfigurationRequest, Empty
    end

    Stub = Service.rpc_stub_class
  end
end
