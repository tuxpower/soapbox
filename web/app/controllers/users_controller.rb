require 'user_pb'

class UsersController < ApplicationController
  before_action :assign_user, only: %i[show omniauth]
  skip_before_action :require_login, except: %i[show omniauth]

  def new
    @form = CreateUserForm.new
  end

  def create
    @form = CreateUserForm.new(params[:user])
    if @form.valid?
      req = Soapbox::CreateUserRequest.new(name: @form.name, email: @form.email, password: @form.password)
      @user = $api_client.users.create_user(req, skip_auth: true)
      session[:current_user_email] = @user.email
      redirect_to profile_user_path
    else
      render :new
    end
  end

  def login
    @form = LoginUserForm.new
  end

  def attempt_login
    @form = LoginUserForm.new(params[:user])
    if @form.valid?
      req = Soapbox::LoginUserRequest.new(email: @form.email, password: @form.password)
      res = $api_client.users.login_user(req, metadata: { skip_auth: 'true' })
      if !res.error.blank?
        @form.errors.add(:password, :invalid)
        render :login
      else
        session[:current_user_email] = res.user.email
        session[:login_token] = res.hmac
        redirect_to profile_user_path
      end
    else
      render :login
    end
  end

  def omniauth
    auth = request.env['omniauth.auth']
    @user.github_oauth_access_token = auth['credentials']['token']
    @user = $api_client.users.assign_github_omniauth_token_to_user(@user, user_metadata)
    refresh_user
    redirect_to profile_user_path
  end

  private

    def assign_user
      @user = current_user
    end
end
