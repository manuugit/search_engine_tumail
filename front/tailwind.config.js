/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html","./src/**/*.{html,js,vue}"],
  theme: {
    extend: {
      container: {
        center: true,
      },
    },
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}
