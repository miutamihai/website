defmodule Blog.Components.Post do
  use Phoenix.Component
  import Phoenix.HTML
  import Components.Layout

  embed_templates("./post.html")
end
