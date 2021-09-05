defmodule RakonWeb.PageController do
  use RakonWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
