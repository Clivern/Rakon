defmodule Rakon.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Ecto repository
      Rakon.Repo,
      # Start the Telemetry supervisor
      RakonWeb.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: Rakon.PubSub},
      # Start the Endpoint (http/https)
      RakonWeb.Endpoint
      # Start a worker by calling: Rakon.Worker.start_link(arg)
      # {Rakon.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Rakon.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    RakonWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
