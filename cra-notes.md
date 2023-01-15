# Setup React (CRA) application

CRA (create-react-app) takes care of all the build and configuration process for you, so you don't need to write a batch script or a webpack configuration file.

Here is an example of how you can set up your project using CRA:

```
npx create-react-app client
cd client
npm start
```

CRA also provides an option to run the server-side code and the client-side code together using the same command, it is called react-scripts-ts.

## Directory structure

Here's an example of the directory structure you might see when using Create-React-App (CRA).

to create this structure you have to first create the server and client dirs, cd to client dir and type:

```
npx create-react-app my-app --template typescript
npm install
```


```
my-app/
|
+-- client/
|   |
|   +-- node_modules/
|   +-- package.json
|   +-- package-lock.json
|   +-- public/
|   |   |
|   |   +-- index.html
|   |   +-- favicon.ico
|   |   +-- manifest.json
|   |
|   +-- src/
|       |
|       +-- components/
|       |   |
|       |   +-- App.tsx
|       |   +-- ...
|       |
|       +-- index.tsx
|       +-- ...
|   +-- .gitignore
|   +-- README.md
|   +-- .env
|   +-- tsconfig.json
|
+-- server/
```
