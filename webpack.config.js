var path = require("path");
var HtmlWebpackPlugin = require("html-webpack-plugin");
var CopyWebpackPlugin = require("copy-webpack-plugin");

module.exports = {
  watch: true,
  entry: "./assets/index.js",
  output: {
    path: path.resolve(__dirname, "public"),
    filename: "index_bundle.js",
    publicPath: "http://localhost:8000/",
  },
  module: {
    rules: [
      { test: /\.(js)$/, use: "babel-loader" },
      { test: /\.css$/, use: ["style-loader", "css-loader"] },
      { test: /\.(png|svg|jpg|gif)$/, use: ["file-loader"] },
    ],
  },
  mode: "development",
  plugins: [
    new HtmlWebpackPlugin({
      template: "./assets/index.html",
    }),
    // new CopyWebpackPlugin({
    //   patterns: [{ from: "./assets/images", to: "images" }],
    // }),
  ],
};
