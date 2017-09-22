class ApplicationController < ActionController::Base
  protect_from_forgery with: :exception
  before_action :require_login
  helper_method :current_user
  helper_method :refresh_user
  helper_method :octokit

  def require_login
    redirect_to login_user_path unless current_user && session[:login_token]
    wrap_api_calls
  end

  # Finds the User with the ID stored in the session with the key
  # :current_user_email
  def current_user
    @_current_user ||= session[:current_user_email] &&
                       get_user(session[:current_user_email])
  end

  def refresh_user
    @_current_user = session[:current_user_email] &&
                     get_user(session[:current_user_email])
  end

  def get_user(email)
    req = Soapbox::GetUserRequest.new(email: email)
    $api_client.users.get_user(req)
  end

  def octokit
    if current_user && current_user.github_oauth_access_token
      @octokit ||= Octokit::Client.new(access_token: current_user.github_oauth_access_token)
    end
  end

  private

  def wrap_api_calls
    $api_client.to_h.values.each do |service|
      service.class.class_exec do
        old = instance_method(:request_response)
        define_method(:request_response) do |route, req, marshal, unmarshal, _|
          old.bind(self).call(
            route, req, marshal, unmarshal,
            user_id: current_user.id,
            login_token: session[:login_token]
          )
        end
      end
    end
  end
end
