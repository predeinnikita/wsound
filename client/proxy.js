const express = require("express");
const { createProxyMiddleware } = require("http-proxy-middleware");
const path = require("path");

const app = express();

const proxyMiddleware = createProxyMiddleware({
  target: "http://localhost:8080",
  changeOrigin: true,
  pathRewrite: {
    "http://localhost:5500/api": "",
  },
});

app.use("/api", proxyMiddleware);

app.get("/index.html", (req, res) => {
  res.sendFile(path.join(__dirname, "/index.html"));
});

app.get("/index.css", (req, res) => {
  res.sendFile(path.join(__dirname, "/index.css"));
});

app.get("/project.js", (req, res) => {
  res.sendFile(path.join(__dirname, "/project.js"));
});

app.get("/projects/:id", (req, res) => {
  res.status(200).sendFile(path.join(__dirname, "/project.html"));
});

app.get("/header.html", (req, res) => {
  res.status(200).sendFile(path.join(__dirname, "/header.html"));
});

app.get("/create-project.html", (req, res) => {
  res.sendFile(path.join(__dirname, "/create-project.html"));
});

app.listen(5500, () => {
  console.log("proxy start on port 5500");
});
