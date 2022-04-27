module.exports = {
  root: true,
  parser: "@typescript-eslint/parser",
  plugins: ["@typescript-eslint", "import", "unused-imports"],
  extends: [
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking",
    "plugin:prettier/recommended",
    "plugin:import/typescript",
  ],
  parserOptions: {
    sourceType: "module",
    project: [`${__dirname}/tsconfig.json`],
  },
  rules: {
    "prettier/prettier": 0,
    "unused-imports/no-unused-imports": 1,
    "@typescript-eslint/no-non-null-assertion": 2,
    "@typescript-eslint/no-explicit-any": 2,
    "import/no-cycle": ["error", { ignoreExternal: true, maxDepth: 42 }],
    "import/no-unresolved": 2,
    "@typescript-eslint/no-empty-function": 0,
    "@typescript-eslint/prefer-regexp-exec": 0,
    "@typescript-eslint/explicit-module-boundary-types": 2,
    "no-shadow": 0,
    "@typescript-eslint/no-shadow": 1,
    "no-unused-vars": 0,
    "@typescript-eslint/no-unused-vars": 1,
    "@typescript-eslint/no-namespace": 0,
    "@typescript-eslint/require-await": 1,
  },
  settings: {
    "import/parsers": {
      "@typescript-eslint/parser": [".ts", ".tsx"],
    },
    "import/resolver": {
      typescript: {
        project: [`${__dirname}/tsconfig.json`],
      },
    },
  },
};
