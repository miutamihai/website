module.exports = {
	content: [
		"./**/*_templ.go",
	],
	plugins: [
		require("@tailwindcss/typography"),
	],
	theme: {
		fontFamily: {
			sans: ['Merriweather Sans', 'sans-serif']
		},
		extend: {
			colors: {
				primary: '#2e3440',
				secondary: '#d08770',
				white: '#eceff4'
			}
		}
	}
};
