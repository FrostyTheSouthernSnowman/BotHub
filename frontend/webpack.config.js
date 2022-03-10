const path = require('path');

module.exports = {
    context: path.resolve(__dirname),
    devtool: 'inline-source-map',
    entry: './main.ts',
    mode: 'development',
    module: {
        rules: [{
            test: /\.ts$/,
            use: 'ts-loader',
            exclude: /node_modules/
        }]
    },
    output: {
        filename: 'freeBots.js',
        path: path.resolve(__dirname),
    },
    resolve: {
        extensions: ['.ts', '.js']
    },
};