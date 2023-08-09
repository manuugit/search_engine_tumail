# Search Engine tumail :open_file_folder:
Project to search text from an emails database, like a search engine would do.

## Visualization - mvp
![image](https://github.com/manuugit/search_engine_tumail/assets/77813632/46316f1c-dde1-4c54-a074-69f7a68a4f76)

## Technologies:
* [Zincsearch 0.4.7](https://zincsearch-docs.zinc.dev)
* Go 1.20.6 
* [API Router chi](https://go-chi.io/#/pages/routing)
* [Vue 3](https://vuejs.org)
* [Tailwindcss](https://tailwindcss.com/docs/installation)
* Node js v18.17.0 
* npm v9.8.1

:page_facing_up: :page_facing_up: :page_facing_up: <br>
### Notes:
* To execute zincsearch in Ubuntu use the following command 
`ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 ./zincsearch`
* Zincsearch is hosted in :4080
* Backend service is locally hosted in :8080
* Frontend service is locally hosted in :5173
* indexer.go loads data to zincsearch

Go configurations to use and import modules
* Create `go.mod` file using this command
    `go mod init serverDomain/packageName` For example: 'git.com/project' <br>
Now, we can import external packages in go
Note: for each package/folder in the project it can exist only one main package
* `go.sum` file is generated and updated automatically 

### Notes - front
How to create a vue project using npm
```
    npm create vue@latest
    cd <your-project-name>
    npm install
    npm run dev
```

Tailwind configuration to use css styles:
* Installation
```
    npm install -D tailwindcss
    npm install @headlessui/vue @heroicons/vue
    npx tailwindcss init
```
* Configure template paths in tailwind.config.js file
```
    /** @type {import('tailwindcss').Config} */
    module.exports = {
    content: ["./src/**/*.{html,js}"],
    theme: {
        extend: {},
    },
    plugins: [],
    }
```
* Add the tailwind directives in your main css file
```
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
```
* Building tailwind css file (the file you are using in the project is output.css) <br>
    `npx tailwindcss -i ./src/main.css -o ./dist/output.css --watch`
  
* Import `/dist/output.css` file in your main component
