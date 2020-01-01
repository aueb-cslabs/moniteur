module.exports = {
    publicPath : process.env.NODE_ENV === 'production' ? './' : '/',
    pluginOptions: {
        electronBuilder: {
            builderOptions: {
                // options placed here will be merged with default configuration and passed to electron-builder
                "appId": "org.aueb.moniteur",
                "nsis": {
                    "oneClick": false,
                    "artifactName": "${productName}-Setup-v${version}.${ext}"
                },
                "linux": {
                    "target": "AppImage",
                    "artifactName": "${productName}-v${version}.${ext}"
                },
                "dmg": {
                    "artifactName": "${productName}-Setup-v${version}.${ext}",
                },
                "extraFiles": [
                    {
                        "from": "config/config.example.yml",
                        "to": "config.yml"
                    }
                ]
            }
        }
    }
}