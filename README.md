# Scissor
 URL Shortened using Go! | Fiber, GORM, PostgreSQL, Svelte. AltSchool Final Project for SOE.# 


---
## PostgreSQL Docker Image
I used postgres:14 and can run it
```bash
$ docker run --name name-of-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=test -d postgres:14
```
You can name the container anything you want, or not name it all. `--name` flag is really used to run docker commands easier, rather than using the randomly generated UUID.

## For frontend I use Sveltejs

. Run
```bash
$ npx degit sveltejs/template view
```
## To create folder name view cd to the folder

. Run
```bash
yarn install
```
then 
```bash
yarn run dev
```
To add modal run 
```bash
yarn add svelte-modals
---
:zap: Happy Coding!
