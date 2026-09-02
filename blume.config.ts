import { defineConfig } from "blume";

export default defineConfig({
  title: "Sonr Crypto",
  description:
    "Developer documentation for the Sonr cryptography library — curves, threshold signatures, MPC, zero-knowledge proofs, and identity primitives in Go.",

  github: {
    owner: "sonr-io",
    repo: "crypto",
    branch: "main",
  },

  theme: {
    accent: "teal",
    radius: "md",
    mode: "system",
  },

  navigation: {
    sidebar: {
      display: "group",
    },
  },

  search: {
    provider: "orama",
  },

  lastModified: true,

  seo: {
    og: { enabled: true },
    sitemap: true,
    robots: true,
    structuredData: true,
  },

  ai: {
    llmsTxt: true,
  },

  // GitHub Pages project site: served from https://sonr-io.github.io/crypto
  deployment: {
    output: "static",
    site: "https://sonr-io.github.io",
    base: "/crypto",
  },
});
