require 'version_pb'

class AboutController < ApplicationController
  def index
    @version = $api_client.versions.get_version(Soapbox::Empty.new)
  end
end
