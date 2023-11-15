# Ivaldea
This is the Elixir code that produces Ivaldea's website.

## Deployment
* `mix deps.get`
* `mix site.build`

The resulting website will be stored in the `/output` folder.

## Installation

If [available in Hex](https://hex.pm/docs/publish), the package can be installed
by adding `website` to your list of dependencies in `mix.exs`:

```elixir
def deps do
  [
    {:website, "~> 0.1.0"}
  ]
end
```

Documentation can be generated with [ExDoc](https://github.com/elixir-lang/ex_doc)
and published on [HexDocs](https://hexdocs.pm). Once published, the docs can
be found at <https://hexdocs.pm/website>.
