module.exports = {
    publicPath : process.env.NODE_ENV === 'production' ? './' : '/',
    pluginOptions: {
        electronBuilder: {
            builderOptions: {
                // options placed here will be merged with default configuration and passed to electron-builder
                "appId": "org.aueb.moniteur-admin",
                "productName": "Moniteur Admin",
                "nsis": {
                    "oneClick": false,
                    "artifactName": "moniteur-admin-Setup-${version}.${ext}",
                },
                "linux": {
                    "target": "deb"
                },
                "dmg": {

                },
                "extraFiles": [
                    "config/*"
                ]
            }
        }
    }
}