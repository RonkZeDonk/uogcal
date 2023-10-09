/** @type {import('tailwindcss').Config} */
export default {
  content: ["index.html", "./src/**/*.{js,ts,jsx,tsx}"],
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
