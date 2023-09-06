/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.hbs"],
  theme: {
    extend: {
      colors: {
        background: {
          700: "#1D181E",
          800: "#19141A",
        },
        secondary: "#5C4194",
        accent: "#DD6336",
      },
    },
  },
  plugins: [],
};
