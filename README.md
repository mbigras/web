# Web

> App web is for experimenting with application delivery workflows.

## Summary

This app is for experimenting with application delivery workflows.

## Getting started

The following procedure describes how to build and run web app.

1. Get the code.

   ```
   git clone git@github.com:mbigras/web.git
   cd web
   ```

1. Build container.

   ```
   docker build -t web .
   ```

1. Run container.

   ```
   docker run -it -p 8080:8080 web
   ```

1. Navigate to http://localhost:8080/ in your browser.

1. To stop all apps, navigate back to your terminal and press Ctrl-C.

   Your output should look like the following:

   ```
   $ docker run -it -p 8080:8080 web
   2023/12/19 20:09:44 running as pid 1
   2023/12/19 20:09:44 listening on port 8080
   ^C2023/12/19 20:09:55 caught signal: interrupt; shutting down; byeee!!
   ```
