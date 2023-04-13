const path = require('path');

const resolvePath = p => path.resolve(__dirname, p)

module.exports = {
    webpack: {
        alias: {
            '@components': resolvePath('./src/components'),
            '@pages': resolvePath('./src/pages'),
            '@services': resolvePath('./src/services'),
            '@scss': resolvePath('./src/scss'),
            '@assets': resolvePath('./src/assets'),
        }
    },
}