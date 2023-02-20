export const devServer = {
    proxy: {
        '/': {
            target: 'http://localhost:60200',
            ws: true,
            changeOrigin: true
        }
    }
};