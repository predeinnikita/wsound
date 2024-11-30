const { createProxyMiddleware } = require("http-proxy-middleware");

module.exports = function (app) {
  app.use(
    "/api",
    createProxyMiddleware({
      target: "http://localhost:8080",
      // changeOrigin: true,
      pathRewrite: {
        "/api": "",
      },
      on: {
        proxyReq: (p, req, res) => console.log(p, req, res),
        proxyRes: (p, req, res) => console.log(p, req, res),
        start: (p, req, res) => console.log(p, req, res),
        error: (p, req, res) => console.log(p, req, res),
      },
    })
  );
};
