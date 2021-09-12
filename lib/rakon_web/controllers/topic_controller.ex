# Copyright 2022 Clivern. All rights reserved.
# Use of this source code is governed by the MIT
# license that can be found in the LICENSE file.

defmodule RakonWeb.TopicController do
  use RakonWeb, :controller

  def index(conn, _params) do
    render(conn, "topic.html")
  end
end
