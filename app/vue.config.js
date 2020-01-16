module.exports = {
    publicPath: process.env.NODE_ENV === 'production' ? './' : '/',
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
                    "target": {
                        "target": 'AppImage',
                        "arch": ['armv7l', 'x64']
                    },
                    "artifactName": "${productName}-${arch}-v${version}.${ext}"
                },
                "dmg": {
                    "artifactName": "${productName}-Setup-v${version}.${ext}",
                },
                "extraFiles": [
                    {
                        "from": "config/config.example.yml",
                        "to": "config.yml"
                    }
                ],
                "publish": [
                    {
                        "provider": "github",
                        "releaseType": "release"
                    }
                ]
            }
        }
    }
}