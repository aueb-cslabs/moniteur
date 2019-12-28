module.exports = {
    publicPath : process.env.NODE_ENV === 'production' ? './' : '/',
    pluginOptions: {
        electronBuilder: {
            builderOptions: {
                // options placed here will be merged with default configuration and passed to electron-builder
                "appId": "org.aueb.moniteur-admin",
                "productName": "moniteur-admin",
                "nsis": {
                    "oneClick": false,
                    "artifactName": "${productName}-Setup-${version}.${ext}",
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