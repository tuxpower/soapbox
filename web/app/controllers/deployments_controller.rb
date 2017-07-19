require 'deployment_pb'

class DeploymentsController < ApplicationController
  before_action :set_application, only: [:index, :create, :new, :show]
  before_action :set_environments, only: [:new, :create]

  def index
    req = Soapbox::ListDeploymentRequest.new(application_id: params[:application_id].to_i)
    res = $api_deployment_client.list_deployments(req)
    if res.deployments.count == 0
      redirect_to new_application_deployment_path
    else
      @deployments = res.deployments.sort_by { |d| -Time.parse(d.created_at).to_i }
    end
  end

  def new
    @form = CreateDeploymentForm.new
  end

  def create
    @form = CreateDeploymentForm.new(params[:deployment])
    if @form.valid?
      env = Soapbox::Environment.new(id: @form.environment_id)
      app = Soapbox::Application.new(id: params[:application_id].to_i)
      req = Soapbox::Deployment.new(committish: @form.committish, application: app, env: env)
      res = $api_deployment_client.start_deployment(req)
      redirect_to application_deployment_path(id: res.id)
    else
      render :new
    end
  end

  def show
    req = Soapbox::GetDeploymentStatusRequest.new(id: params[:id].to_i)
    res = $api_deployment_client.get_deployment_status(req)
    @state = res.state

    respond_to do |format|
      format.html
      format.json { render plain: @state }
    end
  end

  private

  def list_environments(application_id)
    req = Soapbox::ListEnvironmentRequest.new(application_id: application_id)
    $api_environment_client.list_environments(req).environments
  end

  def set_application
    req = Soapbox::GetApplicationRequest.new(id: params[:application_id].to_i)
    @app = $api_client.get_application(req)
  end

  def set_environments
    @environments = list_environments(params[:application_id].to_i)
  end
end