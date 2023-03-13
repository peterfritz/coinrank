<div align="center">
  <h1>CoinRank API</h1>
</div>

<br />

The CoinRank API is a powerful tool that allows users to rate and evaluate cryptocurrencies using the familiar "upvote" and "downvote" system.

Project developed as part of a talent selection/recruitment process for [Klever](https://klever.io/en).

<br />

<hr />
<br />

## Deploy

Spin up your own instance of the CoinRank API using Vercel and Railway with just two clicks.

<br />

<div align="center">
  <a href="https://railway.app/template/lBuqJY?referralCode=ptr" target="_blank">
    <img alt="Deploy the back-end (API) on Railway" title="Deploy the back-end (API) on Railway" src="https://img.shields.io/badge/DEPLOY%20API%20ON%20RAILWAY-000000?style=for-the-badge&logo=railway" />
  </a>
  <a href="https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fpeterfritz%2Fcoinrank&env=NEXT_PUBLIC_API_URL&envDescription=Your%20CoinRank%20API's%20URL%20or%20https%3A%2F%2Fapi.coinrank.ptr.red%2F&project-name=coinrank&repository-name=coinrank&&root-directory=frontend&install-command=pnpm%20install&build-command=pnpm%20build&demo-title=coinrank&demo-description=A%20powerful%20tool%20that%20allows%20users%20to%20rate%20and%20evaluate%20cryptocurrencies%20using%20the%20familiar%20%22upvote%22%20and%20%22downvote%22%20system.&demo-url=https%3A%2F%2Fcoinrank.ptr.red&demo-image=https%3A%2F%2Fuser-images.githubusercontent.com%2F61599784%2F224681680-a338b9b7-eaa6-4d1b-a510-277aa1aec2dd.png" target="_blank">
    <img alt="Deploy the front-end (interface) on Vercel" title="Deploy the front-end (interface) on Vercel" src="https://img.shields.io/badge/DEPLOY%20INTERFACE%20ON%20VERCEL-000000?style=for-the-badge&logo=vercel" />
  </a>
</div>

<br />

## Documentation / Endpoints

The API's request/response structure and the available endpoints are fully documented using the OpenAPI specification. You can explore it [here](https://swaggerviewer.ptr.red/spec/aHR0cHM6Ly9naXRodWIuY29tL3BldGVyZnJpdHovY29pbnJhbmsvcmF3L21haW4vYmFja2VuZC9zd2FnZ2VyLnlhbWw%3D) or by following the link bellow.

<br />

<div align="center">
  <a href="https://swaggerviewer.ptr.red/spec/aHR0cHM6Ly9naXRodWIuY29tL3BldGVyZnJpdHovY29pbnJhbmsvcmF3L21haW4vYmFja2VuZC9zd2FnZ2VyLnlhbWw%3D" target="_blank">
    <img alt="Open on Swagger Viewer" title="Open on Swagger Viewer" src="https://img.shields.io/badge/OPEN%20ON%20SWAGGER%20VIEWER-099268?style=for-the-badge&logo=data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjU4IiBoZWlnaHQ9IjEwMCIgdmlld0JveD0iMCAwIDI1OCAxMDAiIGZpbGw9Im5vbmUiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CjxwYXRoIGQ9Ik0xMTkuMTg3IDExLjQ4MDZDMTIxLjIyMiAxMy4xMDA3IDEyNC43NDIgMTUuMDQxNiAxMjkgMTUuMDQxNkMxMzMuMjU4IDE1LjA0MTYgMTM2Ljc3NyAxMy4xMDA3IDEzOC44MTMgMTEuNDgwNkMxMzguODQ3IDExLjQ1MzggMTM4Ljg4MiAxMS40MzE2IDEzOC45MTggMTEuNDE1QzEzOC45NSAxMS4zOTkzIDEzOC45ODMgMTEuMzg4MiAxMzkuMDE2IDExLjM3OTlMMTQ2LjA3NCA1LjQzMTYyQzE1MC4yMzcgMS45MjMzIDE1NS41MDQgMCAxNjAuOTQ2IDBIMjQ3LjAzM0MyNTMuMDkgMCAyNTggNC45MTMxMyAyNTggMTAuOTc0MUMyNTggMTIuNzIyNyAyNTcuNTkxIDE0LjM3NTIgMjU2Ljg2NCAxNS44NDI5QzI1Ni40ODkgMTYuNjAwNyAyNTYuMDc4IDE3LjM1NzcgMjU1LjY2NSAxOC4xMTc0TDI1NS42NDcgMTguMTUwN0MyNTQuMTY0IDIwLjg3OCAyNTIuNjYgMjMuNjQyMyAyNTIuNjYgMjYuNjgxMVY1MS41NTM2QzI1Mi42NiA1MS43NDAzIDI1Mi42NTkgNTEuOTI2MSAyNTIuNjU4IDUyLjExMjhDMjUwLjQ1OCA3OC45MjUxIDIyOC4wMTYgMTAwIDIwMC42NTggMTAwSDE5MC4zODJDMTY1LjIzNCAxMDAgMTQ0LjI0MSA4Mi4xOTQxIDEzOS4zMDQgNTguNDkyNkMxMzguOTggNTYuMjI2NCAxMzguODEyIDUzLjkwOTQgMTM4LjgxMiA1MS41NTM2VjM3LjYxODNDMTM4LjgxMiAzMi4xOTUgMTM0LjQxOSAyNy43OTg1IDEyOSAyNy43OTg1QzEyMy41ODEgMjcuNzk4NSAxMTkuMTg4IDMyLjE5NSAxMTkuMTg4IDM3LjYxODNWNTEuNTUzNkMxMTkuMTg4IDUzLjkwOTQgMTE5LjAyIDU2LjIyNjQgMTE4LjY5NSA1OC40OTI2QzExMy43NTkgODIuMTk0MSA5Mi43NjU4IDEwMCA2Ny42MTc1IDEwMEg1Ny4zNDIxQzI5Ljk4MjkgMTAwIDcuNTQxOTQgNzguOTI1MSA1LjM0MjAxIDUyLjExMjhDNS4zNDAwNyA1MS45MjYxIDUuMzM4OTYgNTEuNzQwMyA1LjMzODk2IDUxLjU1MzZWMjYuNjgxMUM1LjMzODk2IDIzLjY0MjMgMy44MzUwNiAyMC44NzcxIDIuMzUyNjggMTguMTUwN0wyLjMzNDEyIDE4LjExNzRDMS45MjE1IDE3LjM1NzcgMS41MTA1NSAxNi42MDA3IDEuMTM1MTUgMTUuODQyOUMwLjQwODY0NCAxNC4zNzUyIDAgMTIuNzIyNyAwIDEwLjk3NDFDMCA0LjkxMzEzIDQuOTEgMCAxMC45NjY0IDBIOTcuMDUzNUMxMDIuNDk2IDAgMTA3Ljc2MiAxLjkyMzMgMTExLjkyNSA1LjQzMTYyTDExOC45ODQgMTEuMzc5OUMxMTkuMDA5IDExLjM4NjMgMTE5LjAzNSAxMS4zOTQ2IDExOS4wNTkgMTEuNDA1N0wxMTkuMDgyIDExLjQxNUMxMTkuMTE4IDExLjQzMTYgMTE5LjE1MyAxMS40NTM4IDExOS4xODcgMTEuNDgwNloiIGZpbGw9IiMwQzBDMEMiLz4KPHBhdGggZD0iTTE0Ni41MTYgMzEuNDk5MUMxNDYuNTE2IDE4LjczNzUgMTU2Ljg1MiA4LjM5Mzc0IDE2OS42MDMgOC4zOTM3NEgyMjEuNDM1QzIzNC4xODYgOC4zOTM3NCAyNDQuNTIzIDE4LjczNzUgMjQ0LjUyMyAzMS40OTkxVjQ3Ljc4MTlDMjQ0LjUyMyA3Mi4wMjc3IDIyNC44ODMgOTEuNjgyMSAyMDAuNjU3IDkxLjY4MjFIMTkwLjM4MUMxNjYuMTU1IDkxLjY4MjEgMTQ2LjUxNiA3Mi4wMjc3IDE0Ni41MTYgNDcuNzgxOVYzMS40OTkxWiIgZmlsbD0iIzA1MDUwNSIgc3Ryb2tlPSIjMDUwNTA1IiBzdHJva2Utd2lkdGg9IjAuMiIvPgo8cGF0aCBkPSJNMTMuNDc1OCAzMS40OTkxQzEzLjQ3NTggMTguNzM3NSAyMy44MTIzIDguMzkzNzQgMzYuNTYzIDguMzkzNzRIODguMzk0OUMxMDEuMTQ2IDguMzkzNzQgMTExLjQ4MiAxOC43Mzc1IDExMS40ODIgMzEuNDk5MVY0Ny43ODE5QzExMS40ODIgNzIuMDI3NyA5MS44NDMyIDkxLjY4MjEgNjcuNjE2OCA5MS42ODIxSDU3LjM0MTVDMzMuMTE1IDkxLjY4MjEgMTMuNDc1OCA3Mi4wMjc3IDEzLjQ3NTggNDcuNzgxOVYzMS40OTkxWiIgZmlsbD0iIzA1MDUwNSIgc3Ryb2tlPSIjMDUwNTA1IiBzdHJva2Utd2lkdGg9IjAuMiIvPgo8L3N2Zz4K" />
  </a>
</div>

<br />

## The stack

### Back-end

The back-end is written in [Go](https://go.dev/), uses [Gin](https://github.com/gin-gonic/gin) as its HTTP framework and is fully CORS compliant thanks to [Gin's CORS middleware](https://github.com/gin-contrib/cors).
The API's endpoints and methods are tested using [Testify](https://github.com/stretchr/testify).
The database chosen was [MongoDB](https://www.mongodb.com/) beacause of its ease of use, pricing and performance.

### Front-end

The front-end is written in Typescript and uses [React](https://reactjs.org/), [Next.js](https://nextjs.org/) and [SWR](https://swr.vercel.app/) under the hood to increase performance, accessibility and maintainability.
The styles and animations are made with [Mantine](https://mantine.dev/) and [Framer Motion](https://www.framer.com/motion/).
All the icons were provided by [Font Awesome](https://fontawesome.com/) through [React Icons](https://react-icons.github.io/react-icons/icons?name=fa).

<br />

## Development

1. Clone the repository running `git clone https://github.com/peterfritz/coinrank.git`, `git clone git@github.com:peterfritz/coinrank.git` or `gh repo clone peterfritz/coinrank` depending on your setup.
2. Open the directory cloned running `cd coinrank`.
3. Run the database, the back-end, and the front-end using `docker-compose up -d`.
4. Open a new tab on your browser and navigate to [`http://localhost:80/`](http://localhost:80/). If everything worked you should be able to see the application running.

> **Note**
> 
> The front-end uses the port [80](http://localhost:80/) and the back-end uses the port [8080](http://localhost:8080/).
