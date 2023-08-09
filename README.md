# Search Engine tumail
Project to search text from an emails database, like a search engine would do.

Technologies:
Zincsearch 0.4.7
Go 1.20.6
API Router chi
Vue 3
Tailwindcss
Node js v18.17.0
npm v9.8.1

Notes:
* To execute zincsearch in Ubuntu use the following command 
ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 ./zincsearch
* Zincsearch is hosted in :4080
* Backend service is locally hosted in :8080
* Frontend service is locally hosted in :5173

Go configurations to use and import modules
* Create go.mod file using this command
    go mod init serverDomain/packageName For example: 'git.com/project'
Now, we can import external packages in go
Note: for each package/folder in the project it can exist 1 main package
* go.sum file is generated and updated automatically 

How to create a vue project using npm
    npm create vue@latest
    cd <your-project-name>
    npm install
    npm run dev

Tailwind configuration to use css styles:
* Installation 
    npm install -D tailwindcss
    npm install @headlessui/vue @heroicons/vue
    npx tailwindcss init
* Configure template paths in tailwind.config.js file
    /** @type {import('tailwindcss').Config} */
    module.exports = {
    content: ["./src/**/*.{html,js}"],
    theme: {
        extend: {},
    },
    plugins: [],
    }
* Add the tailwind directives in your main css file
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
* Building tailwind css file (the file you are using in the project is output.css)
    npx tailwindcss -i ./src/main.css -o ./dist/output.css --watch
* Import /dist/output.css file in your main component


