/** @type {import('tailwindcss').Config} */
export default {
    content: [
    "./index.html",             // for Vite
    "./src/**/*.{ts,tsx}"     // only scan TS/TSX files as it's Vite Typescript project
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}

