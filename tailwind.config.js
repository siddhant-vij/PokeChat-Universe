/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["cmd/web/templates/**/*.{html,templ}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
};
