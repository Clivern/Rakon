defmodule Rakon.Repo do
  use Ecto.Repo,
    otp_app: :rakon,
    adapter: Ecto.Adapters.Postgres
end
