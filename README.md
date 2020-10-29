# A Next.js starter for the [JAMstack](https://jamstack.org)
This is a boilerplate for using [Next.js](https://nextjs.org/) as a static site generator with a Go background function to download user gravatars and resize them on signup.

[![Deploy to Netlify](https://www.netlify.com/img/deploy/button.svg)](https://app.netlify.com/start/deploy?repository=https://github.com/futuregerald/next-function)

## Usage

### Getting started

To start your project, either:

1. Deploy to Netlify using the button above, or
2. Clone this repository and run:

```bash
npm install
```

This will take some time and will install all packages necessary to run the starter.

### Development

While developing your website, use:

```bash
npm start
```

Then visit http://localhost:3000/ to preview your new website. The Next.js development server will automatically reload the CSS or refresh the whole page, when stylesheets or content changes.

### Static build

To build a static version of the website inside the `/dist` folder, run:

```bash
npm run build
```

See [package.json](package.json) for all tasks.

## Basic Concepts

You can read more about building sites and apps with Next.js in their documentation here:

https://nextjs.org/docs

## Doing dynamic things

A few resources for doing anything you can imagine with a 100% static site/app on the JAMstack
using Next.js. If you would like to add more resources please open a pull request!

- [Using Next.js as a Static Site Generator for Netlify](https://scotch.io/@sw-yx/using-nextjs-as-a-static-site-generator-for-netlify) - [Shawn Wang](https://twitter.com/swyx)
- [Serverless Next.js 9 on Netlify Functions](https://community.netlify.com/t/serverless-next-js-9-on-netlify-functions/1956) - [Shawn Wang](https://twitter.com/swyx)

## Deploying to Netlify

The deploy to Netlify button above will create a new site and repo in one click. If you've created your repo manually, you can deploy to Netlify as follows:

- Push your clone to your own GitHub repository.
- [Create a new site on Netlify](https://app.netlify.com/start) and link the repository.

Now Netlify will build and deploy your site whenever you push to git.

## Background Function

This boilerplate has a background function written in go that is triggered on user signup. The function downloads their gravatar and resizes it to 3 different sizes and saves it in a github repo. Note that the repo name is hard-coded to this one and needs to be changed. In addition, you need to specify the following environment variables on your Netlify site:

- `GITHUB_COMMITTER_EMAIL`
- `GITHUB_COMMITTER_NAME`
- `GITHUB_OWNER`
- `GITHUB_REPO_NAME`
- `GITHUB_TOKEN`

In addition, the netlify.toml has an environment variable called `GO_IMPORT_PATH` that is not specific to this function, but to deploying Go functions in general. The value needs to be set to your github repo; for example `github.com/netlify/next-function` . You can read more about deploying Go functions on Netlify in the [Netlify docs](https://docs.netlify.com/functions/build-with-go/).