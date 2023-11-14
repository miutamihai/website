defmodule Website.MixProject do
  use Mix.Project

  def project do
    [
      app: :website,
      version: "0.1.0",
      elixir: "~> 1.15",
      start_permanent: Mix.env() == :prod,
      deps: deps(),
      aliases: aliases()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger]
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:nimble_publisher, "~> 0.1.3"},
      {:phoenix_live_view, "~> 0.18.2"},
      {:esbuild, "~> 0.5"},
      {:tailwind, "~> 0.1.8"}
    ]
  end

  defp aliases() do
    [
      "site.build": ["build", "tailwind default --minify", "esbuild default --minify"]
    ]
  end
end
