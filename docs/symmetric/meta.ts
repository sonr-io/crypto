import { defineMeta } from "blume";

export default defineMeta({
  title: "Symmetric & Secrets",
  icon: "lock",
  order: 3,
  pages: ["index", "aead", "deterministic-aead", "key-derivation", "secrets"],
});
