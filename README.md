# Board-App

> Svelte와 Go언어를 사용한 CRUD 어플리케이션입니다

![go-image] 1.20.6 <br>
![mysql] 8.3.0 <br>
![svelte] 2.0.0 <br>

Svelte Kit으로 프론트를, Go언어로 백엔드를 사용한 CRUD 게시판 입니다.<br>
데이터베이스는 MySQL를 사용하였습니다. <br>

현재 업데이트 프론트 기능이 작동을 안합니다.<br>
백엔드는 정상 작동합니다. (고쳐주실분은 컨트리뷰트를..)

## 디렉토리

```
📦Board-App
 ┣ 📜 front - 스벨트 파일
 ┣ 📜 Back - Go언어 서버 파일
 ┗ 📜 README.md - 프로젝트의 README 파일
```

## Go

### 사용 프레임워크

gonic-gin <br>
google/uuid <br>
gorm <br>
godotenv

### install

```shell
go get -u
```

### Using

```shell
go run main.go
```

## Svelte

### Developing

Once you've created a project and installed dependencies with `npm install` (or `pnpm install` or `yarn`), start a development server:

```bash
npm run dev

# or start the server and open the app in a new browser tab
npm run dev -- --open
```

### Building

To create a production version of your app:

```bash
npm run build
```

You can preview the production build with `npm run preview`.

> To deploy your app, you may need to install an [adapter](https://kit.svelte.dev/docs/adapters) for your target environment.

<!-- Markdown link & img dfn's -->

[go-image]: https://img.shields.io/badge/go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[mysql]: https://img.shields.io/badge/mysql-4479A1?style=for-the-badge&logo=mysql&logoColor=white
[svelte]: https://img.shields.io/badge/svelte_kit-FF3E00?style=for-the-badge&logo=svelte&logoColor=white
