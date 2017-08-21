var devServerPort = '9000';

module.exports = {
  entry: [
    'webpack/hot/only-dev-server',
    './src/index.js'
  ],
  output: {
    path: __dirname + '../frontend',
    publicPath: 'http://localhost:' + devServerPort + '/public/',
    filename: 'bundle.js'
  },
  module: {
    loaders: [{
      test: /\.jsx?$/,
      exclude: /node_modules/,
      loader: 'react-hot-loader!babel-loader'
    }]
  },
  resolve: {
    extensions: ['*', '.js', '.jsx']
  },
  devServer: {
    contentBase: './dist',
    hot: true
  }
};
